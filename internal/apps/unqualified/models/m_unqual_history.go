package models

import (
	"github.com/nekohor/mangosteen/internal/app/model"
	"time"
)

type UnqualHistory struct {
	model.BaseModel
	CoilId         string
	ProductTime    *time.Time
	SteelGrade     string
	AimThk         float64
	AimWid         int
	UnqualCategory string
	UnqualSource   string
	UnqualType     string
	UnqualValue    string

}


func (*UnqualHistory) TableName() string {
	return "unqual_histories"
}
