package rules

import (
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
)

func CalculateRoll(actorState *charlist.State, targetState *charlist.State,
	request *models.RollRequest, environment models.Environment) *models.Roll {
	skillID := request.GetSkillID()
	target := actorState.GetSkillValue(skillID)

	return &models.Roll{
		TargetValue:     target,
		TotalModifier:   0,
		TotalSLModifier: 0,
		ModifyOptions:   []models.RollModifyOption{},
	}
}
