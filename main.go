package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/mu7ammad1951/gator/internal/config"
	"github.com/mu7ammad1951/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer db.Close()

	myState := state{
		cfg: &cfg,
		db:  database.New(db),
	}

	cmds := commands{
		handlers: make(map[string]func(s *state, cmd command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)

	if len(os.Args) < 2 {
		log.Fatal("Error: not enough arguments provided")
	}

	name := os.Args[1]  // The command name
	args := os.Args[2:] // Everything after the command name

	myCommand := command{
		name: name,
		args: args,
	}

	err = cmds.run(&myState, myCommand)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}

}
