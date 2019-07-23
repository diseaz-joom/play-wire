package g

import (
	"fmt"
	"github.com/diseaz-joom/play-wire/boxes/e"
	"github.com/diseaz-joom/play-wire/common"
)

type Box struct {
	*common.Printer
}

type I fmt.Stringer

func MakeBox(be e.I) *Box {
	return &Box{
		Printer: common.NewPrinter("g.Box", be),
	}
}
