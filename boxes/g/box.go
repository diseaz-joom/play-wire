package g

import (
	"fmt"
	"projects/play-wire/boxes/e"
	"projects/play-wire/common"
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
