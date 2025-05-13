package actions

import (
	"fmt"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/weapons"
)

type Type string

const (
	Attack Type = "attack"
	Defend      = "defend"
)

type RollRequest struct {
	ActionType Type

	// only one of next 2 fields should be not null
	SkillId skills.ID
	Weapon  *weapons.Value
}

func (r *RollRequest) GetSkillID() skills.ID {
	result := r.SkillId
	if r.Weapon != nil {
		result = r.Weapon.Skill
	}
	return result
}

func (r *RollRequest) String() string {
	return fmt.Sprintf(`
	Action : %s,
	Skill  : %s,
`,
		r.ActionType, skills.TemplateValue[r.GetSkillID()].FullName())
}

type RollModifyOption string

const (
	Reroll  RollModifyOption = "reroll"
	Reverse RollModifyOption = "reverse"
	Select  RollModifyOption = "select"
)

type Roll struct {
	TargetValue     int
	TotalModifier   int
	TotalSLModifier int
	ModifyOptions   []RollModifyOption
}

func (r *Roll) String() string {
	return fmt.Sprintf(`
	Target Value : %d,
	Total Modifier : %d,
	Total SL Modifier : %d,
`,
		r.TargetValue, r.TotalModifier, r.TotalSLModifier,
	)
}

type RollResult struct {
	Dice       int
	IsPassed   bool
	SL         int
	IsCritical bool
	IsFumble   bool
}

func (r *RollResult) String() string {
	return fmt.Sprintf(`
	Dice : %d,
	IsPassed : %t,
	SL : %d,
	IsCritical : %t,
	IsFumble : %t
`,
		r.Dice, r.IsPassed, r.SL, r.IsCritical, r.IsFumble)
}
