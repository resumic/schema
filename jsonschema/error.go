package jsonschema

import "reflect"

type InvalidTagError struct {
	Tag  string
	Kind string
}

func (e *InvalidTagError) Error() string {
	return "jsonschema: " + e.Tag + " is an invalid tag for type " + e.Kind
}

type UnsupportedKindError struct {
	Kind reflect.Kind
}

func (e *UnsupportedKindError) Error() string {
	return "jsonschema: " + e.Kind.String() + " is not supported"
}
