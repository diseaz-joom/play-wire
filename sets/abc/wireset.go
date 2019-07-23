package abc

import (
	"projects/play-wire/boxes/a"
	"projects/play-wire/boxes/b"
	"projects/play-wire/boxes/c"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	a.MakeBox,
	b.MakeBox,
	c.MakeBox,
	wire.Bind(new(a.I), new(*a.Box)),
	wire.Bind(new(b.I), new(*b.Box)),
	wire.Bind(new(c.I), new(*c.Box)),
)
