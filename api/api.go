package api

import (
	"fmt"
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/actions"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
)

func HelloWorld() string {
	msg := "Hello from Go!"
	fmt.Println(msg)
	return msg
}

func CalculateRoll(actorState *charlist.State, targetState *charlist.State,
	request actions.RollRequest, environment models.Environment) *actions.Roll {
	var skillID = request.Skill
	if request.Weapon != nil {
		skillID = request.Weapon.Skill
	}
	var skill = actorState.Skills[skillID]
	var char = actorState.Chars[skill.CharID()]

	var target = char.Total() + skill.Total()

	return &actions.Roll{
		TargetCharValue: char.Total(),
		TargetValue:     target,
		TotalModifier:   0,
		TotalSLModifier: 0,
		ModifyOptions:   []actions.RollModifyOption{actions.Reroll},
	}
}

func MakeDramaticTest(dice int, roll *actions.Roll) *actions.RollResult {
	return &actions.RollResult{}
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
