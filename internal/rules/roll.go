package rules

import (
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/actions"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
)

func CalculateRoll(actorState *charlist.State, targetState *charlist.State,
	request actions.RollRequest, environment models.Environment) *actions.Roll {
	skillID := request.SkillId
	if request.Weapon != nil {
		skillID = request.Weapon.Skill
	}

	target := actorState.GetSkillValue(skillID)

	return &actions.Roll{
		TargetValue:     target,
		TotalModifier:   0,
		TotalSLModifier: 0,
		ModifyOptions:   []actions.RollModifyOption{},
	}
}
