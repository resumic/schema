package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/resumic/schema/jsonschema"
	"github.com/resumic/schema/schema"
	"github.com/spf13/cobra"
	"github.com/xeipuuv/gojsonschema"
)

func validateRun(cmd *cobra.Command, args []string) error {
	resumePath := args[0]
	schema, err := jsonschema.NewSchema(schema.Schema{}, "", "Resumic Schema")
	if err != nil {
		return fmt.Errorf("Couldn't generate schema: %s", err)
	}
	schemaJSON, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return fmt.Errorf("Couldn't marshal the schema to json: %s", err)
	}
	schemaLoader := gojsonschema.NewBytesLoader(schemaJSON)

	resumeJSON, err := ioutil.ReadFile(resumePath)
	if err != nil {
		return fmt.Errorf("Couldn't read the resume json file from %s: %s", resumePath, err)
	}
	resumeLoader := gojsonschema.NewBytesLoader(resumeJSON)

	result, err := gojsonschema.Validate(schemaLoader, resumeLoader)
	if err != nil {
		return fmt.Errorf("Couldn't validate the resume: %s", err)
	}
	if !result.Valid() {
		errors := ""
		for _, err := range result.Errors() {
			errors += "\n\t" + err.String()
		}
		return fmt.Errorf("%s is not a valid resume:%s", resumePath, errors)
	}
	fmt.Printf("%s is a valid resume json", resumePath)
	return nil
}

var validateCmd = &cobra.Command{
	Use:   "validate PATH",
	Short: "Validate a json resume",
	Args:  cobra.ExactArgs(1),
	RunE:  validateRun,
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
