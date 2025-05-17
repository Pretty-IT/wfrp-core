package rules

import (
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
)

var emptyResult = &models.AttackResult{
	Damage: &models.AttackDamage{},
}

func Attack(attacker *models.Opposed, defender *models.Opposed) (*models.AttackResult, *models.AttackResult) {
	opposedResult := OpposedTest(attacker, defender)

	if opposedResult.IsTie {
		return emptyResult, emptyResult
	}
	if !opposedResult.IsAttackerWinner {
		return emptyResult, emptyResult
	}
	if opposedResult.IsAttackerWinner {
		return &models.AttackResult{Damage: assignWeaponDamage(attacker, opposedResult.SL)}, emptyResult
	}

	return nil, nil
}

func assignWeaponDamage(winner *models.Opposed, opposedSL int) *models.AttackDamage {
	weaponDamage := winner.RollRequest.Weapon.Damage
	attackerBonusChar := winner.RollRequest.Weapon.BonusChar

	if attackerBonusChar != chars.Undefined {
		weaponDamage += winner.State.Chars[attackerBonusChar].Bonus()
	}

	return &models.AttackDamage{
		FromWeapon: weaponDamage,
		FromSL:     opposedSL,
	}
}
