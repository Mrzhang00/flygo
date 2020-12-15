package validator

import (
	"github.com/billcoding/flygo/reflectx"
	"log"
)

//Define Validator struct
type Validator struct {
	structPtr interface{}
	items     []*Item
	fieldPos  map[string]int
}

var (
	defaultCode    = 500
	defaultMessage = "parameter is invalid"
)

//New
func New(structPtr interface{}) *Validator {
	items := make([]*Item, 0)
	reflectx.CreateFromTag(structPtr, &items, "alias", "validate")
	fieldPos := reflectx.GetTagFieldPos(structPtr, "validate")
	return &Validator{
		structPtr: structPtr,
		items:     items,
		fieldPos:  fieldPos,
	}
}

//Pos
func (v *Validator) Pos(pos int) string {
	if v.fieldPos != nil {
		for k, v := range v.fieldPos {
			if v == pos {
				return k
			}
		}
	}
	return ""
}

//Validate
func (v *Validator) Validate() {
	for pos, item := range v.items {
		vresult := item.Validate(v.structPtr, v.Pos(pos))
		if vresult != nil {
			log.Println(vresult)
			break
		}
	}
}
