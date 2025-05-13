package actions

import (
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
)

type Opposed struct {
	State       *charlist.State
	RollRequest *RollRequest
	RollResult  *RollResult
}

type OpposedResult struct {
	IsAttackerWinner bool
	IsTie            bool
	SL               int
}
