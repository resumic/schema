package cmd

import (
	"os"
	"path"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "resumic",
	Short: "Resumic is a standardized and generic data schema for your resume/CV",
}

func init() {
	userCache, err := os.UserCacheDir()
	if err != nil {
		userCache = path.Join(os.TempDir(), "cache")
	}
	cacheDir := path.Join(userCache, "resumic")
	rootCmd.PersistentFlags().String("cacheDir", cacheDir, "Filesystem path to cache directory.")
}

func Execute() {
	rootCmd.Execute()
}
