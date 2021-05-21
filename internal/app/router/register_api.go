package router

import "github.com/gin-gonic/gin"

func (r *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")

	v1 := g.Group("/v1")
	{
		gCategory := v1.Group("category")
		{
			gCategory.POST("steelgrade", r.CategoryController.GetClassificationResultBySteelGrade)
		}

		gCoilMain := v1.Group("coilmain")
		{
			gCoilMain.POST("records", r.CoilMainController.GetCoilMainRecordsByDate)
		}

		gUnqualified := v1.Group("unqualified")
		{
			gUnqualified.POST("rollbreak", r.RollBreakController.GetRollBreakStatResult)
		}
	}
}