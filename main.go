package main

import (
	"fmt"
	"github.com/Pretty-IT/wfrp-core/api"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
)

func main() {
	api.HelloWorld()

	fmt.Printf("%+v\n", charlist.CommonCharacter())
}
