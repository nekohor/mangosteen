package controllers

import (
	"github.com/nekohor/mangosteen/internal/app/ginx"
	"github.com/nekohor/mangosteen/internal/apps/category/gradeklass"
	"github.com/nekohor/mangosteen/internal/apps/category/services"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	GradeKlassService *services.GradeKlassService
}

type SteelGradeCategoryRequest struct {
	//CoilId string `json:"coilId"`
	SteelGrade string `json:"steelGrade"`
}

// GetClassificationResultBySteelGrade swagger
// @Tags  查询数据
// @Summary 根据钢种牌号查询钢种分类结果，目前有4个维度
// @Security ApiKeyAuth
// @Param steelGrade body string true "钢种牌号" default({"steelGrade":"M510L"})
// @Success 200 {object} schema.Response 	"{"code": 200, "data": {"steelGrade": "M510L", "category1": "商品材", "category2": "高强商品材", "category3": "大梁钢", "category4": ""}, "message": ""}"
// @Failure 401 {object} schema.Response "未授权"
// @Failure 500 {object} schema.Response "服务器错误"
// @Router /api/v1/category/steelgrade [post]
func (ctrl *CategoryController) GetClassificationResultBySteelGrade(c *gin.Context) {
	var err error
	var req SteelGradeCategoryRequest
	var res *gradeklass.SteelGradeClassificationResult

	err = ginx.ParseJSON(c, &req)
	if err != nil {
		ginx.ResSuccess(c, "亲，您的参数无法绑定，注意牌号的键名为steelGrade，且注意大小写哦！")
	} else {
		res = ctrl.GradeKlassService.GetClassificationResultBySteelGrade(req.SteelGrade)
		ginx.ResSuccess(c, res)
	}
}
