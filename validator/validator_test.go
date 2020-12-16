package validator

import (
	"fmt"
	"testing"
)

func TestValidator(t *testing.T) {
	type userModel struct {
		Id   float32  `validate:"required(T)"`
		Name []string `validate:"required(T) minLength(10) maxLength(10) fixed(1234567890x) enumsx(aaa,bbb,ccc) regexx(^\\d{2}$)"`
		Age  int      `validate:"required(T) min(10) max(200) message(年龄不合法) code(10010)"`
	}
	vs := New(&userModel{
		Id:   1.89,
		Name: []string{"1234567890x"},
		Age:  8,
	}, "参数不合法", 5000)
	fmt.Println(vs.Validate())
}
