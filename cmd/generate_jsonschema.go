package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/resumic/schema/schema"
	"github.com/spf13/cobra"
)

func generateJSONSchemaRun(cmd *cobra.Command, args []string) error {
	var schemaPath string
	if len(args) == 0 {
		schemaPath = "resume.json"
	} else {
		schemaPath = args[0]
	}
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
	Use:   "jsonschema PATH",
	Short: "Generate the jsonschema for the resumic schema",
	Args:  cobra.MaximumNArgs(1),
	RunE:  generateJSONSchemaRun,
}

func init() {
	generateCmd.AddCommand(generateJSONSchemaCmd)
}
