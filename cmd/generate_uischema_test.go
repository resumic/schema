package cmd

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerateUISchemaRunValid(t *testing.T) {
	tempFile, err := ioutil.TempFile(os.TempDir(), "resumic-schema-*.json")
	defer os.Remove(tempFile.Name())
	if err != nil {
		t.Fatalf("Could not generate temp file: %s", err)
	}
	err = Execute([]string{"generate", "uischema", "-o", tempFile.Name()})
	if err != nil {
		t.Fatalf("Could not generate schema, %s", err)
	}
	r, err := compareJSONFiles("../ui.json", tempFile.Name())
	if err != nil {
		t.Fatalf("could not compare ui.json and generated schema: %s", err)
	}
	if !r {
		t.Fatal("ui.json and generated schema are not same")
	}
}
