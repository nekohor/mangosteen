package services

import (
	"github.com/nekohor/mangosteen/internal/apps/category/gradeklass"
)
// steel grade family service
type GradeKlassService struct {
}

func (s *GradeKlassService) GetClassificationResultBySteelGrade(steelGrade string) *gradeklass.SteelGradeClassificationResult {
	var res *gradeklass.SteelGradeClassificationResult
	var k *gradeklass.GradeKlass
	k = &gradeklass.GradeKlass{}
	res = k.GetClassificationResultBySteelGrade(steelGrade)
	return res
}