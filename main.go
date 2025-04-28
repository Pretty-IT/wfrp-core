package main

import (
	"fmt"
	"github.com/Pretty-IT/wfrp-core/api"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
)

func main() {
	api.HelloWorld()

	var test = *chars.Template_value[chars.WeaponSkill]

	fmt.Printf("%+v\n", test)
	fmt.Printf("%+v\n", chars.Template_value[chars.WeaponSkill])
}
