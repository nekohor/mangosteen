package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekohor/mangosteen/internal/app/ginx"
	"github.com/nekohor/mangosteen/internal/apps/fsp/dao"
)

type CoilMainController struct {

}

type CoilMainRecordRequest struct {
	CurrentDate string `json:"currentDate"`
	MillLineTag string `json:"millLineTag"`
}

type CoilMainRecordRespond struct {
	Records []*CoilMainRecord `json:"records"`
}

type CoilMainRecord struct {
	CoilId string `json:"coilId"`
	ProductTime string `json:"productTime"`
	SteelGrade string `json:"steelGrade"`
	AimThick float64 `json:"aimThick"`
	AimWidth int `json:"aimWidth"`
	Weight float64 `json:"weight"`
	UsageCode string `json:"usageCode"`
}

func (ctrl *CoilMainController) TransToCoilMainRecords(records []*dao.CoilMainRecord) []*CoilMainRecord {
	recs := make([]*CoilMainRecord, 0)
	if len(records) == 0 {
	} else {
		for _, record := range records {
			rec := &CoilMainRecord{}
			rec.CoilId =  record.CoilId
			rec.ProductTime =  record.ProductTime
			rec.SteelGrade =  record.SteelGrade
			rec.AimThick =  record.AimThick
			rec.AimWidth =  record.AimWidth
			rec.Weight =  record.Weight
			rec.UsageCode =  record.UsageCode

			recs = append(recs, rec)
		}
	}
	return recs
}

func (ctrl *CoilMainController) GetCoilMainRecordsByDate(c *gin.Context) {
	var err error
	var req CoilMainRecordRequest
	var res CoilMainRecordRespond

	err = ginx.ParseJSON(c, &req)
	if err != nil {
		ginx.ResSuccess(c, "亲，您的请求参数无法绑定")
	} else {
		d := dao.NewCoilMainDao()
		records := d.GetRecordsByDate(req.CurrentDate, req.MillLineTag)

		recs := ctrl.TransToCoilMainRecords(records)
		res.Records = recs
		ginx.ResSuccess(c, res)
	}
}