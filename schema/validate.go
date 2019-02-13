package schema

import (
	"github.com/xeipuuv/gojsonschema"
)

type InvalidResumeError struct {
	errors []gojsonschema.ResultError
}

func (e *InvalidResumeError) Error() string {
	msg := "Resume is invalid:"
	for _, err := range e.errors {
		msg += "\n\t" + err.String()
	}
	return msg
}

func ValidateResume(resume []byte) error {
	schema, err := GenerateJSONSchema()
	if err != nil {
		return err
	}

	schemaLoader := gojsonschema.NewBytesLoader(schema)
	resumeLoader := gojsonschema.NewBytesLoader(resume)

	result, err := gojsonschema.Validate(schemaLoader, resumeLoader)
	if err != nil {
		return err
	}
	if !result.Valid() {
		return &InvalidResumeError{errors: result.Errors()}
	}
	return nil
}
