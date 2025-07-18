package reflects

import "reflect"

func GetStructFieldFromStructType(structTarget reflect.Type, fieldName string) *reflect.StructField {
	if structTarget.Kind() == reflect.Ptr {
		structTarget = structTarget.Elem()
	}

	if structTarget.Kind() != reflect.Struct {
		return nil
	}

	field, found := structTarget.FieldByName(fieldName)
	if !found {
		return nil
	}

	return &field
}
