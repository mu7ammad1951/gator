package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/mu7ammad1951/gator/internal/database"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	params := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		ID:            nextFeed.ID,
	}
	_, err = s.db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error updating last_fetched_at: %w", err)
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed by url: %w", err)
	}
	fmt.Println()
	fmt.Printf("Feed => %v\n", feed.Channel.Title)

	for _, item := range feed.Channel.Item {
		now := time.Now().UTC()
		pubDate, err := parseDate(item.PubDate)
		if err != nil {
			fmt.Printf("error parsing the date, using 'now': %v\n", err)
			pubDate = now
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   now,
			UpdatedAt:   now,
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: pubDate,
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code != "23505" {
					return fmt.Errorf("error creating post %s from %s: %w", item.Title, feed.Channel.Title, err)
				}
			}
		}
	}
	fmt.Println()

	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	var rssFeed *RSSFeed = &RSSFeed{}
	if err = xml.Unmarshal(data, rssFeed); err != nil {
		return &RSSFeed{}, err
	}

	unescapeRSSFeed(rssFeed)

	return rssFeed, nil

}

func unescapeRSSFeed(rssFeed *RSSFeed) {
	if rssFeed == nil {
		return
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	for i, rssItem := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(rssItem.Title)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(rssItem.Description)
	}

}

func parseDate(dateStr string) (time.Time, error) {
	formats := []string{
		time.RFC3339,
		time.RFC1123Z,
		time.RFC1123,
		"Mon, 02 Jan 2006 15:04:05 -0700",
		// Add more formats as needed
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("could not parse date: %s", dateStr)
}
