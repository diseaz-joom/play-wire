package b

import (
	"fmt"
	"projects/play-wire/common"
)

type Box struct {
	common.Printer
}

type I interface {
	fmt.Stringer
}

func MakeBox() *Box {
	r := &Box{}
	r.Init("b.Box")
	return r
}
