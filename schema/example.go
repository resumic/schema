package schema

import (
	"reflect"
	"strconv"
)

func generateIntExample(typ reflect.Type, tags schemaTags) (*reflect.Value, error) {
	example := reflect.New(typ).Elem()
	placeholder, ok := tags["placeholder"]
	if !ok {
		return &example, nil
	}

	bits := example.Type().Bits()
	value, err := strconv.ParseInt(placeholder, 0, bits)
	if err != nil {
		return nil, err
	}
	example.SetInt(value)
	return &example, nil
}

func generateUintExample(typ reflect.Type, tags schemaTags) (*reflect.Value, error) {
	example := reflect.New(typ).Elem()
	placeholder, ok := tags["placeholder"]
	if !ok {
		return &example, nil
	}

	bits := example.Type().Bits()
	value, err := strconv.ParseUint(placeholder, 0, bits)
	if err != nil {
		return nil, err
	}
	example.SetUint(value)
	return &example, nil
}

func generateFloatExample(typ reflect.Type, tags schemaTags) (*reflect.Value, error) {
	example := reflect.New(typ).Elem()
	placeholder, ok := tags["placeholder"]
	if !ok {
		return &example, nil
	}

	bits := example.Type().Bits()
	value, err := strconv.ParseFloat(placeholder, bits)
	if err != nil {
		return nil, err
	}
	example.SetFloat(value)
	return &example, nil
}

func generateStirngExample(typ reflect.Type, tags schemaTags) (*reflect.Value, error) {
	example := reflect.New(typ).Elem()
	placeholder, ok := tags["placeholder"]
	if !ok {
		return &example, nil
	}

	example.SetString(placeholder)
	return &example, nil
}

func generateBoolExample(typ reflect.Type, tags schemaTags) (*reflect.Value, error) {
	example := reflect.New(typ).Elem()
	placeholder, ok := tags["placeholder"]
	if !ok {
		return &example, nil
	}

	value, err := strconv.ParseBool(placeholder)
	if err != nil {
		return nil, err
	}
	example.SetBool(value)
	return &example, nil
}

func generateSliceExample(typ reflect.Type, tags schemaTags) (*reflect.Value, error) {
	itemType := typ.Elem()
	itemTags := tags.itemsTags()
	exampleItem, err := generateExample(itemType, itemTags)
	if err != nil {
		return nil, err
	}

	example := reflect.MakeSlice(typ, 0, 1)
	example = reflect.Append(example, *exampleItem)
	return &example, nil
}

func generateStructExample(typ reflect.Type, tags schemaTags) (*reflect.Value, error) {
	example := reflect.New(typ).Elem()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldTags := newSchemaTags(field.Tag)
		fieldExample, err := generateExample(field.Type, fieldTags)
		if err != nil {
			return nil, err
		}
		example.Field(i).Set(*fieldExample)
	}
	return &example, nil
}

func generateExample(typ reflect.Type, tags schemaTags) (*reflect.Value, error) {
	kind := typ.Kind()
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return generateIntExample(typ, tags)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return generateUintExample(typ, tags)
	case reflect.Float32, reflect.Float64:
		return generateFloatExample(typ, tags)
	case reflect.String:
		return generateStirngExample(typ, tags)
	case reflect.Bool:
		return generateBoolExample(typ, tags)
	case reflect.Slice:
		return generateSliceExample(typ, tags)
	case reflect.Struct:
		return generateStructExample(typ, tags)
	default:
		return nil, &UnsupportedKindError{kind: kind}
	}
}

func GenerateExample() (Schema, error) {
	typ := reflect.TypeOf(Schema{})
	example, err := generateExample(typ, schemaTags{})
	if err != nil {
		return Schema{}, err
	}
	return example.Interface().(Schema), nil
}
