package main

import (
	"fmt"
	"log"
	"os"

	"commands/workitems"

	"github.com/urfave/cli/v2"
)

func main() {
	var commands []*cli.Command = []*cli.Command{}
	commands = append(commands, workitems.Definition)

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "dt",
				Usage: "Track time",
				Action: func(*cli.Context) error {
					fmt.Println("action: Track time")
					return nil
				},
			}, {
				Name:  "add",
				Usage: "Add time",
				Action: func(*cli.Context) error {
					fmt.Println("action: add")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
