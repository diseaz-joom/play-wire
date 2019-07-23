package a

import "github.com/diseaz-joom/play-wire/common"

type Box struct {
	common.Printer
}

type I interface {
	String() string
}

func MakeBox() *Box {
	r := &Box{}
	r.Init("a.Box")
	return r
}
