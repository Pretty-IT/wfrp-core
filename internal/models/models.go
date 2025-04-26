package models

type Modifier struct {
}

type Environment struct{}

type TestResult struct {
}

type TestModifyOption string

const (
	Reroll  TestModifyOption = "reroll"
	Reverse TestModifyOption = "reverse"
	Select  TestModifyOption = "select"
)

type OpposedTestResult struct {
}

type ResolveAttackOption struct {
	_type ResolveAttackOptionType
}

type ResolveAttackOptionType string

const (
	Damage ResolveAttackOptionType = "damage"
)
