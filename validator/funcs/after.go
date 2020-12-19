package funcs

import (
	"fmt"
	"reflect"
	"time"
)

//Define afterFunc struct
type afterFunc struct {
	after time.Time
}

//AfterFunc
func AfterFunc(after time.Time) VFunc {
	return &afterFunc{after}
}

//Accept
func (f *afterFunc) Accept(typ reflect.Type) bool {
	return typ.Kind() == reflect.String || typ == reflect.TypeOf(time.Time{})
}

//Pass
func (f *afterFunc) Pass(value reflect.Value) bool {
	if f.after.IsZero() {
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
		return f.after.Before(value.Interface().(time.Time))
	}
	return true
}
