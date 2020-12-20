package funcs

import (
	"fmt"
	"reflect"
	"strings"
)

//Define joinFunc struct
type joinFunc struct {
	join   bool
	joinsp string
}

//SplitFunc
func JoinFunc(join bool, joinsp string) BFunc {
	if joinsp == "" {
		joinsp = ","
	}
	return &joinFunc{
		join:   join,
		joinsp: joinsp,
	}
}

//Bind
func (j *joinFunc) Bind(inValue reflect.Value) (outValue reflect.Value) {
	if j.join && inValue.IsValid() {
		if inValue.Type().Kind() == reflect.Slice || inValue.Type().Kind() == reflect.Array {
			joins := make([]string, 0)
			for i := 0; i < inValue.Len(); i++ {
				joins = append(joins, fmt.Sprintf("%v", inValue.Index(i)))
			}
			return reflect.ValueOf(strings.Join(joins, j.joinsp))
		}
	}
	return inValue
}
