package actions

import (
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
	Skill  skills.ID
	Weapon *weapons.Value
}

type RollModifyOption string

const (
	Reroll  RollModifyOption = "reroll"
	Reverse RollModifyOption = "reverse"
	Select  RollModifyOption = "select"
)

type Roll struct {
	TargetCharValue int
	TargetValue     int
	TotalModifier   int
	TotalSLModifier int
	ModifyOptions   []RollModifyOption
}

type RollResult struct {
	IsPassed bool
	SL       int
}
