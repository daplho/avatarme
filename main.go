package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "cool avatar"
	app.Usage = "blow them out of the water"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom!")
		return nil
	}

	app.Run(os.Args)
}
