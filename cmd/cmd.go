package cmd

import (
	"context"

	"github.com/urfave/cli/v2"
	"github.com/nekohor/mangosteen/internal/app"
)


func NewRunCmd(ctx context.Context, cliApp *cli.App) *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "运行命令行",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "配置文件(.json,.yaml,.toml)",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			return app.RunCommand(ctx,
				app.SetConfigFile(c.String("conf")),
				app.SetVersion(cliApp.Version))
		},
	}
}

func NewWebCmd(ctx context.Context, cliApp *cli.App) *cli.Command {
	return &cli.Command{
		Name:  "serve",
		Usage: "运行web服务",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "配置文件(.json,.yaml,.toml)",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			return app.RunService(ctx,
				app.SetConfigFile(c.String("conf")),
				app.SetVersion(cliApp.Version))
		},
	}
}
