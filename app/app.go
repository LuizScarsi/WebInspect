package app

import (
	"github.com/urfave/cli/v2"
)

// Generate returns the CLI App ready to be executed
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
					Aliases: []string{"n"},
					Usage: "Specify the host for IP search",
					Required: true,
				},
			},
			Action: SearchIps,
		},
		{
			Name: "txt",
			Usage: "Search DNS TXT records",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "domain",
					Aliases: []string{"d"},
					Usage: "Specify the domain name for TXT search",
					Required: true,
				},
			},
			Action: SearchDNSTxt,
		},
		{
			Name: "servers",
			Usage: "Search server names",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "domain",
					Aliases: []string{"d"},
					Usage: "Specify the domain name for server name search",
					Required: true,
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
					Usage: "Specify port to be checked",
					Aliases: []string{"p"},
				},
				&cli.StringFlag{
					Name: "domain",
					Aliases: []string{"d"},
					Usage: "Specify the domain name for health check",
					Required: true,
				},
			},
			Action: HealthCheck,
		},
	}

	return app
}