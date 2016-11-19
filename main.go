package main

import (
	"crypto/sha1"
	"fmt"
	"os"
	"strings"

	"gopkg.in/urfave/cli.v1"
)

func createHash(params ...string) string {
	joinedParams := strings.Join(params, " ")

	h := sha1.New()
	h.Write([]byte(joinedParams))
	byteSlice := h.Sum(nil)

	return string(byteSlice)
}

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

		userHash := createHash(email, ip)
		fmt.Printf("generated hash: %s\n", userHash)

		return nil
	}

	app.Run(os.Args)
}
