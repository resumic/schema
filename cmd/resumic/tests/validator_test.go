package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/resumic/schema/cmd/resumic/validate"
)

//TestValidatorValidExamples used to test valid JSON passed in ValidateJSON
func TestValidatorValidExamples(t *testing.T) {
	files := run("../../../examples/valid")
	for _, file := range files {
		fmt.Printf("In %s:", file)
		result := validate.ValidateJSON(file)
		if result != true {
			t.Errorf("Expected valid state, got error")
		}
	}

}

//TestValidatorInValidExamples used to test valid JSON passed in ValidateJSON
func TestValidatorInvalidExamples(t *testing.T) {
	files := run("../../../examples/invalid")
	for _, file := range files {
		fmt.Printf("In %s:", file)
		result := validate.ValidateJSON(file)
		if result != false {
			t.Errorf("Expected invalid state, got valid")
		}

	}

}

func run(searchDir string) []string {
	ext := ".json"
	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == ext {
			fileList = append(fileList, path)
		}

		return err
	})

	if e != nil {
		panic(e)
	}

	return fileList
}
