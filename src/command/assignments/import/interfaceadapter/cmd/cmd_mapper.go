package importcmdadapter

import (
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/hekmekk/git-team/src/command/adapter"
	"github.com/hekmekk/git-team/src/command/assignments/import/ldap/interfaceadapter/cmd"
)

// Command the add command
func Command(root commandadapter.CommandRoot) *kingpin.CmdClause {
	assignmentsImport := root.Command("import", "Import alias to co-author assignments")
	forceOverride := assignmentsImport.Flag("force-override", "force override").Short('f').Bool()

	importldapcmdadapter.Command(assignmentsImport, forceOverride)

	return assignmentsImport
}
