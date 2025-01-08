package command

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func getLongAutocompleteText(rootCmdName string, shell string) string {
	switch shell {
	case "bash":
		return `Steps to setup command-line completion for bash:
1. Ensure you have bash-completion installed on your system

2. Execute once:
    ` + "```" + `
    ` + rootCmdName + ` completion bash > ~/.bash_completion.d/ncbi-` + rootCmdName + `.bash
    ` + "```" + `

3. Create/edit the file ` + "`" + `~/.bash_completion` + "`" + `, adding the following:
    ` + "```" + `
    for f in ~/.bash_completion.d/* ; do
        . $f
        [[ -f "$f" ]] && source "$f"
    done
    ` + "```" + `

Depending on your setup, you may need to source ` + "`" + `~/.bash_completion` + "`" + `
`
	case "zsh":
		return `Steps to setup command-line completion for zsh:

1. If shell completion is not already enabled in your environment, enable it by
executing the following once:
    ` + "```" + `
    echo "autoload -U compinit; compinit" >> ~/.zshrc
    ` + "```" + `

2. To load completions for each session, execute once:
    ` + "```" + `
    ` + rootCmdName + ` completion zsh > "${fpath[1]}/_ncbi_` + rootCmdName + `"
    ` + "```" + `

You will need to start a new shell for this setup to take effect.
`

	case "fish":
		return `To load completions for each session, execute once:
    ` + "```" + `
    ` + rootCmdName + ` completion fish > ~/.config/fish/completions/ncbi-` + rootCmdName + `.fish
    ` + "```" + `
`
	case "powershell":
		return `Generate powershell autocompletion script.

See the [golang cobra documentation](https://github.com/spf13/cobra/blob/master/powershell_completions.md) for details.
`
	}
	return ""
}

func shellCompletion(shell string, cmd *cobra.Command) error {
	switch shell {
	case "bash":
		return cmd.Root().GenBashCompletion(os.Stdout)
	case "zsh":
		return cmd.Root().GenZshCompletion(os.Stdout)
	case "fish":
		return cmd.Root().GenFishCompletion(os.Stdout, true)
	case "powershell":
		return cmd.Root().GenPowerShellCompletion(os.Stdout)
	}
	return errors.New("Internal error: unsupported shell")
}

func generateAutocompleteCmd(rootCmdName string, shell string) *cobra.Command {
	cmd := cobra.Command{
		Use:   shell,
		Short: fmt.Sprintf("Generate %s autocompletion script", shell),
		Long:  getLongAutocompleteText(rootCmdName, shell),
		RunE: func(cmd *cobra.Command, args []string) error {
			return shellCompletion(shell, cmd)
		},
		Args: cobra.MaximumNArgs(0),
	}
	return &cmd
}

func addAutocompletionShellCommands(rootCmdName string, cmd *cobra.Command) {
	autocompletes := []string{
		"bash",
		"zsh",
		"fish",
		"powershell",
	}
	for _, shell := range autocompletes {
		cmd.AddCommand(generateAutocompleteCmd(rootCmdName, shell))
	}
}

func NewAutocompleteCmd(rootCmd *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generate autocompletion scripts",
		Long: `This sub-command generates files needed to enable auto-complete for several popular command-line interpreters.

When enabled, the command-line interpreter can automatically fill in subcommands and options.

A good introduction of command-line completion is found on [wikipedia](https://en.wikipedia.org/wiki/Command-line_completion).`,
		Args: cobra.MaximumNArgs(0),
	}
	addAutocompletionShellCommands(rootCmd.Name(), cmd)
	return cmd

}
