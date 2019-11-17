package assignmentscmdadapter

import (
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/hekmekk/git-team/src/command/adapter"
	"github.com/hekmekk/git-team/src/command/assignments/add/interfaceadapter/cmd"
	"github.com/hekmekk/git-team/src/command/assignments/import/interfaceadapter/cmd"
	"github.com/hekmekk/git-team/src/command/assignments/list/interfaceadapter/cmd"
	"github.com/hekmekk/git-team/src/command/assignments/remove/interfaceadapter/cmd"
)

// Command the assignments command
func Command(root commandadapter.CommandRoot) *kingpin.CmdClause {
	assignments := root.Command("assignments", "Manage your alias -> co-author assignments")

	assignmentsLs := listcmdadapter.Command(assignments)
	assignmentsLs.Default()

	addcmdadapter.Command(assignments)

	removecmdadapter.Command(assignments)

	importcmdadapter.Command(assignments)

	return assignments
}
