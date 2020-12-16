package validator

import (
	"github.com/billcoding/flygo/reflectx"
	"reflect"
)

//Define Validator struct
type Validator struct {
	structPtr   interface{}
	items       []*Item
	fields      []reflect.StructField
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
func (v *Validator) Validate() error {
	for pos, item := range v.items {
		result := item.Validate(v.structPtr, v.fields[pos], v.defaultMsg, v.defaultCode)
		if result != nil {
			return result
		}
	}
	return nil
}
