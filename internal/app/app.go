package app

import (
	"context"
	"fmt"
	"log"
	"github.com/nekohor/mangosteen/internal/app/dbx"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nekohor/mangosteen/internal/app/config"
	"github.com/nekohor/mangosteen/pkg/logger/zap/logger"
)

type options struct {
	ConfigFilePath string
	WWWDir         string
	Version        string
}

// Option 定义配置项
type Option func(*options)

// SetConfigFile 设定配置文件
func SetConfigFile(s string) Option {
	return func(o *options) {
		if s == "" {
			o.ConfigFilePath = "conf.toml"
		} else {
			o.ConfigFilePath = s
		}
	}
}

// SetVersion 设定版本号
func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

func PrimaryInit(configFilePath string) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// init config
	config.MustLoad(configFilePath)
	//config.MustLoad("D:/NutCloudSync/proj/mangosteen/cmd/mangosteen/conf.toml")
	//config.PrintWithJSON()

	// init logger module
	InitSimpleLogger()

	//init database
	_, err := dbx.InitGormDB()
	if err != nil {
		log.Println(err)
	}
}

// Init 应用初始化
func Init(ctx context.Context, opts ...Option) (func(), error) {

	var o options
	for _, opt := range opts {
		opt(&o)
	}

	// init config
	config.MustLoad(o.ConfigFilePath)
	config.PrintWithJSON()

	// print init
	fmt.Printf("服务启动，运行模式：%s，版本号：%s，进程号：%d", config.C.RunMode, o.Version, os.Getpid())
	fmt.Println()

	// init logger module
	InitSimpleLogger()
	loggerCleanFunc, err := InitLogger()
	if err != nil {
		return nil, err
	}

	//init database
	dbCleanFunc, err := InitDB()
	if err != nil {
		return nil, err
	}

	// 初始化依赖注入器
	injector, injectorCleanFunc, err := BuildInjector()
	if err != nil {
		return nil, err
	}

	httpServerCleanFunc := InitHTTPServer(ctx, injector.Engine)

	return func() {
		httpServerCleanFunc()
		injectorCleanFunc()
		dbCleanFunc()
		loggerCleanFunc()
	}, nil
}

// Run 运行服务
func RunService(ctx context.Context, opts ...Option) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := Init(ctx, opts...)
	if err != nil {
		return err
	}

	// 在Init()中的另一goroutine中打开server
EXIT:
	for {
		sig := <-sc
		logger.Infof("接收到信号[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	logger.Infof("服务退出")
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}

// Run 运行服务
func RunCommand(ctx context.Context, opts ...Option) error {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := Init(ctx, opts...)
	if err != nil {
		return err
	}

	//RunOnce()

	cleanFunc()
	logger.Infof("结束程序")
	return nil
}
