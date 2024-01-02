package app

import (
	"github.com/urfave/cli/v2"
)

// Returns the CLI App ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "WebInspect"
	app.Usage = "Searches for a website IP address and server name"
	
	app.Commands = []*cli.Command{
		{
			Name: "ip",
			Usage: "Search IP address",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "domain",
					Aliases: []string{"d"},
				},
			},
			Action: SearchIps,
		},
		{
			Name: "servers",
			Usage: "Search server names",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "domain",
					Aliases: []string{"d"},
				},
			},
			Action: SearchServers,
		},
		{
			Name: "health",
			Usage: "Checks the server health",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "port",
					Aliases: []string{"p"},
				},
				&cli.StringFlag{
					Name: "domain",
					Aliases: []string{"d"},
				},
			},
			Action: HealthCheck,
		},
	}

	return app
}