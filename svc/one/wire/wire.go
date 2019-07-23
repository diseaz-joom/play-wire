//+build wireinject

package wire

import (
	"github.com/diseaz-joom/play-wire/sets/abc"
	"github.com/diseaz-joom/play-wire/sets/de"
	"github.com/diseaz-joom/play-wire/sets/fg"
	"github.com/diseaz-joom/play-wire/svc/one"

	"github.com/google/wire"
)

func MakeService() *one.Service {
	wire.Build(
		abc.Set,
		de.Set,
		fg.Set,
		wire.Struct(new(one.Service), "*"),
	)
	return &one.Service{}
}
