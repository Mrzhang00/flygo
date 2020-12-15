package validator

import (
	"fmt"
	"testing"
)

func TestValidator(t *testing.T) {
	type userModel struct {
		Id   int    `validate:"required(TRUE) min(100)     max(10000) regex(^\\d{2}$)"`
		Name string `validate:"required(TRUE) fixed(1) "`
	}
	vs := Validators(&userModel{})
	fmt.Println(vs)
}
