package main

import (
	"log"

	"github.com/jbabin91/net-go-api/cmd"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
