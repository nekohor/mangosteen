package services

import (
	"log"
	fspDaoPack "github.com/nekohor/mangosteen/internal/apps/fsp/dao"
	"github.com/nekohor/mangosteen/internal/apps/unqualified/schemas"
	"github.com/nekohor/mangosteen/pkg/logger/zap/logger"
	"github.com/nekohor/mangosteen/pkg/tags"
	"github.com/nekohor/mangosteen/pkg/util/transcode"
	"strings"
)
type UnqualFspService struct {
}

func NewUnqualifiedFspService() *UnqualFspService {
	return &UnqualFspService{}
}

func (s *UnqualFspService) SaveUnqualHistoriesByCoilId(coilId string) {
	coilMainDao := fspDaoPack.NewCoilMainDao()
	cmr := coilMainDao.GetRecordByCoilId(coilId)
	log.Println(cmr)
	if cmr.CoilId != "" {
		s.SaveUnqualHistoriesByCoilMainRecord(cmr)
	}
}

func (s *UnqualFspService) SaveUnqualHistoriesByCoilMainRecord(cmr *fspDaoPack.CoilMainRecord) {
	quDao := fspDaoPack.NewUnqualifiedReportDao()
	urrs := quDao.GetRecordsByCoilId(cmr.CoilId)
	if len(urrs) == 0 {
		return
	}
	for _, urr := range urrs {
		if tags.IsCoilIdValid(urr.CoilId) {
			s.Trigger(cmr, urr, schemas.ThkFspUnqualCategory, schemas.ThkUnqualCategory, schemas.ThkUnqualType)
			s.Trigger(cmr, urr, schemas.WidFspUnqualCategory, schemas.WidUnqualCategory, schemas.WidUnqualType)
			s.Trigger(cmr, urr, schemas.FdtFspUnqualCategory, schemas.FdtUnqualCategory, schemas.FdtUnqualType)
			s.Trigger(cmr, urr, schemas.CtFspUnqualCategory, schemas.CtUnqualCategory, schemas.CtUnqualType)
			s.Trigger(cmr, urr, schemas.WdgFspUnqualCategory, schemas.WdgUnqualCategory, schemas.WdgUnqualType)
			s.Trigger(cmr, urr, schemas.CrnFspUnqualCategory, schemas.CrnUnqualCategory, schemas.CrnUnqualType)
			s.Trigger(cmr, urr, schemas.SymFltFspUnqualCategory, schemas.SymFltUnqualCategory, schemas.SymFltUnqualType)
			s.Trigger(cmr, urr, schemas.AsymFltFspUnqualCategory, schemas.AsymFltUnqualCategory, schemas.AsymFltUnqualType)
		}
	}
}

func (s *UnqualFspService) Trigger(cmr *fspDaoPack.CoilMainRecord, urr *fspDaoPack.UnqualifiedReportRecord, fspUnqualCategory schemas.UnqualIdiom, unqualCategory schemas.UnqualIdiom, unqualType schemas.UnqualIdiom) {
	if strings.TrimSpace(s.TransGbkToUtf8(urr.UnqualifiedContent)) == fspUnqualCategory.Name() {
		uhi := GetUnqualHistoryInfo(unqualCategory, schemas.FspUnqualSource, unqualType, s.TransGbkToUtf8(urr.UnqualifiedReason))
		SaveUnqualHistory(cmr, uhi)
	}
}

func (s *UnqualFspService) TransGbkToUtf8(content string) string {
	var transferred []byte
	var err error
	transferred, err = transcode.GbkToUtf8([]byte(content))
	if err != nil {
		logger.Error(err)
	}
	return string(transferred)
}

