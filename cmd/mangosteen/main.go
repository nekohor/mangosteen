package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/nekohor/mangosteen/internal/app"
	_ "github.com/nekohor/mangosteen/internal/app/swagger"
)

// VERSION 版本号，可以通过编译的方式指定版本号：go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "0.0.1"

// @title 宇阳的杂七杂八微服务
// @version 0.0.1
// @description 提供与FSP平台基础业务有关的杂七杂八微服务

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @schemes http https
// @basePath /
// @contact.name 王宇阳
// @contact.email nekohor@foxmail.com
func main() {

	ctx := context.WithValue(context.Background(), "__name__", "__main__")
	application := cli.NewApp()
	application.Name = "mangosteen"
	application.Version = VERSION
	application.Usage = "Public Api"
	application.Commands = []*cli.Command{
		newWebCmd(ctx),
		newRunCmd(ctx),
	}
	err := application.Run(os.Args)
	if err != nil {
		panic(err.Error())
	}
}

func newRunCmd(ctx context.Context) *cli.Command {
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
				app.SetVersion(VERSION))
		},
	}
}

func newWebCmd(ctx context.Context) *cli.Command {
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
				app.SetVersion(VERSION))
		},
	}
}
