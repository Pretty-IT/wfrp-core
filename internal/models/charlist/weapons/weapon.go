package weapons

import (
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
)

type Weapon struct {
	Name      string
	Skill     *skills.Skill
	Damage    int
	BonusChar *chars.Characteristic
}

func New(name string, skill *skills.Skill, damage int, bonusChar *chars.Characteristic) *Weapon {
	return &Weapon{
		Name:      name,
		Skill:     skill,
		Damage:    damage,
		BonusChar: bonusChar,
	}
}

func (w *Weapon) String() string { return w.Name }
