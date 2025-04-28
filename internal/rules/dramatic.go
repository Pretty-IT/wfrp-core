package rules

import "github.com/Pretty-IT/wfrp-core/internal/models/actions"

// DramaticTest - dice should be in interval [0-99]
func DramaticTest(dice int, roll *actions.Roll) *actions.RollResult {
	var autoPassed = 6
	var autoFailed = 96

	var isDiceDouble = dice/10 == dice%10
	if dice == 0 {
		dice = 100
	}
	var result = roll.TargetValue + roll.TotalModifier - dice
	var isPassed = dice < autoPassed || dice < autoFailed && result >= 0

	return &actions.RollResult{
		IsPassed:   isPassed,
		SL:         (result / 10) + roll.TotalSLModifier,
		IsCritical: isDiceDouble && isPassed,
		IsFumble:   isDiceDouble && !isPassed,
	}
}
