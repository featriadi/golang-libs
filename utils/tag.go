package utils

import "reflect"

func GetJsonTagOnField(structTarget interface{}, fieldName string) string {
	structTargetType := reflect.TypeOf(structTarget)
	if structTargetType.Kind() != reflect.Ptr {
		structTargetType = structTargetType.Elem()
	}

	// Check if the field exists on the struct
	field, found := structTargetType.FieldByName(fieldName)
	if !found {
		return fieldName
	}

	// Get the JSON tag for the field
	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return fieldName
	}

	return jsonTag
}
