package charlist

import (
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/weapons"
)

func CommonCharacter() *State {
	ws := chars.New(chars.WeaponSkill, 31)
	s := chars.New(chars.Strength, 31)
	t := chars.New(chars.Toughness, 31)
	wp := chars.New(chars.Willpower, 31)

	brawling := skills.New("Melee", "Brawling", ws, true, 10)

	punches := weapons.New("Punches", brawling, 0, s)
	return NewState(
		[]*chars.Characteristic{ws, s, t, wp},
		[]*skills.Skill{brawling},
		[]*weapons.Weapon{punches},
	)
}
