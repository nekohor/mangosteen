package dao

import (
	"log"
	"testing"
)

func TestCoilMainDao_getRecordsByDate(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	d := NewCoilMainDao()
	records := d.GetRecordsByDate("2021-05-07", "HSM1")
	log.Println(records[0])
}

