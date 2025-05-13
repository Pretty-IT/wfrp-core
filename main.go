package main

import (
	"fmt"
	"github.com/Pretty-IT/wfrp-core/api"
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/actions"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
	"github.com/Pretty-IT/wfrp-core/internal/rules"
)

func main() {
	api.HelloWorld()

	attacker := charlist.PlainCharacterWithName("attacker", 31, 10)
	defender := charlist.PlainCharacterWithName("defender", 31, 10)

	aOpposed := getOpposed(actions.Attack, attacker, defender, 10)
	dOpposed := getOpposed(actions.Defend, defender, attacker, 20)

	oTest := rules.OpposedTest(aOpposed, dOpposed)

	fmt.Printf("%+v\n", aOpposed)
	fmt.Printf("%+v\n", dOpposed)
	fmt.Printf("%+v\n", oTest)
}

func getOpposed(action actions.Type, subject *charlist.State, target *charlist.State, dice int) *actions.Opposed {
	weapon := subject.Weapons[0]
	request := &actions.RollRequest{ActionType: action, Weapon: weapon}

	roll := rules.CalculateRoll(subject, target, request, models.Environment{})
	dramatic := rules.DramaticTest(dice, roll)

	return &actions.Opposed{
		State:       subject,
		RollRequest: request,
		Roll:        roll,
		RollResult:  dramatic,
	}
}
