package services

import (
	"github.com/nekohor/mangosteen/internal/app/dbx"
	"github.com/nekohor/mangosteen/internal/apps/unqualified/models"
	"github.com/nekohor/mangosteen/internal/apps/unqualified/schemas"
)

type UnqualReasonService struct {
}

func NewUnqualReasonService() *UnqualReasonService {
	return &UnqualReasonService{}
}

func (s *UnqualReasonService) GetUnqualHistoriesByDate(queryDate string) []*models.UnqualHistory {
	startDate := queryDate + " 00:00:00"
	endDate := queryDate + " 23:59:59"
	db := dbx.GormDB("unqual")
	histories := make([]*models.UnqualHistory, 0)
	db.Where("product_time > ?", startDate).
		Where("product_time < ?", endDate).
		Find(&histories)
	return histories
}

func (s *UnqualReasonService) GetUnqualCategories() []schemas.UnqualIdiom {
	categories := []schemas.UnqualIdiom{
		schemas.ThkUnqualCategory,
		schemas.WidUnqualCategory,
		schemas.FdtUnqualCategory,
		schemas.CtUnqualCategory,
		schemas.WdgUnqualCategory,
		schemas.CrnUnqualCategory,
		schemas.SymFltUnqualCategory,
		schemas.AsymFltUnqualCategory,
		schemas.CoilShapeUnqualCategory,
		schemas.RollBreakUnqualCategory,
	}
	return categories
}

func (s *UnqualReasonService) GetUnqualHistoriesByUnqualCategory(unqualHistories []*models.UnqualHistory, unqualCategory schemas.UnqualIdiom) []*models.UnqualHistory {
	selected := make([]*models.UnqualHistory, 0)
	for _, unqualHistory := range unqualHistories {
		if unqualHistory.UnqualCategory == unqualCategory.Name() {
			selected = append(selected, unqualHistory)
		}
	}
	return selected
}

func (s *UnqualReasonService) GetUnqualHistoriesByCoilId(unqualHistories []*models.UnqualHistory, coilId string) []*models.UnqualHistory {
	selected := make([]*models.UnqualHistory, 0)
	for _, unqualHistory := range unqualHistories {
		if unqualHistory.CoilId == coilId {
			selected = append(selected, unqualHistory)
		}
	}
	return selected
}

func (s *UnqualReasonService) GetCoilIdsFromUnqualHistories(unqualHistories []*models.UnqualHistory) []string {
	coilIds := make([]string, 0)
	for i, _ := range unqualHistories {
		flag := true
		for j, _ := range coilIds {
			if unqualHistories[i].CoilId == coilIds[j] {
				flag = false  // 存在重复元素，标识为false
				break
			}
		}
		if flag {  // 标识为false，不添加进结果
			coilIds = append(coilIds, unqualHistories[i].CoilId)
		}
	}
	return coilIds
}

func (s *UnqualReasonService) SaveUnqualReasonsByDate(queryDate string) {
	histories := s.GetUnqualHistoriesByDate(queryDate)
	coilIds := s.GetCoilIdsFromUnqualHistories(histories)
	if len(histories) == 0 {
		return
	}
	for _, unqualCategory := range s.GetUnqualCategories() {
		historiesWithSameCategory := s.GetUnqualHistoriesByUnqualCategory(histories, unqualCategory)
		for _, coilId := range coilIds {
			historiesWithSameCategoryAndCoilId := s.GetUnqualHistoriesByCoilId(historiesWithSameCategory, coilId)
		}
	}
}

func (s *UnqualReasonService) GetUnqualReasonByUnqualHistories(historiesWithSameCategoryAndCoilId []*models.UnqualHistory) {

}
