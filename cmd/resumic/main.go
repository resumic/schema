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
)

// GenerateSchema generates the schema.json file
func GenerateSchema(schemaFile string) {
	schema, err := jsonschema.NewSchema(schema.Schema{}, "", "Resumic Schema")
	if err != nil {
		log.Fatalf("Couldn't get the schema struct: %s", err)
	}
	json, err := json.Marshal(schema)
	if err != nil {
		log.Fatal("couldn't parse the schema")
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

func main() {
	doc := flag.String("doc", "../../examples/invalid/invalid_email.json", "Example file")
	schemaFile := flag.String("schema", "./schema.json", "Generate JSON Schema")
	flag.Parse()

	// Verify that a subcommand has been provided
	if len(flag.Args()) < 1 {
		fmt.Println("subcommand is required")
		os.Exit(1)
	}
	switch flag.Args()[0] {
	case "validate":
		validate.ValidateJSON(*doc)
	case "schema":
		GenerateSchema(*schemaFile)
	default:
		log.Fatalf("Unsupported subcommands. Please check --help for commands list")
	}
}
