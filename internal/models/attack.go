package models

import "fmt"

type AttackResult struct {
	Damage *AttackDamage
}

func (ar *AttackResult) String() string {
	return fmt.Sprintf(`Attack Result:
	%v
`,
		ar.Damage)
}

type AttackDamage struct {
	FromWeapon int
	FromSL     int
}

func (ad *AttackDamage) Total() int {
	return ad.FromWeapon + ad.FromSL
}

func (ad *AttackDamage) String() string {
	return fmt.Sprintf(`Attack Damage:
	From Weapon: %d
	From SL: %d
`,
		ad.FromWeapon, ad.FromSL)
}
