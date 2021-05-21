package dao

import (
	"fmt"
	"log"
	"github.com/nekohor/mangosteen/internal/apps/fsp/ibmdbx"
)

type UnqualifiedReportDao struct {
	
}

func NewUnqualifiedReportDao() *UnqualifiedReportDao {
	return &UnqualifiedReportDao{}
}

type UnqualifiedReportRecord struct {
	CoilId string `db:"MAT_ID"`
	UnqualifiedContent string `db:"CONTENT"`
	UnqualifiedReason string `db:"RESULT_REASON"`
}

func (d *UnqualifiedReportDao) getSelectedColumns() string {
	return "MAT_ID, CONTENT, RESULT_REASON"
}

func (d *UnqualifiedReportDao) getTableName() string {
	return "QU.TB_UNQUALIFIED_REPORT"
}

func (d *UnqualifiedReportDao) GetRecordsByCoilId(coilId string) []*UnqualifiedReportRecord {
	var err error
	db := ibmdbx.DB("qu")

	sqlTemplate := "SELECT %s FROM %s where MAT_ID = '%s'"
	sqlStatement := fmt.Sprintf(sqlTemplate, d.getSelectedColumns(), d.getTableName(), coilId)

	records := make([]*UnqualifiedReportRecord, 0)
	err = db.Select(&records, sqlStatement)
	if err != nil {
		log.Println(err)
	}
	if len(records) > 0 {
		log.Println(records[0])
	}
	return records
}