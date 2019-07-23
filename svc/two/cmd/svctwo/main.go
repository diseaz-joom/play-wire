package main

import (
	"fmt"
	"projects/play-wire/svc/two/wire"
)

func main() {
	svc := wire.MakeService()
	fmt.Printf("%s\n", svc)
}
