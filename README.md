# Gator - A CLI Blog Aggregator

Gator is an RSS blog aggregator with a Command Line Interface (CLI). It allows you to register multiple users, add feeds, periodically check for new content, follow feeds from other users, and more.

## Installation Instructions

### Prerequisites

- **Go** (version **1.23.6** or higher)
- **PostgreSQL** (version **16.6** or higher)
- **Goose** (version **3.24.1** or higher)

### Steps

1. **Verify Go Installation**  
   Ensure Go is installed correctly by running:
   ```sh
   go version
   ```
   *Expected output:*
   ```sh
   go version go1.23.6 linux/amd64
   ```

2. **Verify PostgreSQL Installation**  
   Check your PostgreSQL installation with:
   ```sh
   psql --version
   ```
   *Expected output:*
   ```sh
   psql (PostgreSQL) 16.6 (Ubuntu 16.6-0ubuntu0.24.04.1)
   ```

3. **Verify Goose Installation**  
   Confirm that Goose is installed by running:
   ```sh
   goose --version
   ```
   *Expected output:*
   ```sh
   goose version: v3.24.1
   ```

4. **Install Gator**  
   Install Gator using:
   ```sh
   go install github.com/mu7ammad1951/gator@latest
   ```

## Setup

After installing the `gator` program, complete the following steps:

1. **Create the Configuration File**  
   Create a `.gatorconfig.json` file in your home directory (`$HOME`) with the following content:
   ```json
   {
     "db_url": "<your-postgres-url>",
     "current_user_name": "<your-username>"
   }
   ```
   Replace `<your-postgres-url>` with your PostgreSQL connection string in the format:
   ```
   protocol://username:password@host:port/database
   ```

2. **Run Database Migrations**  
   Apply the necessary database migrations with Goose:
   ```sh
   goose postgres <your-postgres-url> up
   ```
   This command creates the necessary tables and columns in your database.

## Usage

### Commands

#### `register <name>`

**Description:**  
Registers a new user with the specified `<name>` and sets that user as the current user.

**Usage Example:**  
```sh
gator register your_username
```
Replace `your_username` with the actual username you wish to register.

---

#### `login <name>`

**Description:**  
Logs in an existing user with the specified `<name>` and sets that user as the current session.

**Usage Example:**  
```sh
gator login your_username
```
Replace `your_username` with the actual username you wish to log in as. Once logged in, you can execute commands that require authentication.

---

### `users`

**Description:**
Lists the names of all the users and identifies the current user.

**Usage Example:**  
```sh
gator users
```

Below is a suggested documentation section for the `agg`, `addfeed`, and `feeds` commands to include in your README:

---

#### `agg <duration-string>`

**Description:**  
Starts an ongoing process that aggregates (scrapes) feeds at a specified time interval. The `<duration-string>` should be a valid Go duration (e.g., `10s`, `5m`, or `1h`).  
If an invalid duration is provided, an error message will indicate the proper format.

**Usage Example:**  
```sh
gator agg 30s
```

**Notes:**  
- The command prints a message indicating the interval at which feeds are being collected.  
- The aggregation runs indefinitely until the process is stopped.

---

#### `addfeed <name> <url>`

**Description:**  
Adds a new feed to the aggregator using the provided `<name>` and `<url>`. Once added, the feed is automatically followed by the current user.

**Usage Example:**  
```sh
gator addfeed TechNews https://example.com/rss
```

**Notes:**  
- This command requires the user to be logged in.  
- On success, it confirms that the feed is followed and displays the user and feed details.

---

#### `feeds`

**Description:**  
Lists all the feeds currently available in the system. For each feed, it displays the feed's name, URL, and the username of the person who added it.

**Usage Example:**  
```sh
gator feeds
```

**Notes:**  
- The command retrieves and prints information for each feed stored in the database.

---

#### `browse [limit]`

**Description:**  
Displays recent posts for the current user from the feeds they follow. This command retrieves a specified number of posts, with a default limit of 2 if no argument is provided.

- **Optional Argument:**  
  - `limit`: An integer value representing the maximum number of posts to display.
  - If the provided argument cannot be converted to an integer or is omitted, the command defaults to 2.

**Usage Examples:**  
```sh
gator browse
```
*This command fetches 2 posts by default.*

```sh
gator browse 5
```
*This command fetches up to 5 posts for the current user.*

**Notes:**  
- Ensure you are logged in before using this command, as it fetches posts specific to the current user.
- In case of an invalid input for the limit (e.g., non-integer), a warning is printed and the limit is defaulted to 2.

---

#### `follow <url>`

**Description:**  
Allows the current user to follow an existing feed by specifying its `<url>`. The command looks up the feed using the given URL and, if found, creates a feed-follow record linking the feed to the current user. If the feed does not exist, an error is returned, prompting you to add the feed using the `addfeed <name> <url>` command.

**Usage Example:**  
```sh
gator follow https://example.com/rss
```

**Notes:**  
- **Authentication Required:** You must be logged in to follow a feed.
- **Prerequisite:** The feed must already exist in the system.

---

#### `following`

**Description:**  
Lists all feeds that the current user is following. For each followed feed, it displays the feed's name. If the user is not following any feeds, a message is printed indicating that no feed follows were found.

**Usage Example:**  
```sh
gator following
```

**Notes:**  
- **Authentication Required:** You must be logged in to view the list of feeds you follow.

---

#### `unfollow <feed-url>`

**Description:**  
Unfollows a feed for the current user by specifying the feed's `<feed-url>`. This command deletes the feed-follow record associated with the provided URL and the current user.

**Usage Example:**  
```sh
gator unfollow https://example.com/rss
```

**Notes:**  
- **Authentication Required:** You must be logged in to unfollow a feed.
- **Action:** On successful execution, it confirms that the feed has been unfollowed.



