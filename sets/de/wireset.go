package de

import (
	"github.com/diseaz-joom/play-wire/boxes/d"
	"github.com/diseaz-joom/play-wire/boxes/e"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	d.MakeBox,
	e.MakeBox,
	wire.Bind(new(d.I), new(*d.Box)),
	wire.Bind(new(e.I), new(*e.Box)),
)
