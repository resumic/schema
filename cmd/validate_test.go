package cmd

import "testing"

func TestValidateRunValid(t *testing.T) {
	resumeFile := "../example.json"
	err := validateRun(validateCmd, []string{resumeFile})
	if err != nil {
		t.Fatalf("Expected valid state, got error, %s", err)
	}
}
