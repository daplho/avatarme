package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	var email, ip string

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "email",
			Value:       "avatar@me.com",
			Usage:       "e-mail address",
			Destination: &email,
		},
		cli.StringFlag{
			Name:        "ip-address, ip",
			Value:       "127.0.0.1",
			Usage:       "ip address",
			Destination: &ip,
		},
	}

	app.Action = func(c *cli.Context) error {
		email := c.String("email")
		fmt.Printf("received address: %s\n", email)

		ip := c.String("ip-address")
		fmt.Printf("received ip address: %s\n", ip)

		return nil
	}

	app.Run(os.Args)
}
