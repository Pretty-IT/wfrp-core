package weapons

import (
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
)

type Value struct {
	Name      string
	Skill     skills.ID
	Damage    int
	BonusChar chars.ID
}

func New(name string, skill skills.ID, damage int, bonusChar chars.ID) *Value {
	return &Value{
		Name:      name,
		Skill:     skill,
		Damage:    damage,
		BonusChar: bonusChar,
	}
}

func (v *Value) String() string { return v.Name }
