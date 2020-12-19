package validator

import (
	"fmt"
	"testing"
	"time"
)

func TestValidator(t *testing.T) {
	type userModel struct {
		//Int1 int8      `validate:"required(T) min(85)" message(int1)`
		//Int2 int16     `validate:"required(T) min(85)"`
		//Int3 int32     `validate:"required(T) min(85)"`
		//Int4 int       `validate:"required(T) min(85)"`
		//Int5 int64     `validate:"required(T) min(85) message(int5不合法)"`
		Name []string  `validate:"required(T) enums(aaa,bbb,ccc) message(名称不合法)"`
		Age  float64   `validate:"required(T) min(10) message(年龄不合法)"`
		Time time.Time `validate:"required(T) after(2011-11-07T00:00:17) before(2020-11-07T00:00:17) message(日期不能为空)"`
	}
	vs := New(&userModel{
		//Int1: 0,
		//Int2: 0,
		//Int3: 0,
		//Int3: 0,
		//Int4: 20,
		//Int5: 5850,
		Age:  1,
		Time: time.Now(),
	})
	result := vs.Validate()
	fmt.Println(fmt.Sprintf("passed : %v", result.Passed))
	for _, item := range result.Items {
		fmt.Println(fmt.Sprintf("%s passed : %v", item.Field.Name, item.Passed))
	}
	fmt.Println(fmt.Sprintf("union message : %v", result.Messages()))
}
