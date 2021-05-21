package gradeklass

import "strings"

func (k *GradeKlass) IsAxleSteel(steelGrade string) bool {
	suffixes := []string{"QK"}
	if k.IsInSuffixes(steelGrade, suffixes) {
		return true
	}
	return false
}


func (k *GradeKlass) IsBeamSteel(steelGrade string) bool {

	if k.IsSuffix(steelGrade, "-P") {
		return false
	}

	if k.IsSuffix(steelGrade, "L") {
		if k.IsGradeNumberMoreThanOrEqual(steelGrade, []string{""}, 3, 300) {
			return true
		}
	}

	if k.IsPrefix(steelGrade, "M") &&
		Substr(strings.ToUpper(steelGrade),4,1) ==  "L" &&
		!k.IsSuffix(steelGrade, "CL") &&
		!k.IsSuffix(steelGrade, "-P") {
		if k.IsGradeNumberMoreThanOrEqual(steelGrade, []string{"M"}, 3, 300) {
			return true
		}
	}

	if k.IsPrefix(steelGrade, "DL") {
		if k.IsGradeNumberMoreThanOrEqual(steelGrade, []string{"DL"}, 3, 300) {
			return true
		}
	}

	if k.IsPrefix(steelGrade, "QSTE") && k.IsSuffix(steelGrade, "TM") {
		if k.IsGradeNumberMoreThanOrEqual(steelGrade, []string{"QSTE"}, 3, 600) {
			return true
		}
	}

	if k.IsPrefix(steelGrade, "S") && Substr(strings.ToUpper(steelGrade),4,2) ==  "MC" && !k.IsSuffix(steelGrade, "-P") {
		if k.IsGradeNumberMoreThanOrEqual(steelGrade, []string{"S"}, 3, 650) {
			return true
		}
	}
	return false
}

func (k *GradeKlass) IsWheelSteel(steelGrade string) bool {
	if k.IsPrefix(steelGrade, "M") && k.IsSuffix(steelGrade, "CL") {
			return true
	}
	return false
}

func (k *GradeKlass) IsPipelineSteel(steelGrade string) bool {
	if k.IsPrefix(steelGrade, "X") {
		if k.IsGradeNumberMoreThanOrEqual(steelGrade, []string{"X"}, 2, 10) {
			return true
		}
	}

	if k.IsPrefix(steelGrade, "L") && Substr(strings.ToUpper(steelGrade),4,1) == "M" {
		if k.IsGradeNumberMoreThanOrEqual(steelGrade, []string{"L"}, 3, 100) {
			return true
		}
	}
	return false
}

func (k *GradeKlass) IsMechanicalMixingSteel(steelGrade string) bool {
	if k.IsPrefix(steelGrade, "M") && k.IsSuffix(steelGrade, "JJ") {
		if k.IsGradeNumberMoreThanOrEqual(steelGrade, []string{"M"}, 2, 100) {
			return true
		}
	}
	return false
}

