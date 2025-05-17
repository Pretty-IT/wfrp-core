package main

import (
	"fmt"
	"github.com/Pretty-IT/wfrp-core/api"
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
	"github.com/Pretty-IT/wfrp-core/internal/rules"
)

func main() {
	api.HelloWorld()

	attacker := charlist.PlainCharacterWithName("attacker", 31, 10)
	defender := charlist.PlainCharacterWithName("defender", 31, 10)

	aOpposed := getOpposed(models.Attack, attacker, defender, 10)
	dOpposed := getOpposed(models.Defend, defender, attacker, 60)

	attackerResult, defenderResult := rules.Attack(aOpposed, dOpposed)
	if attackerResult.Damage.Total() > 0 {
		defender = rules.ApplyDamage(defender, attackerResult.Damage)
	}
	fmt.Printf("%+v\n", aOpposed)
	fmt.Printf("%+v\n", dOpposed)
	fmt.Printf("%+v\n", attackerResult)
	fmt.Printf("%+v\n", defenderResult)

	fmt.Printf("%+v\n", attacker)
	fmt.Printf("%+v\n", defender)
}

func getOpposed(action models.Type, subject *charlist.State, target *charlist.State, dice int) *models.Opposed {
	weapon := subject.Weapons[0]
	request := &models.RollRequest{ActionType: action, Weapon: weapon}

	roll := rules.CalculateRoll(subject, target, request, models.Environment{})
	dramatic := rules.DramaticTest(dice, roll)

	return &models.Opposed{
		State:       subject,
		RollRequest: request,
		Roll:        roll,
		RollResult:  dramatic,
	}
}
