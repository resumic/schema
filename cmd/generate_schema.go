package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/resumic/schema/jsonschema"
	"github.com/resumic/schema/schema"
	"github.com/spf13/cobra"
)

func generateSchemaRun(cmd *cobra.Command, args []string) error {
	schemaPath := args[0]
	schema, err := jsonschema.NewSchema(schema.Schema{}, "", "Resumic Schema")
	if err != nil {
		return fmt.Errorf("Couldn't generate schema: %s", err)
	}
	schemaJSON, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return fmt.Errorf("Couldn't marshal the schema to json: %s", err)
	}
	err = ioutil.WriteFile(schemaPath, schemaJSON, 0600)
	if err != nil {
		return fmt.Errorf("Couldn't write the schema to %s: %s", schemaPath, err)
	}
	fmt.Printf("Schema file created successfully at %s\n", schemaPath)
	return nil
}

var generateSchemaCmd = &cobra.Command{
	Use:   "schema PATH",
	Short: "Generate the jsonschema for the resume schema",
	Args:  cobra.ExactArgs(1),
	RunE:  generateSchemaRun,
}

func init() {
	generateCmd.AddCommand(generateSchemaCmd)
}
