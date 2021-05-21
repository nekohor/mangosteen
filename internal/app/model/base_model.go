package model

import (
	"time"
)
type BaseModel struct {
	Id        uint `gorm:"primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (m *BaseModel) AddCreateTime() {
	c := time.Now()
	m.CreatedAt = &c
	m.AddUpdateTime()
}

func (m *BaseModel) AddUpdateTime() {
	c := time.Now()
	m.UpdatedAt = &c
}