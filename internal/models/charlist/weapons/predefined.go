package weapons

import (
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/chars"
	"github.com/Pretty-IT/wfrp-core/internal/models/charlist/skills"
)

var Punches = New("Punches", skills.MeleeBrawling, 0, chars.Strength)
