package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const packageName string = "Resumic"

func check(e error) bool {
	return e != nil
}

func versionFunc(cmd *cobra.Command, args []string) error {
	pwd, err := os.Getwd()
	dat, err := ioutil.ReadFile(filepath.Join(pwd, `VERSION`))
	if check(err) {
		return err
	}
	fmt.Println(packageName)
	fmt.Print(string(dat))
	return nil
}

var versionCmd *cobra.Command = &cobra.Command{
	Use:   "version",
	Short: "Show version information for this build of Schema",
	Args:  cobra.ExactArgs(0),
	RunE:  versionFunc,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
