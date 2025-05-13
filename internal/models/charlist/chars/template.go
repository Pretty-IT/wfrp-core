package chars

type ID int

const (
	WeaponSkill ID = iota
	BallisticSkill
	Strength
	Toughness
	Initiative
	Agility
	Dexterity
	Intelligence
	Willpower
	Fellowship
)

type Template struct {
	id        ID
	name      string
	shortName string
}

func newTemplate(id ID, name string, shortName string) *Template {
	return &Template{id: id, name: name, shortName: shortName}
}

var TemplateValue = map[ID]*Template{
	WeaponSkill:    newTemplate(WeaponSkill, "Weapon Skill", "WS"),
	BallisticSkill: newTemplate(BallisticSkill, "Ballistic Skill", "BS"),
	Strength:       newTemplate(Strength, "Strength", "S"),
	Toughness:      newTemplate(Toughness, "Toughness", "T"),
	Initiative:     newTemplate(Initiative, "Initiative", "I"),
	Agility:        newTemplate(Agility, "Agility", "Ag"),
	Dexterity:      newTemplate(Dexterity, "Dexterity", "Dex"),
	Intelligence:   newTemplate(Intelligence, "Intelligence", "Int"),
	Willpower:      newTemplate(Willpower, "Willpower", "WP"),
	Fellowship:     newTemplate(Fellowship, "Fellowship", "Fel"),
}
