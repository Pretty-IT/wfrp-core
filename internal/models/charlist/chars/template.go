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
	id   ID
	name string
}

func newTemplate(id ID, name string) *Template {
	return &Template{id: id, name: name}
}

var Template_value = map[ID]*Template{
	WeaponSkill:    newTemplate(WeaponSkill, "Weapon Skill"),
	BallisticSkill: newTemplate(BallisticSkill, "Ballistic Skill"),
	Strength:       newTemplate(Strength, "Strength"),
	Toughness:      newTemplate(Toughness, "Toughness"),
	Initiative:     newTemplate(Initiative, "Initiative"),
	Agility:        newTemplate(Agility, "Agility"),
	Dexterity:      newTemplate(Dexterity, "Dexterity"),
	Intelligence:   newTemplate(Intelligence, "Intelligence"),
	Willpower:      newTemplate(Willpower, "Willpower"),
	Fellowship:     newTemplate(Fellowship, "Fellowship"),
}
