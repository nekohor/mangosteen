package dao

import (
	"fmt"
	"log"
	"github.com/nekohor/mangosteen/internal/apps/fsp/ibmdbx"
)

type Level2AssuringDao struct {

}

func NewLevel2AssuringDao() *Level2AssuringDao {
	return &Level2AssuringDao{}
}

type AssuringValueRecord struct {
	CoilId       string  `db:"COILID"`
	HeadThkAssuringValue float64 `db:"THK_H_ENSURE_VALUE"`
	BodyThkAssuringValue float64 `db:"THK_B_ENSURE_VALUE"`
	WidAssuringValue float64 `db:"WIDTH_ENSURE_VALUE_10_17"`
	CrnAssuringValue float64 `db:"C40_ENSURE_VALUE_0_150"`
	FdtAssuringValue float64 `db:"FDT1_TEMP_ENSURE_VALUE"`
	CtAssuringValue float64 `db:"CT1_ENSURE_VALUE"`
	SymFltAssuringValue float64 `db:"SYM_BEND_ENSURE_VALUE"`
	AsymFltAssuringValue float64 `db:"ASYM_LEVEL_ENSURE_VALUE"`
}

func (d *Level2AssuringDao) getSelectedColumns() []string {
	cols := []string{
		"COILID",
		"THK_H_ENSURE_VALUE",
		"THK_B_ENSURE_VALUE",
		"WIDTH_ENSURE_VALUE_10_17",
		"C40_ENSURE_VALUE_0_150",
		"FDT1_TEMP_ENSURE_VALUE",
		"CT1_ENSURE_VALUE",
		"SYM_BEND_ENSURE_VALUE",
		"ASYM_LEVEL_ENSURE_VALUE",
	}
	return cols
}

func (d *Level2AssuringDao) getSelectedColumnsString() string {
	var columnsString string
	for i, col := range d.getSelectedColumns(){
		if i == 0 {
			columnsString += col
		} else {
			columnsString += fmt.Sprintf(", COALESCE(%s, -1) %s", col, col)
		}
	}
	return columnsString
}

func (d *Level2AssuringDao) getTableName() string {
	return "DC.ENSURE_VALUE"
}

func (d *Level2AssuringDao) GetAssuringValueRecordByCoilId(coilId string) *AssuringValueRecord {
	var err error
	db := ibmdbx.DB("ap")
	sqlTemplate := "SELECT %s FROM %s where COILID = '%s'"
	sqlStatement := fmt.Sprintf(sqlTemplate, d.getSelectedColumnsString(), d.getTableName(), coilId)
	record := &AssuringValueRecord{}
	err = db.Get(record, sqlStatement)
	if err != nil {
		log.Println(coilId)
		log.Println(sqlStatement)
		log.Println(err)
	}
	log.Println(record)
	return record
}
