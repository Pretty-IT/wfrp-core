package skills

import (
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
)

type Value struct {
	template Template
	initial  int
}

func From(id ID, initial int) *Value {
	var tmpl = *TemplateValue[id]

	return &Value{
		template: tmpl,
		initial:  initial,
	}
}

func (v *Value) Name() string {
	return v.template.name
}

func (v *Value) FullName() string {
	return v.template.FullName()
}

func (v *Value) IsBasic() bool { return v.template.isBasic }

func (v *Value) Specialization() string { return v.template.specialization }

func (v *Value) Initial() int { return v.initial }

func (v *Value) IsGrouped() bool { return v.template.IsGrouped() }

func (v *Value) CharID() chars.ID { return v.template.char }

func (v *Value) Total() int { return v.initial }

func (v *Value) String() string { return v.FullName() }

func GetID(v *Value) ID { return v.template.id }
