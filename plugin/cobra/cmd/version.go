package cmd
import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)
var versionCmd = &cobra.Command {
	Use: "version",
	Short: "version subcommand show git version info.",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := ExecuteCommand("git", "version", args...)
		if err != nil {
			Error(cmd, args, err)
		}
		fmt.Fprint(os.Stdout, output)
	},
}


var cloneCmd = &cobra.Command{
	Use: "clone",
	Short: "clone github",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprint(os.Stdout, "git clone.......")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(cloneCmd)
}
