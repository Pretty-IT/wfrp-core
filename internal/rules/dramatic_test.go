package rules

import (
	a "github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

type drTD struct {
	inputDice int
	inputRoll *a.Roll
	expected  *a.RollResult
}

var dramaticTestDataProvider = []*test.Data[*drTD]{
	{"Passed Roll", &drTD{13, &a.Roll{TargetValue: 31}, &a.RollResult{Dice: 13, IsPassed: true, SL: 1}}},
	{"Failed Roll", &drTD{68, &a.Roll{TargetValue: 31}, &a.RollResult{Dice: 68, IsPassed: false, SL: -3}}},
	{"Autopassed Roll", &drTD{5, &a.Roll{TargetValue: 58}, &a.RollResult{Dice: 5, IsPassed: true, SL: 5}}},
	{"Autopassed Roll more target", &drTD{4, &a.Roll{TargetValue: 3}, &a.RollResult{Dice: 4, IsPassed: true, SL: 0}}},
	{"Autofailed Roll", &drTD{98, &a.Roll{TargetValue: 54}, &a.RollResult{Dice: 98, IsPassed: false, SL: -4}}},
	{"Autofailed Roll less target", &drTD{96, &a.Roll{TargetValue: 99}, &a.RollResult{Dice: 96, IsPassed: false, SL: 0}}},
	{"Roll 00 -> 100", &drTD{0, &a.Roll{TargetValue: 30}, &a.RollResult{Dice: 100, IsPassed: false, SL: -7, IsFumble: true}}},
	{"Critical Roll", &drTD{33, &a.Roll{TargetValue: 72}, &a.RollResult{Dice: 33, IsPassed: true, SL: 3, IsCritical: true}}},
	{"Fumble Roll", &drTD{55, &a.Roll{TargetValue: 35}, &a.RollResult{Dice: 55, IsPassed: false, SL: -2, IsFumble: true}}},
}

func TestDramaticTest(t *testing.T) {
	test.ParameterizedTest(t, dramaticTestDataProvider, func(t *testing.T, data *drTD) {
		// Act
		var actual = DramaticTest(data.inputDice, data.inputRoll)

		//Assert
		assert.Equal(t, data.expected, actual)
	})
}
