// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"
	"github.com/nekohor/mangosteen/internal/app/router"
	"github.com/nekohor/mangosteen/internal/apps/category"
	"github.com/nekohor/mangosteen/internal/apps/fsp"
	"github.com/nekohor/mangosteen/internal/apps/unqualified"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		category.AppSet,
		fsp.AppSet,
		unqualified.AppSet,
		InitGinEngine,
		router.RouterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
