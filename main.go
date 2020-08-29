package main

import (
	"log"
	"os"

	"github.com/RyuseiNomi/stash/src/handler"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "stash",
		Usage:   "an effective tool to stash your work",
		Version: "1.0.0",
		Commands: []*cli.Command{
			{
				Name:  "save",
				Usage: "Save your current work",
				Action: func(c *cli.Context) error {
					sw := handler.NewStashWorker("save")
					sw.Handle()
					return nil
				},
			},
			{
				Name:  "pop",
				Usage: "Fetch work which is saved temporarily",
				Action: func(c *cli.Context) error {
					sw := handler.NewStashWorker("pop")
					sw.Handle()
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "Show the list of stashed work",
				Action: func(c *cli.Context) error {
					sw := handler.NewStashWorker("list")
					sw.Handle()
					return nil
				},
			},
			{
				Name:  "show",
				Usage: "Show the detail difference each saved work",
				Action: func(c *cli.Context) error {
					sw := handler.NewStashWorker("show")
					sw.Handle()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
