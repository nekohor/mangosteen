package services

import (
	"log"
	"github.com/nekohor/mangosteen/internal/app"
	"github.com/nekohor/mangosteen/internal/app/config"
	"github.com/nekohor/mangosteen/internal/app/dbx"
	"testing"
)

func TestUnqualSaveService_SaveUnqualHistoriesByCoilId(t *testing.T) {
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
	s := NewUnqualSaveService()
	s.SaveUnqualHistoriesByCoilId("H21066013L")
}

func TestUnqualSaveService_SaveUnqualHistoriesByDate(t *testing.T) {
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
	s := NewUnqualSaveService()
	millLines := []string{"MG2250", "MG1580"}
	//millLines := []string{"MG1580"}
	for _, millLine := range millLines {
		s.SaveUnqualHistoriesByDate("2021-05-25", millLine)
	}
}