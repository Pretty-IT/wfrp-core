package rules

import (
	a "github.com/Pretty-IT/wfrp-core/internal/models/actions"
	"github.com/Pretty-IT/wfrp-core/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

type td struct {
	inputDice int
	inputRoll *a.Roll
	expected  *a.RollResult
}

var dramaticTestDataProvider = []*test.Data[*td]{
	{"Passed Roll", &td{13, &a.Roll{TargetValue: 31}, &a.RollResult{IsPassed: true, SL: 1}}},
	{"Failed Roll", &td{68, &a.Roll{TargetValue: 31}, &a.RollResult{IsPassed: false, SL: -3}}},
	{"Autopassed Roll", &td{5, &a.Roll{TargetValue: 58}, &a.RollResult{IsPassed: true, SL: 5}}},
	{"Autopassed Roll more target", &td{4, &a.Roll{TargetValue: 3}, &a.RollResult{IsPassed: true, SL: 0}}},
	{"Autofailed Roll", &td{98, &a.Roll{TargetValue: 54}, &a.RollResult{IsPassed: false, SL: -4}}},
	{"Autofailed Roll less target", &td{96, &a.Roll{TargetValue: 99}, &a.RollResult{IsPassed: false, SL: 0}}},
	{"Roll 00 -> 100", &td{0, &a.Roll{TargetValue: 30}, &a.RollResult{IsPassed: false, SL: -7, IsFumble: true}}},
	{"Critical Roll", &td{33, &a.Roll{TargetValue: 72}, &a.RollResult{IsPassed: true, SL: 3, IsCritical: true}}},
	{"Fumble Roll", &td{55, &a.Roll{TargetValue: 35}, &a.RollResult{IsPassed: false, SL: -2, IsFumble: true}}},
}

func TestDramaticTest(t *testing.T) {
	test.ParameterizedTest(t, dramaticTestDataProvider, func(t *testing.T, data *td) {
		// Act
		var actual = DramaticTest(data.inputDice, data.inputRoll)

		//Assert
		assert.Equal(t, data.expected, actual)
	})
}
