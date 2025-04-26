package chars

type CharacteristicName string

const (
	WeaponSkill    CharacteristicName = "Weapon Skill"
	BallisticSkill                    = "Ballistic Skill"
	Strength                          = "Strength"
	Toughness                         = "Toughness"
	Initiative                        = "Initiative"
	Agility                           = "Agility"
	Dexterity                         = "Dexterity"
	Intelligence                      = "Intelligence"
	Willpower                         = "Willpower"
	Fellowship                        = "Fellowship"
)

type Characteristic struct {
	name    CharacteristicName
	initial int
}

func New(name CharacteristicName, initial int) *Characteristic {
	return &Characteristic{
		name:    name,
		initial: initial,
	}
}

func (c *Characteristic) Name() CharacteristicName { return c.name }

func (c *Characteristic) Initial() int { return c.initial }

func (c *Characteristic) Total() int { return c.initial }

func (c *Characteristic) Bonus() int { return c.Total() / 10 }

func (c *Characteristic) String() string {
	return string(c.name)
}
