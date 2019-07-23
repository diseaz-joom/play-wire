package fg

import (
	"projects/play-wire/boxes/f"
	"projects/play-wire/boxes/g"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	// de.Set,  <- does not work, complains about duplicate providers
	f.MakeBox,
	g.MakeBox,
	wire.Bind(new(f.I), new(*f.Box)),
	wire.Bind(new(g.I), new(*g.Box)),
)
