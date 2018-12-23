package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
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

func compareFiles(file1, file2 string) (bool, error) {
	data1, err := ioutil.ReadFile(file1)
	if err != nil {
		return false, fmt.Errorf("could not read file1: %s", err)
	}
	data2, err := ioutil.ReadFile(file2)
	if err != nil {
		return false, fmt.Errorf("could not read file2: %s", err)
	}
	return bytes.Equal(data1, data2), nil
}
