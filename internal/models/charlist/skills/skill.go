package skills

import (
	"fmt"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
)

type Skill struct {
	name           string
	characteristic *chars.Characteristic
	isBasic        bool
	specialization string

	initial int
}

func New(name string, specialization string, characteristic *chars.Characteristic, isBasic bool, initial int) *Skill {
	return &Skill{
		name:           name,
		characteristic: characteristic,
		isBasic:        isBasic,
		specialization: specialization,
		initial:        initial,
	}
}

func (skill *Skill) Name() string { return skill.name }

func (skill *Skill) IsBasic() bool { return skill.isBasic }

func (skill *Skill) Specialization() string { return skill.specialization }

func (skill *Skill) Initial() int { return skill.initial }

func (skill *Skill) IsGrouped() bool { return skill.specialization != "" }

func (skill *Skill) Total() int { return skill.characteristic.Total() + skill.initial }

func (skill *Skill) String() string {
	result := skill.name
	if skill.IsGrouped() {
		result = fmt.Sprintf("%s(%s)", skill.name, skill.specialization)
	}
	return result
}
