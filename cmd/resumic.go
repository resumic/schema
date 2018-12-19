package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "resumic",
	Short: "Resumic is a standardized and generic data schema for your resume/CV",
}

func Execute() {
	rootCmd.Execute()
}
