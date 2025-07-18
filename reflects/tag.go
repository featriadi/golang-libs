package reflects

import (
	"reflect"
	"strings"
)

func GetJSONTagFromStructField(field reflect.StructField) string {
	jsonTag := field.Tag.Get("json")
	jsonTagName := strings.Split(jsonTag, ",")[0]

	if jsonTagName == "" {
		return strings.ToLower(field.Name)
	}

	return jsonTagName
}

func GetJsonTagFromStructInterface(structInterface interface{}, fieldName string) string {
	structType := reflect.TypeOf(structInterface)

	field := GetStructFieldFromStructType(structType, fieldName)
	if field == nil {
		return ""
	}

	return GetJSONTagFromStructField(*field)
}

func GetJsonTagFromStruct[T any](fieldName string) string {
	var structType T
	structTarget := reflect.TypeOf(structType)

	field := GetStructFieldFromStructType(structTarget, fieldName)
	if field == nil {
		return ""
	}

	return GetJSONTagFromStructField(*field)
}

func MappingJsonTagFromStruct[T any]() []string {
	var structType T
	structTarget := reflect.TypeOf(structType)

	jsonTags := make([]string, structTarget.NumField())

	for i := 0; i < structTarget.NumField(); i++ {
		field := structTarget.Field(i)
		jsonTags[i] = GetJSONTagFromStructField(field)
	}

	return jsonTags
}
