package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"log"
)

func main() {
	app := cli.NewApp()
	app.Name = "Boom"
	app.Usage = "Make an explosive entrance"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
