package validator

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestValidator(t *testing.T) {
	type userModel struct {
		Id   float32   `validate:"required(T)"`
		Name []string  `validate:"required(T) enumsx(aaa,bbb,ccc) regexx(^\\d{2}$)"`
		Age  int       `validate:"required(T) min(1)"`
		Time time.Time `validate:"required(T) after(2010-11-07 00:00:17) before(2010-11-07 00:00:17)"`
	}
	vs := New(&userModel{
		Id:   1.89,
		Name: []string{"aaa"},
		Age:  8,
	}, "参数不合法", 5000)
	fmt.Println(vs.Validate())
}

func TestValidator2(t *testing.T) {
	fmt.Println(reflect.TypeOf(time.Now()).String())
}
