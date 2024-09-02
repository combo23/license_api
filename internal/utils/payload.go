package utils

import "reflect"

// IsValidPayload checks if the payload is valid eg. all fields are not empty
func IsValidPayload(payload interface{}) bool {
	v := reflect.ValueOf(payload)
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // Dereference the pointer if it is a pointer to a struct
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String && field.Len() == 0 {
			return false
		}
	}

	return true
}
