package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"time"
)

//Define Item struct
type Item struct {
	Required    bool      `alias:"required"`
	Min         float64   `alias:"min"`
	Max         float64   `alias:"max"`
	MinLength   int       `alias:"minLength"`
	MaxLength   int       `alias:"maxLength"`
	Length      int       `alias:"length"`
	Fixed       string    `alias:"fixed"`
	Enums       []string  `alias:"enums"`
	Regex       string    `alias:"regex"`
	Before      time.Time `alias:"before"`
	After       time.Time `alias:"after"`
	Message     string    `alias:"message"`
	Code        int       `alias:"code"`
	defaultMsg  string
	defaultCode int
}

//setDefault
func (i *Item) setDefault() {
	if i.Message == "" {
		i.Message = i.defaultMsg
	}
	if i.Code == 0 {
		i.Code = i.defaultCode
	}
}

//Error
func (i *Item) Error(fieldName string) error {
	i.setDefault()
	err := struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	}{
		Msg:  fmt.Sprintf("[%s]%s", fieldName, i.Message),
		Code: i.Code,
	}
	bytes, _ := json.Marshal(&err)
	return errors.New(string(bytes))
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

//Validate
func (i *Item) Validate(modelPtr interface{}, field reflect.StructField, defaultMsg string, defaultCode int) error {
	i.defaultMsg = defaultMsg
	i.defaultCode = defaultCode
	if !i.Required {
		return nil
	}
	fieldv := reflect.ValueOf(modelPtr).Elem().FieldByName(field.Name)
	if !fieldv.IsValid() {
		panic("[Validate]field value is invalid")
	}
	passed := true
	switch fieldv.Type().Kind() {
	case reflect.String:
		passed = i.vstring(fieldv)
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		passed = i.vint(fieldv)
	case reflect.Float32, reflect.Float64:
		passed = i.vfloat(fieldv)
	case reflect.Slice:
		if fieldv.Type().Elem().Kind() == reflect.String {
			for pos := 0; pos < fieldv.Len(); pos++ {
				if passed = i.vstring(fieldv.Index(pos)); !passed {
					break
				}
			}
		}
	}
	if !passed {
		return i.Error(field.Name)
	}
	return nil
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
