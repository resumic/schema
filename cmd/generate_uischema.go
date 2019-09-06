package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"go.resumic.org/schema/schema"
)

var uiSchemaPath string

func generateUISchemaRun(cmd *cobra.Command, args []string) error {
	schema, err := schema.GenerateUISchema()
	if err != nil {
		return fmt.Errorf("Couldn't generate UI schema: %s", err)
	}
	err = ioutil.WriteFile(uiSchemaPath, schema, 0600)
	if err != nil {
		return fmt.Errorf("Couldn't write the schema to %s: %s", uiSchemaPath, err)
	}
	fmt.Printf("UI schema file created successfully at %s\n", uiSchemaPath)
	return nil
}

var generateUISchemaCmd = &cobra.Command{
	Use:   "uischema",
	Short: "Generate the uischema for the resumic schema",
	Args:  cobra.MaximumNArgs(1),
	RunE:  generateUISchemaRun,
}

func init() {
	generateUISchemaCmd.Flags().StringVarP(&uiSchemaPath, "output", "o", "resume.json", "Specify an output file for the resume data")
	generateCmd.AddCommand(generateUISchemaCmd)
}
