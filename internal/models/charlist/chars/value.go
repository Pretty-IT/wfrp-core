package chars

type Value struct {
	template Template
	initial  int
}

func From(id ID, initial int) *Value {
	var tmpl = *Template_value[id]

	return &Value{
		template: tmpl,
		initial:  initial,
	}
}

func (v *Value) Name() string { return v.template.name }

func (v *Value) Initial() int { return v.initial }

func (v *Value) Total() int { return v.initial }

func (v *Value) Bonus() int { return v.Total() / 10 }

func (v *Value) String() string {
	return v.template.name
}

func GetID(v *Value) ID {
	return v.template.id
}
