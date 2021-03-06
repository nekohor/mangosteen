// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package app

import (
	"github.com/nekohor/mangosteen/internal/app/router"
	"github.com/nekohor/mangosteen/internal/apps/category/controllers"
	"github.com/nekohor/mangosteen/internal/apps/category/services"
	controllers2 "github.com/nekohor/mangosteen/internal/apps/fsp/controllers"
	controllers3 "github.com/nekohor/mangosteen/internal/apps/unqualified/controllers"
)

import (
	_ "github.com/nekohor/mangosteen/docs/swagger"
)

// Injectors from wire.go:

func BuildInjector() (*Injector, func(), error) {
	gradeKlassService := &services.GradeKlassService{}
	categoryController := &controllers.CategoryController{
		GradeKlassService: gradeKlassService,
	}
	coilMainController := &controllers2.CoilMainController{}
	rollBreakController := &controllers3.RollBreakController{}
	routerRouter := &router.Router{
		CategoryController:  categoryController,
		CoilMainController:  coilMainController,
		RollBreakController: rollBreakController,
	}
	engine := InitGinEngine(routerRouter)
	injector := &Injector{
		Engine: engine,
	}
	return injector, func() {
	}, nil
}
