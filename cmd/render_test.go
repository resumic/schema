package cmd

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestRenderRun(t *testing.T) {
	resumePath := "../example.json"
	htmlPath := "../example.html"
	tempFile, err := ioutil.TempFile(os.TempDir(), "resumic-html-*.json")
	if err != nil {
		t.Fatalf("Could not make temp file: %s", err)
	}
	err = renderRun(renderCmd, []string{resumePath, tempFile.Name()})
	if err != nil {
		t.Fatalf("Could not render the resume html: %s", err)
	}
	r, err := compareFiles(htmlPath, tempFile.Name())
	if err != nil {
		t.Fatalf("Could not compare the html and temp file: %s", err)
	}
	if !r {
		t.Fatalf("%s and generated html are not same", htmlPath)
	}
}
