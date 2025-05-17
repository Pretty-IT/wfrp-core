package charlist

import (
	"fmt"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/weapons"
	"github.com/Pretty-IT/wfrp-core/internal/utils"
	"maps"
	"slices"
	"strings"
)

type State struct {
	Name    string
	Chars   map[chars.ID]*chars.Value
	Skills  map[skills.ID]*skills.Value
	Wounds  int
	Weapons []*weapons.Value
}

func NewState(name string, c []*chars.Value, s []*skills.Value, w []*weapons.Value) *State {
	result := &State{
		Name:    name,
		Chars:   utils.SliceToMap(c, chars.GetID),
		Skills:  utils.SliceToMap(s, skills.GetID),
		Wounds:  0,
		Weapons: w,
	}

	result.SetWounds(result.MaxWounds())
	return result
}

func From(state *State) *State {
	result := &State{
		Name:    state.Name,
		Chars:   utils.CopyMap(state.Chars),
		Skills:  utils.CopyMap(state.Skills),
		Wounds:  state.Wounds,
		Weapons: utils.CopyList(state.Weapons),
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

func (s *State) GetSkillValue(id skills.ID) int {
	skill := s.Skills[id]
	char := s.Chars[skill.CharID()]
	return char.Total() + skill.Total()
}

func (s *State) GetCharValue(skillID skills.ID) int {
	skill := s.Skills[skillID]
	char := s.Chars[skill.CharID()]
	return char.Total()
}

func (s *State) String() string {
	return fmt.Sprintf(`
	Name   : %s,
	Chars  : %s,
	Skills : %s,
	Wounds : %d
`,
		s.Name, s.getCharsString(), s.getSkillsString(), s.Wounds)
}

func (s *State) getCharsString() string {
	results := utils.MapMap(s.Chars, func(v *chars.Value) string {
		return fmt.Sprintf("%s: %d", v.ShortName(), v.Total())
	})
	return strings.Join(slices.Collect(maps.Values(results)), ", ")
}

func (s *State) getSkillsString() string {
	results := utils.MapMap(s.Skills, func(v *skills.Value) string {
		return fmt.Sprintf("%s: %d", v.FullName(), v.Total())
	})
	return strings.Join(slices.Collect(maps.Values(results)), ", ")
}
