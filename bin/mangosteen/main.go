package main

import (
	"context"
	"os"
	_ "github.com/nekohor/mangosteen/docs/swagger"
	"github.com/nekohor/mangosteen/cmd"
	"github.com/urfave/cli/v2"
)

// VERSION 版本号，可以通过编译的方式指定版本号：go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "0.1.1"

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
		cmd.NewWebCmd(ctx, application),
		cmd.NewRunCmd(ctx, application),
	}
	err := application.Run(os.Args)
	if err != nil {
		panic(err.Error())
	}
}


