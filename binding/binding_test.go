package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBinding(t *testing.T) {
	typeOf := reflect.TypeOf(&userModel{})
	field, _ := typeOf.Elem().FieldByName("Id")
	fmt.Println(field.Tag.Get("binding"))
	fmt.Println(field.Tag.Get("validate"))
}
