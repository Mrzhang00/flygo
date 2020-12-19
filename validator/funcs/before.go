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
		t, err := time.Parse("2006-01-02T15:04:05", value.String())
		if err != nil {
			t, err = time.Parse("2006-01-02 15:04:05", value.String())
			if err != nil {
				panic(fmt.Sprintf("[afterFunc][Pass]%v", err))
			}
		}
		return f.Pass(reflect.ValueOf(t))
	} else if value.Type() == reflect.TypeOf(time.Time{}) {
		return f.before.After(value.Interface().(time.Time))
	}
	return true
}
