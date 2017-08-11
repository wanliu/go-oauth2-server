package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/wanliu/go-oauth2-server/cmd"
)

var (
	cliApp        *cli.App
	configBackend string
)

func init() {
	// Initialise a CLI app
	cliApp = cli.NewApp()
	cliApp.Name = "wanliu-oauth2-server"
	cliApp.Usage = "Wanliu OAuth 2.0 Server"
	cliApp.Author = "Hysios Hu"
	cliApp.Email = "hysios@gmail.com"
	cliApp.Version = "0.0.0"
	cliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "configBackend",
			Value:       "etcd",
			Destination: &configBackend,
		},
	}
}

func main() {
	// Set the CLI app commands
	cliApp.Commands = []cli.Command{
		{
			Name:  "migrate",
			Usage: "run migrations",
			Action: func(c *cli.Context) error {
				return cmd.Migrate(configBackend)
			},
		},
		{
			Name:  "loaddata",
			Usage: "load data from fixture",
			Action: func(c *cli.Context) error {
				return cmd.LoadData(c.Args(), configBackend)
			},
		},
		{
			Name:  "runserver",
			Usage: "run web server",
			Action: func(c *cli.Context) error {
				return cmd.RunServer(configBackend)
			},
		},
		{
			Name:  "setup",
			Usage: "setup oauth server",
			Action: func(c *cli.Context) error {
				return cmd.Setup(configBackend)
			},
		},
	}

	// Run the CLI app
	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
