package gradeklass

import (
	"strings"
)


func (k *GradeKlass) IsAgreementProduct(steelGrade string) bool {
	if strings.HasSuffix(strings.ToUpper(steelGrade), "-X") {
		return true
	} else {
		return false
	}
}

func (k *GradeKlass) IsPicklePlate(steelGrade string) bool {
	if k.IsSpecialPicklePlate(steelGrade) {
		return true
	}

	if k.IsHighSurfacePicklePlate(steelGrade) {
		return true
	}

	suffixes := []string{"-P", "-P2", "-PO"}
	if k.IsInSuffixes(steelGrade, suffixes) {
		return true
	}
	return false
}

func (k *GradeKlass) IsSpecialPicklePlate(steelGrade string) bool {
	specialNames := []string{"MCFC", "S355J2WP"}
	if Contain(specialNames, strings.ToUpper(steelGrade)) {
		return true
	}
	if k.IsPrefix(steelGrade, "HR") && k.IsSuffix(steelGrade, "LA") {
		return true
	}
	if k.IsPrefix(steelGrade, "HD") && k.IsSuffix(steelGrade, "LA") {
		return true
	}
	if k.IsPrefix(steelGrade, "GH") && k.IsSuffix(steelGrade, "MC") {
		return true
	}
	headers := []string{"CP", "HR", "GAH"}
	if k.IsGradeNumberWithHeaders(steelGrade, headers, 1,">",0) {
		return true
	}
	return false
}

func (k *GradeKlass) IsHighSurfacePicklePlate(steelGrade string) bool {
	suffixes := []string{
		"-MZ", "-MZ1",
		"-YS", "-YS1", "-YS2", "-YS3",
		"-EM", "-EM1", "-EM2", "-EM3", "-EM4",
		"-DR", "-DR1",
		"-RL", "-RL1",
		"-DB", "-DB1", "-DB2", "-DB3",
		"-EPS", "-BZ", "-HY1", "-GL",
	}
	if k.IsInSuffixes(steelGrade, suffixes) {
		return true
	}
	return false
}

func (k *GradeKlass) IsHighStrengthPicklePlate(steelGrade string) bool {
	var headers []string
	if !k.IsPicklePlate(steelGrade) {
		return false
	}

	if k.IsSpecialPicklePlate(steelGrade) {
		return false
	}

	// MnB
	if strings.Contains(strings.ToUpper(steelGrade), "MNB") {
		return true
	}

	// number more than (2)
	headers = []string{"S", "SS", "QSTE"}
	if k.IsGradeNumberWithHeaders(steelGrade, headers, 1,">", 2) {
		return true
	}

	// number more than (3)
	headers = []string{"Q", "MTC", "MZG", "MJZ", "MHG", "MH", "HR", "A", "SAPH", "SG", "SPFH"}
	if k.IsGradeNumberWithHeaders(steelGrade, headers, 1,">",3) {
		return true
	}
	return false
}

func (k *GradeKlass) IsSilicon(steelGrade string) bool {
	if k.IsOrientedSilicon(steelGrade) {
		return true
	}
	if k.IsNonOrientedSilicon(steelGrade) {
		return true
	}
	return false
}

func (k *GradeKlass) GetOrientedSiliconPrefixes() []string {
	prefixes := []string{"MBRNQ", "MCGO"}
	return prefixes
}

func (k *GradeKlass) IsOrientedSilicon(steelGrade string) bool {
	prefixes := k.GetOrientedSiliconPrefixes()
	if k.IsInPrefixes(steelGrade, prefixes) {
		return true
	}
	return false
}

func (k *GradeKlass) GetNonOrientedSiliconPrefixes() []string {
	prefixes := []string{"MBRW", "MGW"}
	return prefixes
}

func (k *GradeKlass) IsNonOrientedSilicon(steelGrade string) bool {
	prefixes := k.GetNonOrientedSiliconPrefixes()
	if k.IsInPrefixes(steelGrade, prefixes) {
		return true
	}
	return false
}

func (k *GradeKlass) IsHighGradeNonOrientedSilicon(steelGrade string) bool {
	if k.IsNonOrientedSilicon(steelGrade) {
		if k.IsMidLowGradeNonOrientedSilicon(steelGrade) {
			return false
		} else {
			return true
		}
	}
	return false
}

func (k *GradeKlass) IsMidLowGradeNonOrientedSilicon(steelGrade string) bool {
	if k.IsNonOrientedSilicon(steelGrade) {
		prefixes := k.GetNonOrientedSiliconPrefixes()
		if k.IsGradeNumberEqual(steelGrade, prefixes, 4,1300) {
			return true
		}
		if k.IsGradeNumberMoreThan(steelGrade, prefixes, 3,350) {
			return true
		}
	}
	return false
}

func (k *GradeKlass) IsColdBase(steelGrade string) bool {
	prefixes := []string{"MR", "MB"}
	if k.IsInPrefixes(steelGrade, prefixes) {
		return true
	} else {
		return false
	}
}

func (k *GradeKlass) IsHighStrengthColdBase(steelGrade string) bool {
	if k.IsColdBase(steelGrade) {
		tailGradeName := Substr(strings.ToUpper(steelGrade), 5, 3)
		if k.IsGradeNumberMoreThanOrEqual(tailGradeName, []string{""}, 3,290) {
			return true
		} else {
			return false
		}
	}
	return false
}

func (k *GradeKlass) IsNormalStrengthColdBase(steelGrade string) bool {
	if k.IsColdBase(steelGrade) {
		if k.IsHighStrengthColdBase(steelGrade) {
			return false
		} else {
			return true
		}
	}
	return false
}

func (k *GradeKlass) IsHighStrengthCommercialSteel(steelGrade string) bool {
	var headers []string

	if k.IsPicklePlate(steelGrade) {
		return false
	}
	// MnB
	if strings.Contains(strings.ToUpper(steelGrade), "MNB") && !k.IsSuffix(steelGrade, "-P") {
		return true
	}

	if k.IsAxleSteel(steelGrade) {
		return true
	}

	if k.IsBeamSteel(steelGrade) {
		return true
	}

	if k.IsWheelSteel(steelGrade) {
		return true
	}

	if k.IsPipelineSteel(steelGrade) {
		return true
	}

	if k.IsMechanicalMixingSteel(steelGrade) {
		return true
	}

	headers = []string{"X", "GRADE", "H"}
	if k.IsGradeNumberWithHeaders(steelGrade, headers, 2,">",30) {
		return true
	}

	headers = []string{"Q", "A", "ASTM A", "CP", "DL", "GAH", "GH", "L", "HR", "M", "MH", "MQ", "MYS", "QSTE", "S", "SAPH", "SG", "SPFH", "MTC"}
	if k.IsGradeNumberWithHeaders(steelGrade, headers, 3,">",300) {
		return true
	}

	headers = []string{"M"}
	if k.IsGradeNumberWithHeaders(steelGrade, headers, 4,">",1000) {
		return true
	}
	return false
}
