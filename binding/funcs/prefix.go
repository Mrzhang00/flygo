package funcs

import (
	"reflect"
)

//Define prefixFunc struct
type prefixFunc struct {
	prefix string
}

//PrefixFunc
func PrefixFunc(prefix string) BFunc {
	return &prefixFunc{prefix}
}

//Bind
func (p *prefixFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	if p.prefix != "" && inValue.IsValid() {
		if inValue.Type().Kind() == reflect.String {
			return reflect.ValueOf(p.prefix + inValue.String())
		}
	}
	return inValue
}
