package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	version = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "pipeline wait plugin"
	app.Usage = "pipeline wait plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token",
			Usage:  "drone API token",
			EnvVar: "PLUGIN_TOKEN",
			Hidden: true,
		},
		cli.StringSliceFlag{
			Name:   "waitpipelines",
			Usage:  "list of wait pipelines",
			EnvVar: "PLUGIN_WAIT_PIPELINES,WEBHOOK_WAIT_PIPELINES",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Number:  c.Int("build.number"),
			Link:    c.String("build.link"),
		},
		Config: Config{
			WaitPipelines: c.StringSlice("waitpipelines"),
			Token:         c.String("token"),
		},
	}

	return plugin.Exec()
}
