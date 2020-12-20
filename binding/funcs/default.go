package funcs

import (
	"reflect"
)

//Define defaultFunc
type defaultFunc struct {
	defaultVal string
}

//DefaultFunc
func DefaultFunc(defaultVal string) BFunc {
	return &defaultFunc{defaultVal}
}

//Bind
func (d *defaultFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	if !inValue.IsValid() {
		return reflect.ValueOf(d.defaultVal)
	}
	return inValue
}
