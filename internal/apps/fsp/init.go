package fsp

import (
	"github.com/google/wire"
	"github.com/nekohor/mangosteen/internal/apps/fsp/controllers"
)

var AppSet = wire.NewSet(
	wire.NewSet(wire.Struct(new(controllers.CoilMainController), "*")),
)