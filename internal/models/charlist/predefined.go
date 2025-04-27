package charlist

import (
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/weapons"
)

func CommonCharacter() *State {
	ws := chars.From(chars.WeaponSkill, 31)
	s := chars.From(chars.Strength, 31)
	t := chars.From(chars.Toughness, 31)
	wp := chars.From(chars.Willpower, 31)

	brawling := skills.From(skills.MeleeBrawling, 10)

	punches := weapons.New("Punches", skills.MeleeBrawling, 0, chars.Strength)

	return NewState(
		[]*chars.Value{ws, s, t, wp},
		[]*skills.Value{brawling},
		[]*weapons.Value{punches},
	)
}
