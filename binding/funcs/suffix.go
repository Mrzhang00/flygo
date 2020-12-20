package funcs

import (
	"reflect"
)

//Define suffixFunc struct
type suffixFunc struct {
	suffix string
}

//SuffixFunc
func SuffixFunc(suffix string) BFunc {
	return &suffixFunc{suffix}
}

//Bind
func (s *suffixFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	if s.suffix != "" && inValue.IsValid() {
		if inValue.Type().Kind() == reflect.String {
			return reflect.ValueOf(inValue.String() + s.suffix)
		}
	}
	return inValue
}
