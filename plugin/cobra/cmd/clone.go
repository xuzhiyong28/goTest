package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cloneCmd = &cobra.Command{
	Use: "clone",
	Short: "clone github",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprint(os.Stdout, "git clone.......")
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
