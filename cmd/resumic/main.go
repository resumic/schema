package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/resumic/schema/hugo"
	"github.com/resumic/schema/jsonschema"
	"github.com/resumic/schema/schema"
	"github.com/xeipuuv/gojsonschema"
)

// ValidateResume validate a resume file
func ValidateResume(resumeFile string) error {
	schema, err := jsonschema.NewSchema(schema.Schema{}, "", "Resumic Schema")
	if err != nil {
		return fmt.Errorf("couldn't get the schema struct: %s", err)
	}
	schemaJSON, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return fmt.Errorf("couldn't marshal the schema")
	}
	schemaLoader := gojsonschema.NewBytesLoader(schemaJSON)

	resumeJSON, err := ioutil.ReadFile(resumeFile)
	if err != nil {
		return fmt.Errorf("couldn't read the resume file: %s", err)
	}
	resumeLoader := gojsonschema.NewBytesLoader(resumeJSON)

	result, err := gojsonschema.Validate(schemaLoader, resumeLoader)
	if err != nil {
		return fmt.Errorf("couldn't validate the resume: %s", err)
	}
	if !result.Valid() {
		errors := ""
		for _, err := range result.Errors() {
			errors += "\n" + err.String()
		}
		log.Fatalf("%s is not a valid resume:%s", resumeFile, errors)
	}
	return nil
}

// GenerateSchema generates the schema.json file
func GenerateSchema(schemaFile string) error {
	schema, err := jsonschema.NewSchema(schema.Schema{}, "", "Resumic Schema")
	if err != nil {
		return fmt.Errorf("couldn't get the schema struct: %s", err)
	}
	json, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return fmt.Errorf("couldn't marshal the schema")
	}
	file, err := os.Create(schemaFile)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("couldn't create the schema file")
	}
	_, err = file.Write(json)
	if err != nil {
		return fmt.Errorf("couldn't write the schema content to given the schema file")
	}
	return nil
}

// GenerateExample generate the example file
func GenerateExample(exampleFile string) error {
	example := schema.NewExample()
	json, err := json.MarshalIndent(example, "", "  ")
	if err != nil {
		return fmt.Errorf("couldn't parse the example")
	}
	file, err := os.Create(exampleFile)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("couldn't create the example file")
	}
	_, err = file.Write(json)
	if err != nil {
		return fmt.Errorf("couldn't write the example content to given the schema file")
	}
	return nil
}

func RenderResume() error {
	hugo, err := hugo.New("/home/arman/test-resumic/site")
	example, err := json.MarshalIndent(schema.NewExample(), "", "  ")
	hugo.NewResume(example)
	return err
}

func main() {
	resumeFile := flag.String("resume", "", "Resume file")
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
		err := ValidateResume(*resumeFile)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("resume is valid.")
	case "schema":
		err := GenerateSchema(*schemaFile)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Schema file created successfully")
	case "example":
		err := GenerateExample(*exampleFile)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Example file created successfully")
	case "render":
		err := RenderResume()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Resume HTML rendered successfully")
	default:
		log.Fatalln("Unsupported subcommands. Please check --help for commands list")
	}
}
