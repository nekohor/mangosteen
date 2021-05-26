package bootstrap

import (
	"github.com/nekohor/mangosteen/internal/app"
	"github.com/nekohor/mangosteen/internal/app/config"
	"github.com/nekohor/mangosteen/internal/app/dbx"
	"log"
)

func InitForTest() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// init config
	config.MustLoad("D:/NutCloudSync/proj/mangosteen/bin/mangosteen/conf.toml")
	//config.PrintWithJSON()

	// init logger module
	app.InitSimpleLogger()

	//init database
	_, err := dbx.InitGormDB()
	if err != nil {
		log.Println(err)
	}
}
