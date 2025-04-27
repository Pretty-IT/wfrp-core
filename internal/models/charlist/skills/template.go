package skills

import "github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"

type ID int

const (
	MeleeBrawling ID = iota
)

type Template struct {
	id             ID
	name           string
	char           chars.ID
	isBasic        bool
	specialization string
}

func newTemplate(id ID, name string, specialization string, char chars.ID, isBasic bool) *Template {
	return &Template{
		id:             id,
		name:           name,
		char:           char,
		specialization: specialization,
		isBasic:        isBasic,
	}
}

var Template_value = map[ID]*Template{
	MeleeBrawling: newTemplate(MeleeBrawling, "Melee", "Brawling", chars.WeaponSkill, true),
}
