package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	// This is similar to the pointer scenario before
	// We are trying to call NumField on our reflect.Value
	// But it doesn't have one as it's not a struct.
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}

	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer || val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
