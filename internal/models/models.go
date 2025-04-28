package models

type Environment struct{}

type ResolveAttackOption struct {
	_type ResolveAttackOptionType
}

type ResolveAttackOptionType string

const (
	Damage ResolveAttackOptionType = "damage"
)
