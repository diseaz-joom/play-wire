package common

import (
	"fmt"
	"strings"
)

type Printer struct {
	Name string
	Refs []fmt.Stringer
}

func NewPrinter(name string, refs ...fmt.Stringer) *Printer {
	p := &Printer{}
	p.Init(name, refs...)
	return p
}

func (p *Printer) Init(name string, refs ...fmt.Stringer) {
	p.Name = name
	p.Refs = refs
}

func (p *Printer) String() string {
	var s []string
	for _, r := range p.Refs {
		s = append(s, r.String())
	}
	return fmt.Sprintf("%s{%s}", p.Name, strings.Join(s, ", "))
}
