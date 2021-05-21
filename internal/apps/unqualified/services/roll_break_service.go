package services

import (
	"github.com/nekohor/mangosteen/internal/app/dbx"
	"github.com/nekohor/mangosteen/internal/apps/unqualified/models"
	"github.com/nekohor/mangosteen/pkg/tags"
)

type RollBreakStatService struct {
}

func NewRollBreakStatService() *RollBreakStatService {
	return &RollBreakStatService{}
}

type RollBreakStatResult struct {
	RollBreakCountTotal  int `json:"rollBreakCountTotal"`
	RollBreakCountMG2250 int `json:"rollBreakCountMG2250"`
	RollBreakCountMG1580 int `json:"rollBreakCountMG1580"`
	HeadBreakCount       int `json:"headBreakCount"`
	TailBreakCount       int `json:"tailBreakCount"`
}

func (*RollBreakStatService) GetRollBreakStatResult(queryDate string) *RollBreakStatResult {
	startDate := queryDate + " 00:00:00"
	endDate := queryDate + " 23:59:59"
	db := dbx.GormDB("unqual")
	histories := make([]*models.UnqualHistory, 0)
	db.Where("product_time > ?", startDate).
		Where("product_time < ?", endDate).
		Where("unqual_category = ?", "轧破甩尾").
		Find(&histories)
	res := &RollBreakStatResult{}
	res.RollBreakCountTotal = len(histories)

	for _, history := range histories {
		if tags.IsMG2250(history.CoilId) {
			res.RollBreakCountMG2250 += 1
		}
		if tags.IsMG1580(history.CoilId) {
			res.RollBreakCountMG1580 += 1
		}
	}
	return res
}
