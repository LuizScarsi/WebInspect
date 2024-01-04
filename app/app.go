package app

import (
	"github.com/urfave/cli/v2"
)

// Returns the CLI App ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "WebInspect"
	app.Usage = "Tool that provides informations about websites"
	
	app.Commands = []*cli.Command{
		{
			Name: "ip",
			Usage: "Search IP address",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "host",
					Usage: "Specify the host for IP seach",
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