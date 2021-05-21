package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	category "github.com/nekohor/mangosteen/internal/apps/category/controllers"
	fsp "github.com/nekohor/mangosteen/internal/apps/fsp/controllers"
	unqualified "github.com/nekohor/mangosteen/internal/apps/unqualified/controllers"
)

var _ IRouter = (*Router)(nil)

// RouterSet 注入router
var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

// IRouter 注册路由
type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

// Router 路由管理器
type Router struct {
	CategoryController *category.CategoryController
	CoilMainController *fsp.CoilMainController
	RollBreakController *unqualified.RollBreakController
}

// Register 注册路由
func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

// Prefixes 路由前缀列表
func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}
