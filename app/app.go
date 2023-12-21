package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Returns the CLI App ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "WebInspect"
	app.Usage = "Searches for a website IP address and server name"
	
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com.br",
		},
	}

	
	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Search IP address",
			Flags: flags,			
			Action: searchIps,
		},
		{
			Name: "servers",
			Usage: "Search server names",
			Flags: flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")
	
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")
	
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}