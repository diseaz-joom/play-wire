package fg

import (
	"github.com/diseaz-joom/play-wire/boxes/f"
	"github.com/diseaz-joom/play-wire/boxes/g"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	// de.Set,  <- does not work, complains about duplicate providers
	f.MakeBox,
	g.MakeBox,
	wire.Bind(new(f.I), new(*f.Box)),
	wire.Bind(new(g.I), new(*g.Box)),
)
