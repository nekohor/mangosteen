package gradeklass

import (
	"strconv"
	"strings"
)

func (k *GradeKlass) IsTrueByOperation(numLeft int64, operation string, numRight int64) bool {
	if operation == ">" {
		return numLeft > numRight
	} else if operation == ">=" {
		return numLeft >= numRight
	} else if operation == "==" || operation == "=" {
		return numLeft == numRight
	} else if operation == "<=" {
		return numLeft <= numRight
	} else if operation == "<" {
		return numLeft < numRight
	} else {
		return false
	}
}

func (k *GradeKlass) IsGradeNumber(steelGrade string, startIndex int, numberLength int, operation string, numberBound int) bool {
	numberString := Substr(strings.ToUpper(steelGrade), startIndex, numberLength)
	number, err := strconv.ParseInt(numberString, 10, 0)
	if err == nil {
		if k.IsTrueByOperation(number, operation, int64(numberBound))  {
			return true
		}
	}
	return false
}

func (k *GradeKlass) IsGradeNumberWithHeader(steelGrade string, header string, numberLength int, operation string, numberBound int) bool {
	if k.IsGradeNumber(steelGrade, len(header), numberLength, operation, numberBound) {
		return true
	}
	return false
}

func (k *GradeKlass) IsGradeNumberWithHeaders(steelGrade string, headers []string, numberLength int, operation string, numberBound int) bool {
	for _, header := range headers {
		if k.IsGradeNumberWithHeader(steelGrade, header, numberLength, operation, numberBound) {
			return true
		}
	}
	return false
}

func (k *GradeKlass) IsGradeNumberMoreThanOrEqual(steelGrade string, patterns []string, numberLength int, numberBound int) bool {
	for _, pattern := range patterns {
		if k.IsGradeNumberWithHeader(steelGrade, pattern, numberLength, ">=", numberBound) {
			return true
		}
	}
	return false
}

func (k *GradeKlass) IsGradeNumberMoreThan(steelGrade string, patterns []string, numberLength int, numberBound int) bool {
	for _, pattern := range patterns {
		if k.IsGradeNumberWithHeader(steelGrade, pattern, numberLength, ">", numberBound) {
			return true
		}
	}
	return false
}

func (k *GradeKlass) IsGradeNumberEqual(steelGrade string, patterns []string, numberLength int, numberBound int) bool {
	for _, pattern := range patterns {
		if k.IsGradeNumberWithHeader(steelGrade, pattern, numberLength, "==", numberBound) {
			return true
		}
	}
	return false
}

func (k *GradeKlass) IsPrefix(steelGrade string, prefix string) bool {
	if strings.HasPrefix(strings.ToUpper(steelGrade), prefix) {
		return true
	}
	return false
}

func (k *GradeKlass) IsSuffix(steelGrade string, suffix string) bool {
	if strings.HasSuffix(strings.ToUpper(steelGrade), suffix) {
		return true
	}
	return false
}

func (k *GradeKlass) IsInPrefixes(steelGrade string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(strings.ToUpper(steelGrade), prefix) {
			return true
		}
	}
	return false
}

func (k *GradeKlass) IsInSuffixes(steelGrade string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(strings.ToUpper(steelGrade), suffix) {
			return true
		}
	}
	return false
}