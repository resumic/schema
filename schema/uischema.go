package schema

import (
	"encoding/json"
	"reflect"
)

func generateFieldUISchema(typ reflect.Type, tags schemaTags) (map[string]interface{}, error) {
	schema := map[string]interface{}{}
	if example, ok := tags["example"]; ok {
		schema["ui:placeholder"] = example
	}
	return schema, nil
}

func generateArrayUISchema(typ reflect.Type, tags schemaTags) (map[string]interface{}, error) {
	schema := map[string]interface{}{}
	items, err := generateUISchema(typ.Elem(), tags.itemsTags())
	if err != nil {
		return nil, err
	}
	schema["items"] = items
	return schema, nil
}

func generateObjectUISchema(typ reflect.Type, tags schemaTags) (map[string]interface{}, error) {
	schema := map[string]interface{}{}
	orders := []string{}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldName := getJSONName(field)
		if fieldName == "" {
			continue
		}
		orders = append(orders, fieldName)
		fieldTags := newSchemaTags(field.Tag)
		property, err := generateUISchema(field.Type, fieldTags)
		if err != nil {
			return nil, err
		}
		schema[fieldName] = property
	}
	schema["ui:order"] = orders
	return schema, nil
}

func generateUISchema(typ reflect.Type, tags schemaTags) (map[string]interface{}, error) {
	kind := typ.Kind()
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.String,
		reflect.Bool:
		return generateFieldUISchema(typ, tags)
	case reflect.Slice:
		return generateArrayUISchema(typ, tags)
	case reflect.Struct:
		return generateObjectUISchema(typ, tags)
	default:
		return nil, &UnsupportedKindError{kind: kind}
	}
}

func GenerateUISchema() ([]byte, error) {
	typ := reflect.TypeOf(Schema{})
	schema, err := generateUISchema(typ, schemaTags{})
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(schema, "", "  ")
}
