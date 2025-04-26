package charlist

import (
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/weapons"
	"github.com/Pretty-IT/wfrp-core/internal/utils"
)

type State struct {
	Chars   []*chars.Characteristic
	Skills  []*skills.Skill
	Wounds  int
	Weapons []*weapons.Weapon
}

func NewState(chars []*chars.Characteristic, skills []*skills.Skill, weapons []*weapons.Weapon) *State {
	result := &State{
		Chars:   chars,
		Skills:  skills,
		Wounds:  0,
		Weapons: weapons,
	}

	result.SetWounds(result.MaxWounds())
	return result
}

func (state *State) MaxWounds() int {
	return state.Char(chars.Strength).Bonus() +
		state.Char(chars.Toughness).Bonus()*2 +
		state.Char(chars.Willpower).Bonus()
}

func (state *State) Char(name chars.CharacteristicName) *chars.Characteristic {
	result, err := utils.FindFirst(state.Chars, isCharHasName(name))
	if err != nil {
		panic(err)
	}
	return result
}

func (state *State) SetWounds(value int) {
	state.Wounds = value
}

func isCharHasName(name chars.CharacteristicName) func(*chars.Characteristic) bool {
	return func(ch *chars.Characteristic) bool {
		return ch.Name() == name
	}
}

type Delta struct{}
