package funcs

import (
	"fmt"
	"reflect"
	"time"
)

//Define beforeFunc struct
type beforeFunc struct {
	before time.Time
}

//BeforeFunc
func BeforeFunc(before time.Time) VFunc {
	return &beforeFunc{before}
}

//Accept
func (f *beforeFunc) Accept(typ reflect.Type) bool {
	return typ.Kind() == reflect.String || typ == reflect.TypeOf(time.Time{})
}

//Pass
func (f *beforeFunc) Pass(value reflect.Value) bool {
	if f.before.IsZero() {
		return true
	}
	if value.Type().Kind() == reflect.String {
		t, err := time.Parse(time.RFC3339, value.String())
		if err != nil {
			panic(fmt.Sprintf("[beforeFunc][Pass]%v", err))
		}
		return f.before.Before(t)
	} else if value.String() == reflect.TypeOf(time.Time{}).String() {
		return f.before.Before(value.Interface().(time.Time))
	}
	return true
}
