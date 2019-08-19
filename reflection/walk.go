package walk

import (
	"reflect"
	"strconv"
)

func walk(x interface{}, fn func(intput string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Int:
			str := strconv.Itoa(int(field.Int()))
			fn(str)
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
