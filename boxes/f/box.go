package f

import (
	"fmt"
	"github.com/diseaz-joom/play-wire/boxes/d"
	"github.com/diseaz-joom/play-wire/boxes/e"
	"github.com/diseaz-joom/play-wire/common"
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
