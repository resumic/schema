package cmd

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerateSchemaRunValid(t *testing.T) {
	tempFile, err := ioutil.TempFile(os.TempDir(), "resumic-schema-*.json")
	defer os.Remove(tempFile.Name())
	if err != nil {
		t.Fatalf("Could not generate temp file: %s", err)
	}
	err = Execute([]string{"generate", "schema", tempFile.Name()})
	if err != nil {
		t.Fatalf("Could not generate schema, %s", err)
	}
	r, err := compareJSONFiles("../schema.json", tempFile.Name())
	if err != nil {
		t.Fatalf("could not compare schema.json and generated schema: %s", err)
	}
	if !r {
		t.Fatal("schema.json and generated schema are not same")
	}
}
