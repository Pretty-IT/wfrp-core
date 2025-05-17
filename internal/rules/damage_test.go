package rules

import (
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
	"github.com/Pretty-IT/wfrp-core/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

type dTD struct {
	state          *charlist.State
	damage         *models.AttackDamage
	expectedWounds int
}

var damageDataProvider = []*test.Data[*dTD]{
	{"Smoke test", &dTD{
		charlist.PlainCharacter(31, 10),
		&models.AttackDamage{FromWeapon: 3, FromSL: 2},
		10,
	}},
}

func TestApplyDamage(t *testing.T) {
	test.ParameterizedTest(t, damageDataProvider, func(t *testing.T, data *dTD) {
		// Act
		actual := ApplyDamage(data.state, data.damage)

		//Assert
		assert.Equal(t, data.expectedWounds, actual.Wounds)
	})
}
