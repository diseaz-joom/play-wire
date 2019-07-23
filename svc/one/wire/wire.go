//+build wireinject

package wire

import (
	"projects/play-wire/sets/abc"
	"projects/play-wire/sets/de"
	"projects/play-wire/sets/fg"
	"projects/play-wire/svc/one"

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
