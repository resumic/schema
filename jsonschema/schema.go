package jsonschema

import (
	"encoding/json"
	"reflect"
	"strings"
)

type numberSchema struct {
	description string
}

func newNumberSchema(tags []string) (*numberSchema, error) {
	s := numberSchema{}

	for _, t := range tags {
		if strings.HasPrefix(t, "description:") {
			s.description = strings.TrimPrefix(t, "description:")
		} else {
			return nil, &InvalidTagError{Tag: t, Kind: "number"}
		}
	}

	return &s, nil
}

func (s *numberSchema) marshal() map[string]interface{} {
	m := map[string]interface{}{"type": "number"}
	if s.description != "" {
		m["description"] = s.description
	}
	return m
}

type stringSchema struct {
	description string
	format      string
}

func newStringSchema(tags []string) (*stringSchema, error) {
	s := stringSchema{}

	for _, t := range tags {
		if strings.HasPrefix(t, "description:") {
			s.description = strings.TrimPrefix(t, "description:")
		} else if strings.HasPrefix(t, "format:") {
			s.format = strings.TrimPrefix(t, "format:")
		} else {
			return nil, &InvalidTagError{Tag: t, Kind: "string"}
		}
	}

	return &s, nil
}

func (s *stringSchema) marshal() map[string]interface{} {
	m := map[string]interface{}{"type": "string"}
	if s.description != "" {
		m["description"] = s.description
	}
	if s.format != "" {
		m["format"] = s.format
	}
	return m
}

type booleanSchema struct {
	description string
	format      string
}

func newBooleanSchema(tags []string) (*booleanSchema, error) {
	s := booleanSchema{}

	for _, t := range tags {
		if strings.HasPrefix(t, "description:") {
			s.description = strings.TrimPrefix(t, "description:")
		} else if strings.HasPrefix(t, "format:") {
			s.format = strings.TrimPrefix(t, "format:")
		} else {
			return nil, &InvalidTagError{Tag: t, Kind: "boolean"}
		}
	}

	return &s, nil
}

func (s *booleanSchema) marshal() map[string]interface{} {
	m := map[string]interface{}{"type": "boolean"}
	if s.description != "" {
		m["description"] = s.description
	}
	if s.format != "" {
		m["format"] = s.format
	}
	return m
}

type arraySchema struct {
	description     string
	additionalItems bool
	items           schema
}

func newArraySchema(tags []string, elemType reflect.Type) (*arraySchema, error) {
	s := arraySchema{}

	var elemTags []string
	for _, t := range tags {
		if strings.HasPrefix(t, "description:") {
			s.description = strings.TrimPrefix(t, "description:")
		} else if t == "additionalItems" {
			s.additionalItems = true
		} else if strings.HasPrefix(t, "items_") {
			elemTags = append(elemTags, strings.TrimPrefix(t, "items_"))
		} else {
			return nil, &InvalidTagError{Tag: t, Kind: "array"}
		}
	}

	items, err := newSchema(elemTags, elemType)
	if err != nil {
		return nil, err
	}
	s.items = items

	return &s, nil
}

func (s *arraySchema) marshal() map[string]interface{} {
	m := map[string]interface{}{"type": "array"}
	if s.description != "" {
		m["description"] = s.description
	}
	m["additionalItems"] = s.additionalItems
	m["items"] = s.items.marshal()
	return m
}

type objectSchema struct {
	description          string
	additionalProperties bool
	format               string
	properties           map[string]schema
}

func newObjectSchema(tags []string, properties []reflect.StructField) (*objectSchema, error) {
	s := objectSchema{}

	for _, t := range tags {
		if strings.HasPrefix(t, "description:") {
			s.description = strings.TrimPrefix(t, "description:")
		} else if strings.HasPrefix(t, "format:") {
			s.format = strings.TrimPrefix(t, "format:")
		} else if t == "additionalProperties" {
			s.additionalProperties = true
		} else {
			return nil, &InvalidTagError{Tag: t, Kind: "object"}
		}
	}

	s.properties = make(map[string]schema)
	for _, p := range properties {
		tt := strings.Split(string(p.Tag.Get("jsonschema")), ",")
		k := p.Name
		if tt[0] != "" {
			if tt[0] != "-" {
				k = tt[0]
			}
		}
		v, err := newSchema(tt[1:], p.Type)
		if err != nil {
			return nil, err
		}
		s.properties[k] = v
	}

	return &s, nil
}

func (s *objectSchema) marshal() map[string]interface{} {
	m := map[string]interface{}{"type": "object"}
	if s.description != "" {
		m["description"] = s.description
	}
	m["additionalProperties"] = s.additionalProperties
	if s.format != "" {
		m["format"] = s.format
	}
	pp := map[string]interface{}{}
	for k, v := range s.properties {
		pp[k] = v.marshal()
	}
	m["properties"] = pp
	return m
}

type schema interface {
	marshal() map[string]interface{}
}

func newSchema(tt []string, ty reflect.Type) (schema, error) {
	kind := ty.Kind()
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return newNumberSchema(tt)
	case reflect.String:
		return newStringSchema(tt)
	case reflect.Bool:
		return newBooleanSchema(tt)
	case reflect.Array, reflect.Slice:
		return newArraySchema(tt, ty.Elem())
	case reflect.Map, reflect.Struct:
		var pt []reflect.StructField
		for i := 0; i < ty.NumField(); i++ {
			pt = append(pt, ty.Field(i))
		}
		return newObjectSchema(tt, pt)
	default:
		return nil, &UnsupportedKindError{Kind: kind}
	}
}

type Schema struct {
	schema schema
	title  string
	draft  string
}

func NewSchema(i interface{}, tags, title string) (*Schema, error) {
	tt := strings.Split(tags, ",")
	if len(tt) == 1 {
		if tt[0] == "" {
			tt = []string{}
		}
	}
	s, err := newSchema(tt, reflect.TypeOf(i))
	if err != nil {
		return nil, err
	}
	return &Schema{
		schema: s,
		title:  title,
		draft:  "http://json-schema.org/draft-04/schema#",
	}, nil
}

func (s *Schema) MarshalJSON() ([]byte, error) {
	m := s.schema.marshal()
	m["title"] = s.title
	m["$schema"] = s.draft
	return json.Marshal(m)
}
