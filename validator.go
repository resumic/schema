package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/gobuffalo/packr"
	"github.com/xeipuuv/gojsonschema"
)

func main() {
	var doc string
	flag.StringVar(&doc, "doc", "examples/invalid/invalid_email.json", "Example file")
	flag.Parse()
	ValidateJSON(doc)
}

// ValidateJSON is used to check for validity
func ValidateJSON(doc string) bool {

	file := path.Join("file:///", GetPath(), "/", doc)
	box := packr.NewBox("./")
	s := box.String("schema.json")
	schemaLoader := gojsonschema.NewStringLoader(s)
	documentLoader := gojsonschema.NewReferenceLoader(file)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
		return true
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
		return false
	}

}

// GetPath is used to get current path
func GetPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
