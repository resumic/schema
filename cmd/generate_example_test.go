package cmd

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerateExampleRunValid(t *testing.T) {
	tempFile, err := ioutil.TempFile(os.TempDir(), "resumic-example-*.json")
	defer os.Remove(tempFile.Name())
	if err != nil {
		t.Fatalf("Could not generate temp file: %s", err)
	}
	err = Execute([]string{"generate", "example", tempFile.Name()})
	if err != nil {
		t.Fatalf("Could not generate example, %s", err)
	}
	r, err := compareJSONFiles("../example.json", tempFile.Name())
	if err != nil {
		t.Fatalf("could not compare example.json and generated schema: %s", err)
	}
	if !r {
		t.Fatal("example.json and generated example are not same")
	}
}
