//+build wireinject

package wire

import (
	"github.com/diseaz-joom/play-wire/boxes/g"
	"github.com/diseaz-joom/play-wire/sets/abc"
	"github.com/diseaz-joom/play-wire/sets/de"
	"github.com/diseaz-joom/play-wire/sets/fg"

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
