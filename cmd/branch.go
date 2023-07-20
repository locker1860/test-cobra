package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{	
	Use: "branch",
	Short: "branch subcommand show git branch list.",

	Run: func(cmd *cobra.Command,args []string) {
		out, err := ExecuteCommand("git", "branch", args...)
		if err != nil {
			Error(cmd, args,err)
		}
		fmt.Fprint(os.Stdout, out)
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
}
