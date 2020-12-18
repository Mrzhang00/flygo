package validator

import (
	"encoding/json"
	"github.com/billcoding/flygo/validator/funcs"
	"reflect"
	"regexp"
	"time"
)

//Define Item struct
type Item struct {
	Required  bool      `alias:"required"`
	Min       float64   `alias:"min"`
	Max       float64   `alias:"max"`
	MinLength int       `alias:"minLength"`
	MaxLength int       `alias:"maxLength"`
	Length    int       `alias:"length"`
	Fixed     string    `alias:"fixed"`
	Enums     []string  `alias:"enums"`
	Regex     string    `alias:"regex"`
	Before    time.Time `alias:"before"`
	After     time.Time `alias:"after"`
	Message   string    `alias:"message"`
}

//contains
func contains(items []string, item string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

//match
func match(pattern, str string) bool {
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}

//vfuncs
func (i *Item) vfuncs() []funcs.VFunc {
	return []funcs.VFunc{
		funcs.RequiredFunc(),
		funcs.MinFunc(i.Min),
		funcs.MaxFunc(i.Max),
		funcs.FixedFunc(i.Fixed),
		funcs.EnumsFunc(i.Enums...),
		funcs.MinLengthFunc(i.MinLength),
		funcs.MaxLengthFunc(i.MaxLength),
		funcs.RegexFunc(i.Regex),
		funcs.BeforeFunc(i.Before),
		funcs.AfterFunc(i.After),
	}
}

//Validate
func (i *Item) Validate(field *reflect.StructField, value reflect.Value) (bool, string) {
	if !i.Required {
		return true, i.Message
	}
	passed := true
	vfuncs := i.vfuncs()
	for _, vFunc := range vfuncs {
		if !vFunc.Accept(field.Type) {
			continue
		}
		passed = vFunc.Pass(value)
		if !passed {
			break
		}
	}
	return passed, i.Message
}

//String
func (i *Item) String() string {
	bytes, _ := json.Marshal(i)
	return string(bytes)
}

//vstring
func (i *Item) vstring(fieldv reflect.Value) bool {
	passed := true
	strv := fieldv.String()
	if strv == "" {
		passed = false
	} else if i.Fixed != "" && strv != i.Fixed {
		passed = false
	} else if i.Length > 0 && len(strv) != i.Length {
		passed = false
	} else if i.MinLength > 0 && len(strv) < i.MinLength {
		passed = false
	} else if i.MaxLength > 0 && len(strv) > i.MaxLength {
		passed = false
	} else if i.Enums != nil && len(i.Enums) > 0 && !contains(i.Enums, strv) {
		passed = false
	} else if i.Regex != "" && !match(i.Regex, strv) {
		passed = false
	}
	return passed
}

//vint
func (i *Item) vint(fieldv reflect.Value) bool {
	passed := true
	stri := fieldv.Int()
	if i.Min > 0 && stri < int64(i.Min) {
		passed = false
	} else if i.Max > 0 && stri > int64(i.Max) {
		passed = false
	}
	return passed
}

//vfloat
func (i *Item) vfloat(fieldv reflect.Value) bool {
	passed := true
	strf := fieldv.Float()
	if i.Min > 0 && strf < i.Min {
		passed = false
	} else if i.Max > 0 && strf > i.Max {
		passed = false
	}
	return passed
}

//vtime
func (i *Item) vtime() bool {
	return false
}
