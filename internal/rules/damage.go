package rules

import (
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
)

func ApplyDamage(state *charlist.State, damage *models.AttackDamage) *charlist.State {
	newState := charlist.From(state)
	resolvedDamage := damage.Total() - newState.Chars[chars.Toughness].Bonus()

	newState.Wounds -= resolvedDamage
	if newState.Wounds < 0 {
		newState.Wounds = 0
	}
	return newState
}
