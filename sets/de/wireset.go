package de

import (
	"projects/play-wire/boxes/d"
	"projects/play-wire/boxes/e"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	d.MakeBox,
	e.MakeBox,
	wire.Bind(new(d.I), new(*d.Box)),
	wire.Bind(new(e.I), new(*e.Box)),
)
