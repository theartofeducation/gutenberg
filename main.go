package main

import (
	"fmt"
	"os"

	"github.com/theartofeducation/gutenberg/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("you need to give us an action to run")

		os.Exit(1)
	}

	action := os.Args[1]

	switch action {
	case "version":
		cmd.Version.Execute()
	case "tag":
		cmd.Tag.Execute()
	case "publish":
		cmd.Publish.Execute()
	default:
		fmt.Println("that's not an option")
	}
}
