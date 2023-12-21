package app

import "github.com/urfave/cli"

// Returns the CLI App ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "WebInspect"
	app.Usage = "Searches for a website IP address and server name"
	
	return app
}