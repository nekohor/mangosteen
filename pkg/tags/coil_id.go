package tags

import "strings"

func IsCoilIdValid(coilId string) bool {
	if strings.HasPrefix(coilId, "H") {
		return true
	}
	if strings.HasPrefix(coilId, "M") {
		return true
	}
	return false
}


func IsMG2250(coilId string) bool {
	if strings.HasPrefix(strings.ToUpper(coilId), "H") {
		return true
	}
	return false
}

func IsMG1580(coilId string) bool {
	if strings.HasPrefix(strings.ToUpper(coilId), "M") {
		return true
	}
	return false
}