package rules

import (
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/weapons"
	"github.com/Pretty-IT/wfrp-core/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

type aTD struct {
	attackerInput    *models.Opposed
	defenderInput    *models.Opposed
	expectedAttacker *models.AttackResult
	expectedDefender *models.AttackResult
}

func createOpposedWithWeapon(action models.Type, weapon *weapons.Value, charValue int, skillValue int, SL int) *models.Opposed {
	return &models.Opposed{
		State:       charlist.PlainCharacter(charValue, skillValue),
		RollRequest: &models.RollRequest{ActionType: action, SkillID: skills.Undefined, Weapon: weapon},
		RollResult:  &models.RollResult{SL: SL},
	}
}

var attackDataProvider = []*test.Data[*aTD]{
	{"Attacker wins with SL", &aTD{
		createOpposedWithWeapon(models.Attack, weapons.Punches, 31, 10, 2),
		createOpposedWithWeapon(models.Defend, weapons.Punches, 31, 10, 0),
		&models.AttackResult{Damage: &models.AttackDamage{FromWeapon: 3, FromSL: 2}},
		&models.AttackResult{Damage: &models.AttackDamage{}},
	}},
	{"Defender wins with SL", &aTD{
		createOpposedWithWeapon(models.Attack, weapons.Punches, 31, 10, 0),
		createOpposedWithWeapon(models.Defend, weapons.Punches, 31, 10, 2),
		&models.AttackResult{Damage: &models.AttackDamage{}},
		&models.AttackResult{Damage: &models.AttackDamage{}},
	}},
	{"A tie", &aTD{
		createOpposedWithWeapon(models.Attack, weapons.Punches, 41, 10, 3),
		createOpposedWithWeapon(models.Defend, weapons.Punches, 41, 10, 3),
		&models.AttackResult{Damage: &models.AttackDamage{}},
		&models.AttackResult{Damage: &models.AttackDamage{}},
	}},
}

func TestAttack(t *testing.T) {
	test.ParameterizedTest(t, attackDataProvider, func(t *testing.T, data *aTD) {
		// Act
		actualAttacker, actualDefender := Attack(data.attackerInput, data.defenderInput)

		//Assert
		assert.Equal(t, data.expectedAttacker, actualAttacker)
		assert.Equal(t, data.expectedDefender, actualDefender)
	})
}
