package e

import (
	"fmt"
	"projects/play-wire/boxes/b"
	"projects/play-wire/boxes/c"
	"projects/play-wire/common"
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
