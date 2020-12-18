package funcs

import (
	"reflect"
)

//Define minFunc struct
type minFunc struct {
	min float64
}

//MinFunc
func MinFunc(min float64) VFunc {
	return &minFunc{min}
}

//Accept
func (f *minFunc) Accept(typ reflect.Type) bool {
	var acceptKinds = map[reflect.Kind]byte{
		// int types
		reflect.Int8:  0,
		reflect.Int16: 0,
		reflect.Int32: 0,
		reflect.Int:   0,
		reflect.Int64: 0,

		//float types
		reflect.Float32: 0,
		reflect.Float64: 0,
	}
	_, have := acceptKinds[typ.Kind()]
	return have
}

//Pass
func (f *minFunc) Pass(value reflect.Value) bool {
	if f.min <= 0 {
		return true
	}
	switch value.Type().Kind() {
	default:
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		return f.min <= float64(value.Int())
	case reflect.Float32, reflect.Float64:
		return f.min <= value.Float()
	}
	return true
}
