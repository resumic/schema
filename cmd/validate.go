package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/resumic/schema/schema"

	"github.com/spf13/cobra"
)

var resumePathValidate string

func validateRun(cmd *cobra.Command, args []string) error {

	resume, err := ioutil.ReadFile(resumePathValidate)
	if err != nil {
		return fmt.Errorf("Couldn't read the resume json file from %s: %s", resumePathValidate, err)
	}

	if err := schema.ValidateResume(resume); err != nil {
		if _, ok := err.(*schema.InvalidResumeError); ok {
			return err
		}
		return fmt.Errorf("Couldn't validate the resume: %s", err)
	}
	fmt.Printf("%s is valid and conforming with the resumic schema\n", resumePathValidate)
	return nil
}

var validateCmd = &cobra.Command{
	Use:   "validate PATH",
	Short: "Validate a json resume",
	Args:  cobra.MaximumNArgs(1),
	RunE:  validateRun,
}

func init() {
	validateCmd.Flags().StringVarP(&resumePathValidate, "resume", "r", "resume.json", "Path to resume data file")
	rootCmd.AddCommand(validateCmd)
}
