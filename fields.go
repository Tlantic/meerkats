package meerkats

import "fmt"

type Fields map[string]interface{}


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