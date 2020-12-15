package validator

import (
	"encoding/json"
)

//Define Item struct
type Item struct {
	Required  bool     `alias:"required"`
	Min       int      `alias:"min"`
	Max       int      `alias:"max"`
	MinLength int      `alias:"minLength"`
	MaxLength int      `alias:"maxLength"`
	Length    int      `alias:"length"`
	Fixed     string   `alias:"fixed"`
	Enums     []string `alias:"enums"`
	Regex     string   `alias:"regex"`
	Message   string   `alias:"message"`
	Code      int      `alias:"code"`
}

//Validate
func (i *Item) Validate(modelPtr interface{}, fieldName string) error {
	if !i.Required {
		return nil
	}

	//fieldt, _ := reflect.TypeOf(modelPtr).Elem().FieldByName(fieldName)
	//fieldv := reflect.ValueOf(modelPtr).Elem().FieldByName(fieldName)

	return nil
}

//String
func (i *Item) String() string {
	bytes, _ := json.Marshal(i)
	return string(bytes)
}
