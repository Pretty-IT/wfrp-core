package actions

import (
	"fmt"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
)

type Opposed struct {
	State       *charlist.State
	RollRequest *RollRequest
	Roll        *Roll
	RollResult  *RollResult
}

func (o *Opposed) String() string {
	return fmt.Sprintf(`Opposed Test Request:
	State: %v 
	RollRequest: %v
	Roll: %v
	RollResult: %v
`, o.State, o.RollRequest, o.Roll, o.RollResult)
}

type OpposedResult struct {
	IsAttackerWinner bool
	IsTie            bool
	SL               int
}

func (r *OpposedResult) String() string {
	return fmt.Sprintf(`Opposed Test Result:
	Is Attacker Winner: %t,
	Is Tie: %t,
	Opposed SL: %d
`,
		r.IsAttackerWinner, r.IsTie, r.SL)
}
