package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// Version Contains the version number for a particular Resumic build
	Version string
	// GitCommit contains the commit that this build is based on
	GitCommit = "none"
)

const packageName string = "Resumic"

func check(e error) bool {
	return e != nil
}

func versionFunc(cmd *cobra.Command, args []string) error {
	// pwd, err := os.Getwd()
	// dat, err := ioutil.ReadFile(filepath.Join(pwd, `VERSION`))
	// if check(err) {
	// 	return err
	// }
	fmt.Println(packageName + ", " + Version)
	fmt.Println("running on...")
	fmt.Println(runtime.GOOS + ", " + runtime.Version())
	fmt.Println("Git Info: " + GitCommit)
	return nil
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information for this build of Schema",
	Args:  cobra.ExactArgs(0),
	RunE:  versionFunc,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
