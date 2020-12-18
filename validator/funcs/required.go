package funcs

import "reflect"

//Define RequiredFunc struct
type requiredFunc struct {
}

//RequiredFunc
func RequiredFunc() VFunc {
	return &requiredFunc{}
}

//Accept
func (r *requiredFunc) Accept(typ reflect.Type) bool {
	return true
}

//Pass
func (r *requiredFunc) Pass(value reflect.Value) bool {
	return value.IsValid()
}
