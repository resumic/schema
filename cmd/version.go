package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// Version Contains the version number for a particular Resumic build
	version = "dev"
	// GitCommit contains the commit that this build is based on
	gitCommit = "none"
)

const packageName string = "Resumic"

func versionFunc(cmd *cobra.Command, args []string) error {
	fmt.Println(packageName + "-" + version)
	fmt.Println(runtime.GOOS + ", " + runtime.Version() + ", " + gitCommit)
	return nil
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information for this build of Resumic",
	Args:  cobra.ExactArgs(0),
	RunE:  versionFunc,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
