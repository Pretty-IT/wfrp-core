package rules

import (
	"github.com/Pretty-IT/wfrp-core/internal/models"
)

func OpposedTest(attacker *models.Opposed, defender *models.Opposed) *models.OpposedResult {
	attackerSkill := attacker.RollRequest.SkillId
	defenderSkill := defender.RollRequest.SkillId

	resultSL := attacker.RollResult.SL - defender.RollResult.SL
	resultSkill := attacker.State.GetSkillValue(attackerSkill) - defender.State.GetSkillValue(defenderSkill)
	resultChar := attacker.State.GetCharValue(attackerSkill) - defender.State.GetCharValue(defenderSkill)

	return &models.OpposedResult{
		SL:               resultSL,
		IsAttackerWinner: resultSL > 0 || resultSL == 0 && resultSkill > 0 || resultSL == 0 && resultSkill == 0 && resultChar > 0,
		IsTie:            resultSL == 0 && resultSkill == 0 && resultChar == 0,
	}
}
