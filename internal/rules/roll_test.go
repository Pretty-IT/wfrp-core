package rules

import (
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
	"github.com/Pretty-IT/wfrp-core/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

type crTD struct {
	actorState   *charlist.State
	targetState  *charlist.State
	actorRequest *models.RollRequest
	environment  *models.Environment

	expected *models.Roll
}

var smokeChar = charlist.PlainCharacter(31, 10)

var calculateRollDataProvider = []*test.Data[*crTD]{
	{"Smoke test", &crTD{
		smokeChar,
		charlist.PlainCharacter(31, 10),
		&models.RollRequest{ActionType: models.Attack, SkillID: skills.Undefined, Weapon: smokeChar.Weapons[0]},
		&models.Environment{},
		&models.Roll{TargetValue: 41, ModifyOptions: []models.RollModifyOption{}},
	}},
}

func TestCalculateRoll(t *testing.T) {
	test.ParameterizedTest(t, calculateRollDataProvider, func(t *testing.T, data *crTD) {
		// Act
		actual := CalculateRoll(data.actorState, data.targetState, data.actorRequest, data.environment)

		//Assert
		assert.Equal(t, data.expected, actual)
	})
}
