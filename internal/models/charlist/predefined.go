package charlist

import (
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/weapons"
)

func CommonCharacter() *State {
	return PlainCharacter(31, 10)
}

// PlainCharacter - predefined character with all characteristics and skills equal to given value
func PlainCharacter(charValue int, skillValue int) *State {
	ws := chars.From(chars.WeaponSkill, charValue)
	s := chars.From(chars.Strength, charValue)
	t := chars.From(chars.Toughness, charValue)
	wp := chars.From(chars.Willpower, charValue)

	brawling := skills.From(skills.MeleeBrawling, skillValue)

	punches := weapons.New("Punches", skills.MeleeBrawling, 0, chars.Strength)

	return NewState(
		[]*chars.Value{ws, s, t, wp},
		[]*skills.Value{brawling},
		[]*weapons.Value{punches},
	)
}
