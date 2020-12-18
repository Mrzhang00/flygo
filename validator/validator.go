package validator

import (
	"github.com/billcoding/flygo/reflectx"
	"reflect"
)

//Define Validator struct
type Validator struct {
	structPtr   interface{}
	items       []*Item
	fields      []*reflect.StructField
	defaultMsg  string
	defaultCode int
}

//New
func New(structPtr interface{}, defaultMsg string, defaultCode int) *Validator {
	items := make([]*Item, 0)
	fields := reflectx.CreateFromTag(structPtr, &items, "alias", "validate")
	if len(items) != len(fields) {
		panic("[New]invalid pos both items and fields")
	}
	return &Validator{
		structPtr:   structPtr,
		items:       items,
		fields:      fields,
		defaultMsg:  defaultMsg,
		defaultCode: defaultCode,
	}
}

//Validate
func (v *Validator) Validate() *Result {
	ritems := make([]*ResultItem, len(v.items), len(v.items))
	passedCount := 0
	for pos, item := range v.items {
		field := v.fields[pos]
		value := reflect.ValueOf(v.structPtr).Elem().FieldByName(field.Name)
		passed, msg := item.Validate(field, value)
		if passed {
			passedCount++
		}
		ritems[pos] = &ResultItem{
			Field:   v.fields[pos],
			Passed:  passed,
			Message: msg,
		}
	}
	return &Result{
		StructPtr: v.structPtr,
		Passed:    len(v.items) == passedCount,
		Items:     ritems,
	}
}
