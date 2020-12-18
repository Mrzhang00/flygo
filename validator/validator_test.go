package validator

import (
	"fmt"
	"testing"
	"time"
)

func TestValidator2(t *testing.T) {
	t1, _ := time.Parse(time.RFC3339, "2011-11-07T00:00:17")
	t2, _ := time.Parse(time.RFC3339, "2012-11-07T00:00:17")
	fmt.Println(t1.Before(t2))
}

func TestValidator(t *testing.T) {
	type userModel struct {
		//Int1 int8  `validate:"required(F) min(85)"`
		//Int2 int16 `validate:"required(F) min(85)"`
		//Int3 int32 `validate:"required(F) min(85)"`
		//Int4 int   `validate:"required(F) min(85)"`
		//Int5 int64 `validate:"required(F) min(85)"`
		//Name []string  `validate:"required(T) enumsx(aaa,bbb,ccc) regexx(^\\d{2}$)"`
		//Age  int       `validate:"required(T) min(1)"`
		Time time.Time `validate:"required(T) after(2011-11-07T00:00:17)"`
	}
	vs := New(&userModel{
		//Int1: 0,
		//Int2: 0,
		//Int3: 0,
		//Int4: 20,
		//Int5: 5850,
	}, "参数不合法", 500)
	result := vs.Validate()
	fmt.Println(fmt.Sprintf("passed : %v", result.Passed))
	for _, item := range result.Items {
		fmt.Println(fmt.Sprintf("%s passed : %v", item.Field.Name, item.Passed))
	}
}
