package services_test

import (
	"github.com/nekohor/mangosteen/internal/app/bootstrap"
	"github.com/nekohor/mangosteen/internal/apps/unqualified/services"
	"testing"
)

func TestUnqualSaveService_SaveUnqualHistoriesByCoilId(t *testing.T) {

	// init
	bootstrap.InitForTest()

	// test
	s := services.NewUnqualSaveService()
	s.SaveUnqualHistoriesByCoilId("H21066013L")
}

func TestUnqualSaveService_SaveUnqualHistoriesByDate(t *testing.T) {

	// init
	bootstrap.InitForTest()

	// test
	s := services.NewUnqualSaveService()
	millLines := []string{"MG2250", "MG1580"}
	//millLines := []string{"MG1580"}
	for _, millLine := range millLines {
		s.SaveUnqualHistoriesByDate("2021-05-24", millLine)
	}
}