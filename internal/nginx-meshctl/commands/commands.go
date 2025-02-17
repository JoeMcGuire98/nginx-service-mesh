// Package commands contains all of the cli commands
package commands // import "github.com/nginxinc/nginx-service-mesh/internal/nginx-meshctl/commands"

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	// client-go auth methods.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// Setup creates the root command and adds all sub commands.
func Setup(cmdName, version, commit string) *cobra.Command {
	return AddSubCommands(cmdName, version, commit, Root(cmdName))
}

// AddSubCommands adds all subcommands to the root command.
func AddSubCommands(cmdName, version, commit string, rootCmd *cobra.Command) *cobra.Command {
	rootCmd.AddCommand(NewStatusCmd())
	rootCmd.AddCommand(NewVersionCmd(cmdName, version, commit))
	rootCmd.AddCommand(Top())
	rootCmd.AddCommand(GetServices())
	rootCmd.AddCommand(GetConfig())
	rootCmd.AddCommand(Inject())
	rootCmd.AddCommand(Deploy())
	rootCmd.AddCommand(Upgrade(version))
	rootCmd.AddCommand(Remove())
	rootCmd.AddCommand(Support(version))

	return rootCmd
}

var errCommandStopped = errors.New("command stopped by user")

// ReadYes reads user input to make a yes/no decision.
func ReadYes(msg string) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(msg)
	fmt.Print("Do you want to continue (y/n)? ")
	letter, _, _ := reader.ReadRune()
	switch letter {
	case 'Y', 'y':
	default:
		fmt.Println()

		return errCommandStopped
	}
	fmt.Println()

	return nil
}

/* TabWriterWithOpts returns a tabwriter.
 * This call, with these numbers were found across the codebase
 * and so were centralized here so that modifications to text attributes
 * could be made from one place.*/
func TabWriterWithOpts() *tabwriter.Writer {
	return tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0) //nolint:gomnd // ignore text opts
}
