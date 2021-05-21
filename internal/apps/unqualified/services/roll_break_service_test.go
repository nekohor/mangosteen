package services

import (
	"log"
	"github.com/nekohor/mangosteen/internal/app"
	"github.com/nekohor/mangosteen/internal/app/config"
	"github.com/nekohor/mangosteen/internal/app/dbx"
	"testing"
)

func TestRollBreakStatService_GetRollBreakStatResult(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// init config
	config.MustLoad("D:/NutCloudSync/proj/mangosteen/cmd/mangosteen/conf.toml")
	//config.PrintWithJSON()

	// init logger module
	app.InitSimpleLogger()

	//init database
	_, err := dbx.InitGormDB()
	if err != nil {
		log.Println(err)
	}

	s := NewRollBreakStatService()
	res := s.GetRollBreakStatResult("2021-05-20")
	log.Println(res.RollBreakTotalCount)
}