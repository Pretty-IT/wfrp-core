package rules

import (
	"github.com/Pretty-IT/wfrp-core/internal/models/actions"
)

func OpposedTest(attacker *actions.Opposed, defender *actions.Opposed) *actions.OpposedResult {
	attackerSkill := attacker.RollRequest.SkillId
	defenderSkill := defender.RollRequest.SkillId

	resultSL := attacker.RollResult.SL - defender.RollResult.SL
	resultSkill := attacker.State.GetSkillValue(attackerSkill) - defender.State.GetSkillValue(defenderSkill)
	resultChar := attacker.State.GetCharValue(attackerSkill) - defender.State.GetCharValue(defenderSkill)

	return &actions.OpposedResult{
		SL:               resultSL,
		IsAttackerWinner: resultSL > 0 || resultSL == 0 && resultSkill > 0 || resultSL == 0 && resultSkill == 0 && resultChar > 0,
		IsTie:            resultSL == 0 && resultSkill == 0 && resultChar == 0,
	}
}
