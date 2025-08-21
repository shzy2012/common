package tools

import (
	"fmt"
	"reflect"
	"strings"
)

func SetField(source interface{} /*must be a interface*/, fieldName string, fieldValue string) {
	v := reflect.ValueOf(source).Elem()
	field := v.FieldByName(fieldName)

	if field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
		field.SetString(fieldValue)
	}
}

func SetField2(source interface{}, fieldPath string, fieldValue string) error {
	v := reflect.ValueOf(source).Elem()

	// Split the field path to handle nested fields
	fields := strings.Split(fieldPath, ".")

	for i, fieldName := range fields {
		if v.Kind() == reflect.Struct {
			v = v.FieldByName(fieldName)
			if !v.IsValid() {
				return fmt.Errorf("field %s not found", fieldName)
			}
			if i < len(fields)-1 {
				// If it's not the last field, ensure it's a struct
				if v.Kind() == reflect.Ptr {
					v = v.Elem()
				}
				if v.Kind() != reflect.Struct {
					return fmt.Errorf("field %s is not a struct", fieldName)
				}
			}
		} else {
			return fmt.Errorf("field %s is not a struct", fieldName)
		}
	}

	if v.CanSet() && v.Kind() == reflect.String {
		v.SetString(fieldValue)
		return nil
	}

	return fmt.Errorf("field %s cannot be set", fieldPath)
}
