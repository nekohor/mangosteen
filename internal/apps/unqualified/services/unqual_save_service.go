package services

import (
	"log"
	fspDaoPack "github.com/nekohor/mangosteen/internal/apps/fsp/dao"
	"github.com/nekohor/mangosteen/pkg/tags"
)
type UnqualSaveService struct {
	
}

func NewUnqualSaveService() *UnqualSaveService {
	return &UnqualSaveService{}
}
func (s *UnqualSaveService) SaveUnqualHistoriesByDates(startDate string, endDate string) {

}

func (s *UnqualSaveService) SaveUnqualHistoriesByDate(currentDate string, millLine string) {
	d := fspDaoPack.NewCoilMainDao()
	records := d.GetRecordsByDate(currentDate, millLine)
	log.Println(len(records))
	for _, record := range records {
		if tags.IsCoilIdValid(record.CoilId) {
			s.SaveUnqualHistoriesByCoilMainRecord(record)
		}
	}
}

func (s *UnqualSaveService) SaveUnqualHistoriesByCoilId(coilId string) {

	l2assuringService := NewUnqualifiedLevel2AssuringService()
	l2assuringService.SaveUnqualHistoriesByCoilId(coilId)

	unqualFspService := NewUnqualifiedFspService()
	unqualFspService.SaveUnqualHistoriesByCoilId(coilId)

	unqualInspectionService := NewUnqualInspectionService()
	unqualInspectionService.SaveUnqualHistoriesByCoilId(coilId)
}

func (s *UnqualSaveService) SaveUnqualHistoriesByCoilMainRecord(cmr *fspDaoPack.CoilMainRecord) {
	l2assuringService := NewUnqualifiedLevel2AssuringService()
	l2assuringService.SaveUnqualHistoriesByCoilMainRecord(cmr)

	unqualFspService := NewUnqualifiedFspService()
	unqualFspService.SaveUnqualHistoriesByCoilMainRecord(cmr)

	unqualInspectionService := NewUnqualInspectionService()
	unqualInspectionService.SaveUnqualHistoriesByCoilMainRecord(cmr)
}