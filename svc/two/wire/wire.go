//+build wireinject

package wire

import (
	"projects/play-wire/boxes/g"
	"projects/play-wire/sets/abc"
	"projects/play-wire/sets/de"
	"projects/play-wire/sets/fg"

	"github.com/google/wire"
)

func MakeService() *g.Box {
	wire.Build(
		abc.Set,
		de.Set,
		fg.Set,
	)
	return &g.Box{}
}
