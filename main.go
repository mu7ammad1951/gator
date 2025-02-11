package main

import (
	"fmt"
	"os"

	"github.com/mu7ammad1951/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	myState := state{
		cfg: &cfg,
	}

	commands := commands{
		handlers: make(map[string]func(s *state, cmd command) error),
	}
	commands.register("login", handlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("Error: not enough arguments provided")
		os.Exit(1)
	}

	name := os.Args[1]  // The command name
	args := os.Args[2:] // Everything after the command name

	myCommand := command{
		name: name,
		args: args,
	}

	err = commands.run(&myState, myCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
