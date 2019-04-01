package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func myTest(cmd *cobra.Command, args []string) error {
	fmt.Println(args[1])
	return nil
}

var testCmd = &cobra.Command{
	Use:   "my_test",
	Short: "Run my test command!",
	Args:  cobra.ExactArgs(1),
	RunE:  myTest,
}

func init() {
	rootCmd.AddCommand(testCmd)
}
