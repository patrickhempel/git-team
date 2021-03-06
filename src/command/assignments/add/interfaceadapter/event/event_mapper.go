package addeventadapter

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/hekmekk/git-team/src/command/assignments/add"
	"github.com/hekmekk/git-team/src/core/effects"
	"github.com/hekmekk/git-team/src/core/events"
)

// MapEventToEffects convert assignment events to effects for the cli
func MapEventToEffects(event events.Event) []effects.Effect {
	switch evt := event.(type) {
	case add.AssignmentSucceeded:
		return []effects.Effect{
			effects.NewPrintMessage(color.CyanString(fmt.Sprintf("Assignment added: '%s' →  '%s'", evt.Alias, evt.Coauthor))),
			effects.NewExitOk(),
		}
	case add.AssignmentFailed:
		return []effects.Effect{
			effects.NewPrintErr(evt.Reason),
			effects.NewExitErr(),
		}
	case add.AssignmentAborted:
		return []effects.Effect{
			effects.NewPrintMessage(color.YellowString("Nothing changed")),
			effects.NewExitOk(),
		}
	default:
		return []effects.Effect{}
	}
}
