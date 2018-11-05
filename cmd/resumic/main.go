package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/resumic/schema/cmd/resumic/validate"
	"github.com/resumic/schema/jsonschema"
	"github.com/resumic/schema/schema"
	"github.com/xeipuuv/gojsonschema"
)

// ValidateResume validate a resume file
func ValidateResume(resumeFile string) {
	schema, err := jsonschema.NewSchema(schema.Schema{}, "", "Resumic Schema")
	if err != nil {
		log.Fatalf("couldn't get the schema struct: %s", err)
	}
	schemaJSON, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		log.Fatal("couldn't marshal the schema")
	}
	schemaLoader := gojsonschema.NewBytesLoader(schemaJSON)
	resumeLoader := gojsonschema.NewReferenceLoader(resumeFile)
	result, err := gojsonschema.Validate(schemaLoader, resumeLoader)
	if err != nil {
		log.Fatalf("couldn't validate the resume: %s", err)
	}
	if result.Valid() {
		log.Printf("%s is a valid resume.", resumeFile)
	} else {
		log.Fatalf("%s is not a valid resume", resumeFile)
		for _, desc := range result.Errors() {
			log.Fatal(desc)
		}
	}
}

// GenerateSchema generates the schema.json file
func GenerateSchema(schemaFile string) {
	schema, err := jsonschema.NewSchema(schema.Schema{}, "", "Resumic Schema")
	if err != nil {
		log.Fatalf("couldn't get the schema struct: %s", err)
	}
	json, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		log.Fatal("couldn't marshal the schema")
	}
	file, err := os.Create(schemaFile)
	defer file.Close()
	if err != nil {
		log.Fatal("couldn't create the schema file")
	}
	_, err = file.Write(json)
	if err != nil {
		log.Fatal("couldn't write the schema content to given the schema file")
	}
	log.Printf("Schema file created successfully: %s", file.Name())
}

// GenerateExample generate the example file
func GenerateExample(exampleFile string) {
	example := schema.NewExample()
	json, err := json.MarshalIndent(example, "", "  ")
	if err != nil {
		log.Fatal("couldn't parse the example")
	}
	file, err := os.Create(exampleFile)
	defer file.Close()
	if err != nil {
		log.Fatal("couldn't create the example file")
	}
	_, err = file.Write(json)
	if err != nil {
		log.Fatal("couldn't write the example content to given the schema file")
	}
	log.Printf("Example file created successfully: %s", file.Name())
}

func main() {
	resume := flag.String("resume", "", "Resume file")
	schemaFile := flag.String("schema", "./schema.json", "Generate JSON Schema")
	exampleFile := flag.String("example", "./example.json", "Generate example JSON")
	flag.Parse()

	// Verify that a subcommand has been provided
	if len(flag.Args()) < 1 {
		fmt.Println("subcommand is required")
		os.Exit(1)
	}
	switch flag.Args()[0] {
	case "validate":
		validate.ValidateJSON(*resume)
	case "schema":
		GenerateSchema(*schemaFile)
	case "example":
		GenerateExample(*exampleFile)
	default:
		log.Fatalf("Unsupported subcommands. Please check --help for commands list")
	}
}
