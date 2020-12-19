package funcs

import (
	"reflect"
	"strings"
)

//Define splitFunc struct
type splitFunc struct {
	split   bool
	splitsp string
}

//SplitFunc
func SplitFunc(split bool, splitsp string) BFunc {
	return &splitFunc{
		split:   split,
		splitsp: splitsp,
	}
}

//Bind
func (s *splitFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	if s.split && inValue.IsValid() {
		if inValue.Type().Kind() == reflect.String {
			return reflect.ValueOf(strings.Split(inValue.String(), s.splitsp))
		}
	}
	return inValue
}
