package action

import (
	"github.com/Pretty-IT/wfrp-core/internal/models"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/weapons"
)

type Type string

const (
	Attack Type = "attack"
	Defend      = "defend"
)

type Request struct {
	actionType Type
	// only one of next 2 fields should be not null
	weapon *weapons.Value
	skill  *skills.Value
}

type Opposed struct {
	state    *charlist.State
	request  *Request
	modifier *models.Modifier
}
