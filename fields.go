package meerkats

import (
	"fmt"
	"reflect"
)


type Fields map[string]interface{}

func ( f *Fields) Merge(src...interface{}) {
	for _, field := range src {

		rElem := reflect.ValueOf(field).Elem()
		if (rElem.Kind() == reflect.Ptr) {
			rElem = rElem.Elem()
		}
		switch rElem.Kind() {
		case reflect.Struct:
			for i := 0; i < rElem.NumField(); i++ {
				valueField := rElem.Field(i)
				typeField := rElem.Type().Field(i)
				tag := typeField.Tag

				if tag := tag.Get(FIELD_STRUCT_TAG); tag == "-" {
					continue
				}

				(*f)[typeField.Name] = valueField.Interface()
			}
			break
		case reflect.Map:
			for k, v := range field.(map[string]interface{}) {
				(*f)[k] = v
			}
			break
		case reflect.Array:
			for i, v := range field.([]interface{}) {
				(*f)[string(i)] = v
			}
			break
		}
	}
}
func ( f *Fields) MarshalText() ([]byte, error) {
	var text []byte
	for k, v := range *f {
		text = append(text, []byte(fmt.Sprint(k, "=\"", v, "\" "))...)
	}

	return text, nil
}
func ( f *Fields) String() string {
	text, _ := f.MarshalText()
	return string(text)
}