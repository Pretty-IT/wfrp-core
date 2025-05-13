package charlist

import (
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/weapons"
)

func CommonCharacter() *State {
	return PlainCharacter(31, 10)
}

func PlainCharacter(charValue int, skillValue int) *State {
	return PlainCharacterWithName("", charValue, skillValue)
}

// PlainCharacterWithName - predefined character with all characteristics and skills equal to given value
func PlainCharacterWithName(name string, charValue int, skillValue int) *State {
	ws := chars.From(chars.WeaponSkill, charValue)
	s := chars.From(chars.Strength, charValue)
	t := chars.From(chars.Toughness, charValue)
	wp := chars.From(chars.Willpower, charValue)

	brawling := skills.From(skills.MeleeBrawling, skillValue)

	punches := weapons.Punches

	return NewState(
		name,
		[]*chars.Value{ws, s, t, wp},
		[]*skills.Value{brawling},
		[]*weapons.Value{punches},
	)
}
