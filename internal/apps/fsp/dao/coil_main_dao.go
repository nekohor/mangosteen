package dao

import (
	"fmt"
	"log"
	"github.com/nekohor/mangosteen/internal/apps/fsp/ibmdbx"
)

type CoilMainRecord struct {
	CoilId string `db:"L3BATCH"`
	ProductTime string `db:"PRODUCTIONTIME"`
	SteelGrade string `db:"STEELGRADE"`
	AimThick float64 `db:"THICKNESS"`
	AimWidth int `db:"WIDTH"`
	Weight float64 `db:"WEIGHT"`
	UsageCode string `db:"USAGE_CODE"`
}

type CoilMainDao struct {
}

func NewCoilMainDao() *CoilMainDao {
	d := &CoilMainDao{}
	return d
}


func (d *CoilMainDao) getMillLineCode(millLineTag string) string {
	var code string
	if millLineTag == "HSM1" {
		code = millLineTag
	} else if millLineTag == "HSM2" {
		code = millLineTag
	} else if millLineTag == "MG2250" {
		code = "HSM1"
	} else if millLineTag == "MG1580" {
		code = "HSM2"
	} else {
		code = "HSM1"
	}
	return code
}

func (d *CoilMainDao) getSelectedColumns() string {
	return "L3BATCH, PRODUCTIONTIME, STEELGRADE, THICKNESS, WIDTH, WEIGHT, USAGE_CODE"
}

func (d *CoilMainDao) getTableName() string {
	return "TM.K_M_COIL_MAIN"
}

func (d *CoilMainDao) GetRecordsByDate(queryDate string, millLineTag string) []*CoilMainRecord {
	var err error

	db := ibmdbx.DB("ap")
	curDate := queryDate[:4] + queryDate[5:7] + queryDate[8:]
	sqlTemplate := "SELECT %s FROM %s where PRODUCTIONTIME >= '%s 00:00:00' and  PRODUCTIONTIME <= '%s 23:59:59' AND LINECODE = '%s'"
	sqlStatement := fmt.Sprintf(sqlTemplate, d.getSelectedColumns(),d.getTableName(), curDate, curDate, d.getMillLineCode(millLineTag))

	records := make([]*CoilMainRecord, 0)
	err = db.Select(&records, sqlStatement)
	if err != nil {
		log.Println(err)
	}
	return records
}

func (d *CoilMainDao) GetRecordByCoilId(coilId string) *CoilMainRecord {
	var err error
	db := ibmdbx.DB("ap")
	sqlTemplate := "SELECT %s FROM %s where L3BATCH = '%s'"
	sqlStatement := fmt.Sprintf(sqlTemplate, d.getSelectedColumns(),d.getTableName(), coilId)
	record := &CoilMainRecord{}
	err = db.Get(record, sqlStatement)
	if err != nil {
		log.Println(err)
	}
	return record
}