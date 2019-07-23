package d

import (
	"fmt"
	"projects/play-wire/boxes/a"
	"projects/play-wire/boxes/b"
	"projects/play-wire/common"
)

type Box struct {
	common.Printer
}

type I fmt.Stringer

func MakeBox(ba a.I, bb b.I) *Box {
	r := &Box{}
	r.Init("d.Box", ba, bb)
	return r
}
