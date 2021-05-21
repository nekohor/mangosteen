package services

import (
	"log"
	"github.com/nekohor/mangosteen/internal/app/dbx"
	fspDaoPack "github.com/nekohor/mangosteen/internal/apps/fsp/dao"
	"github.com/nekohor/mangosteen/internal/apps/unqualified/models"
	"github.com/nekohor/mangosteen/internal/apps/unqualified/schemas"
	"github.com/nekohor/mangosteen/pkg/logger/zap/logger"
	"github.com/nekohor/mangosteen/pkg/util/convert"
	"time"
)

func GetUnqualHistoryInfo(unqualCategory schemas.UnqualIdiom, unqualSource schemas.UnqualIdiom,  unqualType schemas.UnqualIdiom, unqualValue string) *schemas.UnqualHistoryInfo {
	uhi := &schemas.UnqualHistoryInfo{}
	uhi.UnqualCategory = unqualCategory.Name()
	uhi.UnqualSource = unqualSource.Name()
	uhi.UnqualType = unqualType.Name()
	uhi.UnqualValue = unqualValue
	return uhi
}

func GetUnqualHistory(cmr *fspDaoPack.CoilMainRecord, uhi *schemas.UnqualHistoryInfo) *models.UnqualHistory {
	var err error
	var productTime time.Time
	e := &models.UnqualHistory{}
	e.CoilId = cmr.CoilId
	productTime, err = time.ParseInLocation("20060102 15:04:05", cmr.ProductTime, time.Local)
	if err != nil {
		logger.Info(err)
	}
	e.ProductTime = &productTime
	e.SteelGrade = cmr.SteelGrade
	e.AimThk = convert.ToFloat64(cmr.AimThick)
	e.AimWid = convert.ToInt(cmr.AimWidth)
	e.UnqualCategory = uhi.UnqualCategory
	e.UnqualSource = uhi.UnqualSource
	e.UnqualType = uhi.UnqualType
	e.UnqualValue = uhi.UnqualValue
	return e
}

func SaveUnqualHistory(cmr *fspDaoPack.CoilMainRecord, uhi *schemas.UnqualHistoryInfo) {
	db := dbx.GormDB("unqual")
	unqualHistory := GetUnqualHistory(cmr, uhi)

	//var cnt int
	//db.Where(unqualHistory).Count(&cnt)
	//db.NewRecord()
	//if cnt != 0 {
	//	unqualHistory.AddCreateTime()
	//	db.Create(unqualHistory)
	//} else {
	//	db.Where(unqualHistory).FirstOrCreate(unqualHistory)
	//}

	err := db.Where(unqualHistory).FirstOrCreate(&models.UnqualHistory{}).Error; if err != nil {
		log.Println(err)
	}
}