package removecmdadapter

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/hekmekk/git-team/src/command/adapter"
	"github.com/hekmekk/git-team/src/command/assignments/remove"
	"github.com/hekmekk/git-team/src/command/assignments/remove/interfaceadapter/event"
	"github.com/hekmekk/git-team/src/core/gitconfig"
)

// Command the rm command
func Command(root commandadapter.CommandRoot) *kingpin.CmdClause {
	rm := root.Command("rm", "Remove an alias to co-author assignment")
	alias := rm.Arg("alias", "The alias to identify the assignment to be removed").Required().String()

	rm.Action(commandadapter.Run(policy(alias), removeeventadapter.MapEventToEffects))

	return rm
}

func policy(alias *string) remove.Policy {
	return remove.Policy{
		Req: remove.DeAllocationRequest{
			Alias: alias,
		},
		Deps: remove.Dependencies{
			GitRemoveAlias: func(alias string) error {
				return gitconfig.UnsetAll(fmt.Sprintf("team.alias.%s", alias))
			},
		},
	}
}
