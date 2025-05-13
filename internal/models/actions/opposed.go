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
	isAttackerWinner bool
	isTie            bool
	SL               int
}
