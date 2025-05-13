package skills

import (
	"fmt"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
)

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

func (t *Template) FullName() string {
	result := t.name
	if t.IsGrouped() {
		result = fmt.Sprintf("%s(%s)", t.name, t.specialization)
	}
	return result
}

func (t *Template) IsGrouped() bool { return t.specialization != "" }

func newTemplate(id ID, name string, specialization string, char chars.ID, isBasic bool) *Template {
	return &Template{
		id:             id,
		name:           name,
		char:           char,
		specialization: specialization,
		isBasic:        isBasic,
	}
}

var TemplateValue = map[ID]*Template{
	MeleeBrawling: newTemplate(MeleeBrawling, "Melee", "Brawling", chars.WeaponSkill, true),
}
