package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/resumic/schema/schema"
	"github.com/spf13/cobra"
)

func generateExampleRun(cmd *cobra.Command, args []string) error {
	var examplePath string
	if len(args) == 0 {
		examplePath = "resume_example.json"
	} else {
		examplePath = args[0]
	}
	example, err := schema.GenerateExample()
	if err != nil {
		return fmt.Errorf("Couldn't generate the example: %s", err)
	}
	exampleJSON, err := json.MarshalIndent(example, "", "  ")
	if err != nil {
		return fmt.Errorf("Couldn't marshal the example to json: %s", err)
	}
	err = ioutil.WriteFile(examplePath, exampleJSON, 0600)
	if err != nil {
		return fmt.Errorf("Couldn't write the example to %s: %s", examplePath, err)
	}
	fmt.Printf("Example file created successfully at %s\n", examplePath)
	return nil
}

var generateExampleCmd = &cobra.Command{
	Use:   "example PATH",
	Short: "Generate an example json resume",
	Args:  cobra.MaximumNArgs(1),
	RunE:  generateExampleRun,
}

func init() {
	generateCmd.AddCommand(generateExampleCmd)
}
