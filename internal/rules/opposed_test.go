package rules

import (
	a "github.com/Pretty-IT/wfrp-core/internal/models/actions"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
	"github.com/Pretty-IT/wfrp-core/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

type opposedTd struct {
	attackerInput *a.Opposed
	defenderInput *a.Opposed
	expected      *a.OpposedResult
}

func createOpposed(action a.Type, skillID skills.ID, charValue int, skillValue int, SL int) *a.Opposed {
	return &a.Opposed{
		State:       charlist.PlainCharacter(charValue, skillValue),
		RollRequest: &a.RollRequest{SkillId: skillID, ActionType: action},
		RollResult:  &a.RollResult{SL: SL},
	}
}

var opposedTestDataProvider = []*test.Data[*opposedTd]{
	{"Attacker wins with SL", &opposedTd{
		createOpposed(a.Attack, skills.MeleeBrawling, 31, 10, 2),
		createOpposed(a.Defend, skills.MeleeBrawling, 31, 10, 0),
		&a.OpposedResult{SL: 2, IsAttackerWinner: true, IsTie: false},
	}},
	{"Defender wins with SL", &opposedTd{
		createOpposed(a.Attack, skills.MeleeBrawling, 31, 10, 0),
		createOpposed(a.Defend, skills.MeleeBrawling, 31, 10, 2),
		&a.OpposedResult{SL: -2, IsAttackerWinner: false, IsTie: false},
	}},
	{"Attacker wins with skill", &opposedTd{
		createOpposed(a.Attack, skills.MeleeBrawling, 31, 20, -2),
		createOpposed(a.Defend, skills.MeleeBrawling, 31, 10, -2),
		&a.OpposedResult{SL: 0, IsAttackerWinner: true, IsTie: false},
	}},
	{"Defender wins with skill", &opposedTd{
		createOpposed(a.Attack, skills.MeleeBrawling, 31, 10, 1),
		createOpposed(a.Defend, skills.MeleeBrawling, 31, 20, 1),
		&a.OpposedResult{SL: 0, IsAttackerWinner: false, IsTie: false},
	}},
	{"Attacker wins with char", &opposedTd{
		createOpposed(a.Attack, skills.MeleeBrawling, 41, 10, -1),
		createOpposed(a.Defend, skills.MeleeBrawling, 31, 10, -1),
		&a.OpposedResult{SL: 0, IsAttackerWinner: true, IsTie: false},
	}},
	{"Defender wins with char", &opposedTd{
		createOpposed(a.Attack, skills.MeleeBrawling, 31, 10, 0),
		createOpposed(a.Defend, skills.MeleeBrawling, 41, 10, 0),
		&a.OpposedResult{SL: 0, IsAttackerWinner: false, IsTie: false},
	}},
	{"A tie", &opposedTd{
		createOpposed(a.Attack, skills.MeleeBrawling, 41, 10, 3),
		createOpposed(a.Defend, skills.MeleeBrawling, 41, 10, 3),
		&a.OpposedResult{SL: 0, IsAttackerWinner: false, IsTie: true},
	}},
}

func TestOpposedTest(t *testing.T) {
	test.ParameterizedTest(t, opposedTestDataProvider, func(t *testing.T, data *opposedTd) {
		// Act
		actual := OpposedTest(data.attackerInput, data.defenderInput)

		//Assert
		assert.Equal(t, data.expected, actual)
	})
}
