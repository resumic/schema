package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/resumic/schema/schema"
	"github.com/spf13/cobra"
)

func generateUISchemaRun(cmd *cobra.Command, args []string) error {
	var schemaPath string
	if len(args) == 0 {
		schemaPath = "resume.json"
	} else {
		schemaPath = args[0]
	}
	schema, err := schema.GenerateUISchema()
	if err != nil {
		return fmt.Errorf("Couldn't generate UI schema: %s", err)
	}
	err = ioutil.WriteFile(schemaPath, schema, 0600)
	if err != nil {
		return fmt.Errorf("Couldn't write the schema to %s: %s", schemaPath, err)
	}
	fmt.Printf("UI schema file created successfully at %s\n", schemaPath)
	return nil
}

var generateUISchemaCmd = &cobra.Command{
	Use:   "uischema PATH",
	Short: "Generate the uischema for the resumic schema",
	Args:  cobra.MaximumNArgs(1),
	RunE:  generateUISchemaRun,
}

func init() {
	generateCmd.AddCommand(generateUISchemaCmd)
}
