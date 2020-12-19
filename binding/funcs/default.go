package funcs

import (
	"github.com/billcoding/flygo/log"
	"reflect"
)

//Define defaultFunc
type defaultFunc struct {
	defaultVal string
	logger     log.Logger
}

//DefaultFunc
func DefaultFunc(defaultVal string) BFunc {
	return &defaultFunc{
		defaultVal: defaultVal,
		logger:     log.New("[defaultFunc]"),
	}
}

//Bind
func (d *defaultFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	if !inValue.IsValid() {
		return reflect.ValueOf(d.defaultVal)
	}
	return inValue
}
