package cmd

import "testing"

func TestValidateRunValid(t *testing.T) {
	resumeFile := "../example.json"
	err := Execute([]string{"validate", "-r", resumeFile}) //validateRun(validateCmd, []string{"-r", resumeFile})
	if err != nil {
		t.Fatalf("Expected valid state, got error, %s", err)
	}
}
