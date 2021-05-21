package services

import (
	"log"
	"github.com/nekohor/mangosteen/internal/app/dbx"
	fspDaoPack "github.com/nekohor/mangosteen/internal/apps/fsp/dao"
	"github.com/nekohor/mangosteen/internal/apps/unqualified/models"
	"github.com/nekohor/mangosteen/internal/apps/unqualified/schemas"
	"github.com/nekohor/mangosteen/pkg/tags"
	"strings"
)

type UnqualInspectionService struct {
	
}

func NewUnqualInspectionService() *UnqualInspectionService {
	return &UnqualInspectionService{}
}

func (s *UnqualInspectionService) GetQualityInspectionAccountByCoilId(coilId string) *models.QualityInspectionAccount {
	db := dbx.GormDB("unqual")
	qia := &models.QualityInspectionAccount{}
	db.Where("coil_id = ?", coilId).First(&qia)
	return qia
}

func (s *UnqualInspectionService) SaveUnqualHistoriesByCoilId(coilId string) {
	coilMainDao := fspDaoPack.NewCoilMainDao()
	cmr := coilMainDao.GetRecordByCoilId(coilId)
	log.Println(cmr)
	if cmr.CoilId != "" {
		s.SaveUnqualHistoriesByCoilMainRecord(cmr)
	}
}

func (s *UnqualInspectionService) SaveUnqualHistoriesByCoilMainRecord(cmr *fspDaoPack.CoilMainRecord) {
	qia := s.GetQualityInspectionAccountByCoilId(cmr.CoilId)

	if !tags.IsCoilIdValid(qia.CoilId) {
		return
	}
	// 卷形
	if qia.CoilId != "" && qia.CoilDefect != "" &&  qia.CoilDefectDesc != "" && !strings.Contains(qia.BlockState, "合格") {
		s.TriggerCoilShape(cmr, qia)
	}
	// 轧破甩尾
	if qia.CoilId != "" && (strings.Contains(qia.SurfaceDefect, "轧破") || strings.Contains(qia.CoilDefect, "烂尾")) {
		s.TriggerRollBreak(cmr, qia)
	}
}

func (s *UnqualInspectionService) TriggerCoilShape(cmr *fspDaoPack.CoilMainRecord, qia *models.QualityInspectionAccount) {
	coilDefectTags := []string{
		"溢出",
		"鱼尾",
		"折叠",
		"松卷",
		"塌卷",
		"折边",
		"交错",
		"错层",
		"塔形",
		"扁卷",
		"喷印不清",
		"夹损",
		"侧喷不清",
		"边损",
		"划伤",
		"辊印",
	}
	var coilDefect string
	for _, coilDefectTag := range coilDefectTags {
		if strings.Contains(qia.CoilDefect, coilDefectTag) {
			if coilDefect == "" {
				coilDefect = coilDefect + coilDefectTag
			} else {
				coilDefect = coilDefect + "|" + coilDefectTag
			}
		}
	}
	if coilDefect != "" {
		coilDefectValue := qia.CoilDefect + "|" + qia.CoilDefectDesc
		uhi := GetUnqualHistoryInfo(schemas.CoilShapeUnqualCategory, schemas.InspectionUnqualSource, schemas.NewUnqualType(coilDefect), coilDefectValue)
		SaveUnqualHistory(cmr, uhi)
	}
}

func (s *UnqualInspectionService) TriggerRollBreak(cmr *fspDaoPack.CoilMainRecord, qia *models.QualityInspectionAccount) {
	//rollBreakDefectTags := []string{
	//	"轧破",
	//	"烂尾",
	//}
	var rollBreakDefect string
	var rollBreakDefectValue string
	if strings.Contains(qia.SurfaceDefect, "轧破") {
		rollBreakDefect = rollBreakDefect + "轧破"
		rollBreakDefectValue = rollBreakDefectValue  + qia.SurfaceDefectDesc
	}
	if strings.Contains(qia.CoilDefect, "烂尾") {
		if rollBreakDefect == "" {
			rollBreakDefect = rollBreakDefect + "烂尾"
			rollBreakDefectValue = rollBreakDefectValue + qia.CoilDefectDesc
		} else {
			rollBreakDefect = rollBreakDefect + "|" + "烂尾"
			rollBreakDefectValue = rollBreakDefectValue + "|" + qia.CoilDefectDesc
		}

	}
	if rollBreakDefect != "" {
		uhi := GetUnqualHistoryInfo(schemas.RollBreakUnqualCategory, schemas.InspectionUnqualSource, schemas.NewUnqualType(rollBreakDefect), rollBreakDefectValue)
		SaveUnqualHistory(cmr, uhi)
	}
}