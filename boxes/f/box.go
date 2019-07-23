package f

import (
	"fmt"
	"projects/play-wire/boxes/d"
	"projects/play-wire/boxes/e"
	"projects/play-wire/common"
)

type Box struct {
	*common.Printer
}

type I fmt.Stringer

func MakeBox(bd d.I, be e.I) *Box {
	return &Box{
		Printer: common.NewPrinter("f.Box", bd, be),
	}
}
