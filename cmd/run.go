package cmd

import (
	"github.com/urfave/cli"
)

var RunCommand = cli.Command{
	Name:      "run",
	Usage:     "[Deprecated] Run command inside Barcelona environment",
	ArgsUsage: "",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		return cli.NewExitError("This command is deprecated. Please use ecs-exec instead, available at https://github.com/komoju/komoju-ecs-exec", 1)
	},
}
