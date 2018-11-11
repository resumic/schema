package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func compareJSONFiles(file1, file2 string) (bool, error) {
	data1, err := ioutil.ReadFile(file1)
	if err != nil {
		return false, fmt.Errorf("could not read file1: %s", err)
	}
	data2, err := ioutil.ReadFile(file2)
	if err != nil {
		return false, fmt.Errorf("could not read file2: %s", err)
	}
	var json1 interface{}
	err = json.Unmarshal(data1, &json1)
	if err != nil {
		return false, fmt.Errorf("could not unmarshal json1: %s", err)
	}
	var json2 interface{}
	err = json.Unmarshal(data2, &json2)
	if err != nil {
		return false, fmt.Errorf("could not unmarshal json2: %s", err)
	}
	return reflect.DeepEqual(json1, json2), nil
}

func TestValidateResumeValid(t *testing.T) {
	resumeFile := "../../example.json"
	err := ValidateResume(resumeFile)
	if err != nil {
		t.Fatalf("Expected valid state, got error, %s", err)
	}
}

func TestGenerateSchemaValid(t *testing.T) {
	tempFile, err := ioutil.TempFile(os.TempDir(), "schema-*.json")
	defer os.Remove(tempFile.Name())
	if err != nil {
		t.Fatalf("Could not generate temp file: %s", err)
	}
	err = GenerateSchema(tempFile.Name())
	if err != nil {
		t.Fatalf("Could not generate schema, %s", err)
	}
	if err != nil {
		t.Fatalf("Could not read schema.json, %s", err)
	}
	r, err := compareJSONFiles("../../schema.json", tempFile.Name())
	if err != nil {
		t.Fatalf("could not compare schema.json and generated schema: %s", err)
	}
	if !r {
		t.Fatal("schema.json and generated schema are not same")
	}
}

func TestGenerateExampleValid(t *testing.T) {
	tempFile, err := ioutil.TempFile(os.TempDir(), "schema-*.json")
	defer os.Remove(tempFile.Name())
	if err != nil {
		t.Fatalf("Could not generate temp file: %s", err)
	}
	err = GenerateExample(tempFile.Name())
	if err != nil {
		t.Fatalf("Could not generate example, %s", err)
	}
	if err != nil {
		t.Fatalf("Could not read example.json, %s", err)
	}
	r, err := compareJSONFiles("../../example.json", tempFile.Name())
	if err != nil {
		t.Fatalf("could not compare example.json and generated schema: %s", err)
	}
	if !r {
		t.Fatal("example.json and generated example are not same")
	}
}
