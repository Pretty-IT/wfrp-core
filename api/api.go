package api

import (
	"fmt"
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/action"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
)

func HelloWorld() string {
	msg := "Hello from Go!"
	fmt.Println(msg)
	return msg
}

func CalculateModifier(actorState *charlist.State, targetState *charlist.State,
	actionRequest action.Request, environment models.Environment) *models.Modifier {
	return &models.Modifier{}
}

func MakeDramaticTest(state *charlist.State, modifier *models.Modifier, roll int) (*models.TestResult, []models.TestModifyOption) {
	return &models.TestResult{}, []models.TestModifyOption{}
}

func MakeOpposedTest(attacker *action.Opposed, defender *action.Opposed) *models.OpposedTestResult {
	return &models.OpposedTestResult{}
}

func MakeAttack(attacker *action.Opposed, defender *action.Opposed) (*models.OpposedTestResult, []*models.ResolveAttackOption) {
	opposedResult := MakeOpposedTest(attacker, defender)
	resolveOptions := resolveAttack(attacker, defender, opposedResult)
	return opposedResult, resolveOptions
}

func resolveAttack(attacker *action.Opposed, defender *action.Opposed, result *models.OpposedTestResult) []*models.ResolveAttackOption {
	return []*models.ResolveAttackOption{}
}

func ChangeState(state *charlist.State, delta *charlist.Delta) *charlist.State {
	result := &charlist.State{} //Create deep copy here
	//apply delta to result
	return result
}
