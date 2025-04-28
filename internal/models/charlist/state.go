package charlist

import (
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/weapons"
	"github.com/Pretty-IT/wfrp-core/internal/utils"
)

type State struct {
	Chars   map[chars.ID]*chars.Value
	Skills  map[skills.ID]*skills.Value
	Wounds  int
	Weapons []*weapons.Value
}

func NewState(c []*chars.Value, s []*skills.Value, w []*weapons.Value) *State {
	result := &State{
		Chars:   utils.SliceToMap(c, chars.GetID),
		Skills:  utils.SliceToMap(s, skills.GetID),
		Wounds:  0,
		Weapons: w,
	}

	result.SetWounds(result.MaxWounds())
	return result
}

func (s *State) MaxWounds() int {
	return s.Chars[chars.Strength].Bonus() +
		s.Chars[chars.Toughness].Bonus()*2 +
		s.Chars[chars.Willpower].Bonus()
}

func (s *State) SetWounds(value int) {
	s.Wounds = value
}

type Delta struct{}
