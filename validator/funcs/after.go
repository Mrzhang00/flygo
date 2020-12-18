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
		t, err := time.Parse(time.RFC3339, value.String())
		if err != nil {
			panic(fmt.Sprintf("[afterFunc][Pass]%v", err))
		}
		return f.after.After(t)
	} else if value.String() == reflect.TypeOf(time.Time{}).String() {
		return f.after.After(value.Interface().(time.Time))
	}
	return true
}
