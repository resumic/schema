package main

import (
	"flag"

	"github.com/resumic/schema/cmd/resumic/validate"
)

func main() {
	var doc string
	flag.StringVar(&doc, "doc", "../../examples/invalid/invalid_email.json", "Example file")
	flag.Parse()
	validate.ValidateJSON(doc)
}
