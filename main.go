package main

import (
	"github.com/c-j-j/faros_deployer/commands"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Faros"
	app.Usage = "Deployment Tool"

	app.Commands = availableCommands()
	app.Run(os.Args)
}

func availableCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "list-docker-containers",
			Aliases: []string{"ldc"},
			Usage:   "List all docker containers",
			Action:  func(c *cli.Context) { commands.ListDockerContainers{}.Execute() },
		},
	}
}
