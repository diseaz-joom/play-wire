package e

import (
	"fmt"
	"github.com/diseaz-joom/play-wire/boxes/b"
	"github.com/diseaz-joom/play-wire/boxes/c"
	"github.com/diseaz-joom/play-wire/common"
)

type Box struct {
	common.Printer
}

type I fmt.Stringer

func MakeBox(bb b.I, bc c.I) *Box {
	r := &Box{}
	r.Init("e.Box", bb, bc)
	return r
}
