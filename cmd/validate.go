package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/resumic/schema/schema"

	"github.com/spf13/cobra"
)

func validateRun(cmd *cobra.Command, args []string) error {
	var resumePath string
	if len(args) == 0 {
		resumePath = "resume.json"
	} else {
		resumePath = args[0]
	}

	resume, err := ioutil.ReadFile(resumePath)
	if err != nil {
		return fmt.Errorf("Couldn't read the resume json file from %s: %s", resumePath, err)
	}

	if err := schema.ValidateResume(resume); err != nil {
		if _, ok := err.(*schema.InvalidResumeError); ok {
			return err
		}
		return fmt.Errorf("Couldn't validate the resume: %s", err)
	}
	fmt.Printf("%s is valid and conforming with the resumic schema\n", resumePath)
	return nil
}

var validateCmd = &cobra.Command{
	Use:   "validate PATH",
	Short: "Validate a json resume",
	Args:  cobra.MaximumNArgs(1),
	RunE:  validateRun,
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
