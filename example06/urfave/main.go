package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

type (
	// Config information.
	Config struct {
		username string
		password string
	}
)

var config Config

func main() {
	app := cli.NewApp()
	app.Name = "Example"
	app.Usage = "Example"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "username,u",
			Usage: "user account",
		},
		cli.StringFlag{
			Name:  "password,p",
			Usage: "user password",
		},
	}

	app.Run(os.Args)
}

func run(c *cli.Context) error {
	config = Config{
		username: c.String("username"),
		password: c.String("password"),
	}

	return exec()
}

func exec() error {
	fmt.Println("username:", config.username)
	fmt.Println("password:", config.password)

	return nil
}
