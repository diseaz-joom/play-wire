package d

import (
	"fmt"
	"github.com/diseaz-joom/play-wire/boxes/a"
	"github.com/diseaz-joom/play-wire/boxes/b"
	"github.com/diseaz-joom/play-wire/common"
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
