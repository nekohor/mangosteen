package category

import (
	"github.com/nekohor/mangosteen/internal/apps/category/controllers"
	"github.com/nekohor/mangosteen/internal/apps/category/services"

	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	wire.NewSet(wire.Struct(new(controllers.CategoryController), "*")),
	wire.NewSet(wire.Struct(new(services.GradeKlassService), "*")),
)
