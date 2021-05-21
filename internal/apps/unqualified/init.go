package unqualified

import (
	"github.com/google/wire"
	"github.com/nekohor/mangosteen/internal/apps/unqualified/controllers"
)

var AppSet = wire.NewSet(
	wire.NewSet(wire.Struct(new(controllers.RollBreakController), "*")),
)