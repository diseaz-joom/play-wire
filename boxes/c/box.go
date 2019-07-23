package c

import "projects/play-wire/common"

type Box struct {
	common.Printer
}

type I interface {
	String() string
}

func MakeBox() *Box {
	r := &Box{}
	r.Init("c.Box")
	return r
}
