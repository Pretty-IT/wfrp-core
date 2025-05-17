package rules

import (
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
	"github.com/Pretty-IT/wfrp-core/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

type opposedTd struct {
	attackerInput *models.Opposed
	defenderInput *models.Opposed
	expected      *models.OpposedResult
}

func createOpposed(action models.Type, skillID skills.ID, charValue int, skillValue int, SL int) *models.Opposed {
	return &models.Opposed{
		State:       charlist.PlainCharacter(charValue, skillValue),
		RollRequest: &models.RollRequest{SkillId: skillID, ActionType: action},
		RollResult:  &models.RollResult{SL: SL},
	}
}

var opposedTestDataProvider = []*test.Data[*opposedTd]{
	{"Attacker wins with SL", &opposedTd{
		createOpposed(models.Attack, skills.MeleeBrawling, 31, 10, 2),
		createOpposed(models.Defend, skills.MeleeBrawling, 31, 10, 0),
		&models.OpposedResult{SL: 2, IsAttackerWinner: true, IsTie: false},
	}},
	{"Defender wins with SL", &opposedTd{
		createOpposed(models.Attack, skills.MeleeBrawling, 31, 10, 0),
		createOpposed(models.Defend, skills.MeleeBrawling, 31, 10, 2),
		&models.OpposedResult{SL: -2, IsAttackerWinner: false, IsTie: false},
	}},
	{"Attacker wins with skill", &opposedTd{
		createOpposed(models.Attack, skills.MeleeBrawling, 31, 20, -2),
		createOpposed(models.Defend, skills.MeleeBrawling, 31, 10, -2),
		&models.OpposedResult{SL: 0, IsAttackerWinner: true, IsTie: false},
	}},
	{"Defender wins with skill", &opposedTd{
		createOpposed(models.Attack, skills.MeleeBrawling, 31, 10, 1),
		createOpposed(models.Defend, skills.MeleeBrawling, 31, 20, 1),
		&models.OpposedResult{SL: 0, IsAttackerWinner: false, IsTie: false},
	}},
	{"Attacker wins with char", &opposedTd{
		createOpposed(models.Attack, skills.MeleeBrawling, 41, 10, -1),
		createOpposed(models.Defend, skills.MeleeBrawling, 31, 10, -1),
		&models.OpposedResult{SL: 0, IsAttackerWinner: true, IsTie: false},
	}},
	{"Defender wins with char", &opposedTd{
		createOpposed(models.Attack, skills.MeleeBrawling, 31, 10, 0),
		createOpposed(models.Defend, skills.MeleeBrawling, 41, 10, 0),
		&models.OpposedResult{SL: 0, IsAttackerWinner: false, IsTie: false},
	}},
	{"A tie", &opposedTd{
		createOpposed(models.Attack, skills.MeleeBrawling, 41, 10, 3),
		createOpposed(models.Defend, skills.MeleeBrawling, 41, 10, 3),
		&models.OpposedResult{SL: 0, IsAttackerWinner: false, IsTie: true},
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
