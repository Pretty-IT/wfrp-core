package rules

import (
	"github.com/Pretty-IT/wfrp-core/internal/models"
)

// DramaticTest - dice should be in interval [0-99]
func DramaticTest(dice int, roll *models.Roll) *models.RollResult {
	var autoPassed = 6
	var autoFailed = 96

	var isDiceDouble = dice/10 == dice%10
	if dice == 0 {
		dice = 100
	}
	var result = roll.TargetValue + roll.TotalModifier - dice
	var isPassed = dice < autoPassed || dice < autoFailed && result >= 0

	return &models.RollResult{
		Dice:       dice,
		IsPassed:   isPassed,
		SL:         (result / 10) + roll.TotalSLModifier,
		IsCritical: isDiceDouble && isPassed,
		IsFumble:   isDiceDouble && !isPassed,
	}
}
