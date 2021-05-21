package dao

import (
	"log"
	"testing"
)

func TestLevel2AssuringDao_GetAssuringValueRecordByCoilId(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	d := NewLevel2AssuringDao()
	_ = d.GetAssuringValueRecordByCoilId("H210699330")
}