package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "execute %s args:%v error: %v\n", cmd.Name(), args, err)
	os.Exit(1)
}

func ExecuteCommand(name string, subname string, args ...string) (string, error) {
	args = append([]string{subname}, args...)
	cmd := exec.Command(name, args...)
	bs, err := cmd.CombinedOutput()

	return string(bs), err
}
