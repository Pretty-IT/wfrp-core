package rules

import (
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/actions"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
)

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
		TargetValue:     target,
		TotalModifier:   0,
		TotalSLModifier: 0,
		ModifyOptions:   []actions.RollModifyOption{},
	}
}
