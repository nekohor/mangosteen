package models

import (
	"github.com/nekohor/mangosteen/internal/app/model"
	"time"
)

type UnqualReason struct {
	model.BaseModel
	CoilId         string
	ProductTime    *time.Time
	SteelGrade     string
	AimThk         float64
	AimWid         int
	UnqualCategory string
	HighAccuracyType string
	InternalControlType string
	DeliveryControlType string
	UnqualValues    string
	TechnicianReason string
	FspReason string
	TmeicReason string
	RootReason string
	RootReasonType string
}


func (*UnqualReason) TableName() string {
	return "unqual_reasons"
}
