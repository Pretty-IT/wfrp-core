package api

import (
	"fmt"
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/actions"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
	"github.com/Pretty-IT/wfrp-core/internal/rules"
)

func HelloWorld() string {
	msg := "Hello from Go!"
	fmt.Println(msg)
	return msg
}

func CalculateRoll(actorState *charlist.State, targetState *charlist.State,
	request *actions.RollRequest, environment models.Environment) *actions.Roll {
	return rules.CalculateRoll(actorState, targetState, request, environment)
}

// MakeDramaticTest - dice should be in interval [0-99]
func MakeDramaticTest(dice int, roll *actions.Roll) *actions.RollResult {
	return rules.DramaticTest(dice, roll)
}

func MakeOpposedTest(attacker *actions.RollResult, defender *actions.RollResult) *actions.OpposedResult {
	return &actions.OpposedResult{}
}

func MakeAttack(attacker *actions.Opposed, defender *actions.Opposed) (*actions.OpposedResult, []*models.ResolveAttackOption) {
	opposedResult := MakeOpposedTest(attacker.RollResult, defender.RollResult)
	resolveOptions := resolveAttack(attacker, defender, opposedResult)
	return opposedResult, resolveOptions
}

func resolveAttack(attacker *actions.Opposed, defender *actions.Opposed, result *actions.OpposedResult) []*models.ResolveAttackOption {
	return []*models.ResolveAttackOption{}
}

func ChangeState(state *charlist.State, delta *charlist.Delta) *charlist.State {
	result := &charlist.State{} //Create deep copy here
	//apply delta to result
	return result
}
