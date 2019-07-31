package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/resumic/schema/schema"
	"github.com/spf13/cobra"
)

var schemaPath string

func generateJSONSchemaRun(cmd *cobra.Command, args []string) error {
	schema, err := schema.GenerateJSONSchema()
	if err != nil {
		return fmt.Errorf("Couldn't generate JSON schema: %s", err)
	}
	err = ioutil.WriteFile(schemaPath, schema, 0600)
	if err != nil {
		return fmt.Errorf("Couldn't write the schema to %s: %s", schemaPath, err)
	}
	fmt.Printf("JSON schema file created successfully at %s\n", schemaPath)
	return nil
}

var generateJSONSchemaCmd = &cobra.Command{
	Use:   "jsonschema",
	Short: "Generate the jsonschema for the resumic schema",
	Args:  cobra.MaximumNArgs(1),
	RunE:  generateJSONSchemaRun,
}

func init() {
	generateJSONSchemaCmd.Flags().StringVarP(&schemaPath, "output", "o", "resume.json", "Specify an output file for the resume data")
	generateCmd.AddCommand(generateJSONSchemaCmd)
}
