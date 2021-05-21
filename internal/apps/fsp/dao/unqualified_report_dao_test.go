package dao

import (
	"log"
	"testing"
)

func TestUnqualifiedReportDao_GetRecordsByCoilId(t *testing.T) {
	d := NewUnqualifiedReportDao()
	records := d.GetRecordsByCoilId("H21010780L")
	log.Println(records[0])
}