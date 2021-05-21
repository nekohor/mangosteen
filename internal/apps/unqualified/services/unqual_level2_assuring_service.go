package services

import (
	"fmt"
	"log"
	fspDaoPack "github.com/nekohor/mangosteen/internal/apps/fsp/dao"
	"github.com/nekohor/mangosteen/internal/apps/unqualified/schemas"
	"github.com/nekohor/mangosteen/pkg/tags"
)

type UnqualLevel2AssuringService struct {
}

func NewUnqualifiedLevel2AssuringService() *UnqualLevel2AssuringService {
	return &UnqualLevel2AssuringService{}
}

func (s *UnqualLevel2AssuringService) SaveUnqualHistoriesByCoilId(coilId string) {
	coilMainDao := fspDaoPack.NewCoilMainDao()
	cmr := coilMainDao.GetRecordByCoilId(coilId)
	log.Println(cmr)
	if tags.IsCoilIdValid(cmr.CoilId) {
		s.SaveUnqualHistoriesByCoilMainRecord(cmr)
	}
}

func (s *UnqualLevel2AssuringService) SaveUnqualHistoriesByCoilMainRecord(cmr *fspDaoPack.CoilMainRecord) {
	assuringDao := fspDaoPack.NewLevel2AssuringDao()
	avr := assuringDao.GetAssuringValueRecordByCoilId(cmr.CoilId)

	//threshold := 50.0
	if tags.IsCoilIdValid(avr.CoilId) {
		// head thk
		s.Trigger(avr.HeadThkAssuringValue, 50.0, cmr, schemas.ThkUnqualCategory, schemas.HeadThkUnqualType)
		// body thk
		s.Trigger(avr.BodyThkAssuringValue, 80.0, cmr, schemas.ThkUnqualCategory, schemas.BodyThkUnqualType)
		// wid
		s.Trigger(avr.WidAssuringValue, 15.0, cmr, schemas.WidUnqualCategory, schemas.WidUnqualType)
		// crn
		if cmr.AimThick <= 6.0 {
			s.Trigger(avr.CrnAssuringValue, 15.0, cmr, schemas.CrnUnqualCategory, schemas.CrnUnqualType)
		}
		// fdt
		s.Trigger(avr.FdtAssuringValue, 65.0, cmr, schemas.FdtUnqualCategory, schemas.FdtUnqualType)
		// ct
		s.Trigger(avr.CtAssuringValue, 60.0, cmr, schemas.CtUnqualCategory, schemas.CtUnqualType)
		// flatness
		s.Trigger(avr.SymFltAssuringValue, 85.0, cmr, schemas.SymFltUnqualCategory, schemas.SymFltUnqualType)
		s.Trigger(avr.AsymFltAssuringValue, 35.0, cmr, schemas.AsymFltUnqualCategory, schemas.AsymFltUnqualType)
	}
}

func (s *UnqualLevel2AssuringService) Trigger(val float64, threshold float64, cmr *fspDaoPack.CoilMainRecord, unqualCategory schemas.UnqualIdiom, unqualType schemas.UnqualIdiom) {
	if val < threshold && val >= 0 {
		uhi := GetUnqualHistoryInfo(unqualCategory, schemas.Level2AssuringUnqualSource, unqualType, fmt.Sprintf("%.2f", val))
		SaveUnqualHistory(cmr, uhi)
	}
}




