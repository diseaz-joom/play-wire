package main

import (
	"fmt"
	"github.com/diseaz-joom/play-wire/svc/one/wire"
)

func main() {
	svc := wire.MakeService()
	fmt.Printf("svc.F: %s\nsvc.G: %s\n", svc.F, svc.G)
}
