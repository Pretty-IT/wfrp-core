package api

import (
	"fmt"
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
	"github.com/Pretty-IT/wfrp-core/internal/rules"
)

func HelloWorld() string {
	msg := "Hello from Go!"
	fmt.Println(msg)
	return msg
}

func CalculateRoll(actorState *charlist.State, targetState *charlist.State,
	request *models.RollRequest, environment models.Environment) *models.Roll {
	return rules.CalculateRoll(actorState, targetState, request, environment)
}

// MakeDramaticTest - dice should be in interval [0-99]
func MakeDramaticTest(dice int, roll *models.Roll) *models.RollResult {
	return rules.DramaticTest(dice, roll)
}

func MakeOpposedTest(attacker *models.Opposed, defender *models.Opposed) *models.OpposedResult {
	return rules.OpposedTest(attacker, defender)
}

func MakeAttack(attacker *models.Opposed, defender *models.Opposed) (*models.AttackResult, *models.AttackResult) {
	return rules.Attack(attacker, defender)
}

func ApplyDamage(state *charlist.State, damage *models.AttackDamage) *charlist.State {
	return rules.ApplyDamage(state, damage)
}
