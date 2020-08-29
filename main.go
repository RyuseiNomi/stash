package main

import (
	"github.com/RyuseiNomi/stash/src/handler"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "stash",
		Usage: "an effective tool to stash your work",
		Action: func(c *cli.Context) error {
			sw := handler.NewStashWorker(c.Args().Get(0))
			sw.Handle()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
