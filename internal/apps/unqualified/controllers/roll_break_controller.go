package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekohor/mangosteen/internal/app/ginx"
	"github.com/nekohor/mangosteen/internal/apps/unqualified/services"
)

type RollBreakController struct {
}


type RollBreakStatRequest struct {
	//CoilId string `json:"coilId"`
	QueryDate string `json:"queryDate"`
}

// GetRollBreakStatResult swagger
// @Tags  查询数据
// @Summary 根据日期查询轧破甩尾统计数
// @Security ApiKeyAuth
// @Param queryDate body string true "查询日期" default({"queryDate":"2021-05-20"})
// @Success 200 {object} schema.Response 	"{"code": 200, "data": {"rollBreakCountTotal": 6}"
// @Failure 401 {object} schema.Response "未授权"
// @Failure 500 {object} schema.Response "服务器错误"
// @Router /api/v1/unqualified/rollbreak [post]
func (* RollBreakController) GetRollBreakStatResult(c *gin.Context) {
	var err error
	var req RollBreakStatRequest

	err = ginx.ParseJSON(c, &req)
	if err != nil {
		ginx.ResSuccess(c, "亲，您的请求参数无法绑定")
		return
	}

	s := services.NewRollBreakStatService()
	res := s.GetRollBreakStatResult(req.QueryDate)
	ginx.ResSuccess(c, res)
}