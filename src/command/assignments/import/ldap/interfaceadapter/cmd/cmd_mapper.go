package importldapcmdadapter

import (
	"bufio"
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/hekmekk/git-team/src/command/adapter"
	add "github.com/hekmekk/git-team/src/command/assignments/import/ldap"
	addeventadapter "github.com/hekmekk/git-team/src/command/assignments/import/ldap/interfaceadapter/event"
	"github.com/hekmekk/git-team/src/core/gitconfig"
	"github.com/hekmekk/git-team/src/core/validation"
)

// Command the add command
func Command(root commandadapter.CommandRoot, forceOverride *bool) *kingpin.CmdClause {
	ldap := root.Command("ldap", "Import assignments from ldap")
	configDir := ldap.Flag("config-dir", "Directory where ldap config is located").Short('c').Required().String()

	ldap.Action(commandadapter.Run(policy(configDir, forceOverride), addeventadapter.MapEventToEffects))

	return ldap
}

func policy(configDir *string, forceOverride *bool) add.Policy {
	return add.Policy{
		Req: add.AssignmentRequest{
			Alias:         nil,
			Coauthor:      nil,
			ForceOverride: forceOverride,
		},
		Deps: add.Dependencies{
			SanityCheckCoauthor: validation.SanityCheckCoauthor,
			GitResolveAlias:     commandadapter.ResolveAlias,
			GitAddAlias: func(alias string, coauthor string) error {
				return gitconfig.ReplaceAll(fmt.Sprintf("team.alias.%s", alias), coauthor)
			},
			GetAnswerFromUser: func(question string) (string, error) {
				_, err := os.Stdout.WriteString(question)
				if err != nil {
					return "", err
				}
				return bufio.NewReader(os.Stdin).ReadString('\n')
			},
		},
	}
}
