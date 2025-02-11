package main

import (
	"fmt"

	"github.com/mu7ammad1951/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	fmt.Println(cfg)

	cfg.CurrentUserName = "AbdulKareem"

	err = config.Write(cfg)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println("Error re-reading file: ", err)
		return
	}
	fmt.Println(cfg)

}
