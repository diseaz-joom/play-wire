package main

import (
	"fmt"
	"github.com/diseaz-joom/play-wire/svc/two/wire"
)

func main() {
	svc := wire.MakeService()
	fmt.Printf("%s\n", svc)
}
