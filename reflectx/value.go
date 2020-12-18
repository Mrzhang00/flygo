package reflectx

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//SetFieldValue
func SetFieldValue(val string, fieldValue reflect.Value) {
	switch fieldValue.Kind() {
	case reflect.Bool:
		vboolval, err := strconv.ParseBool(val)
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(vboolval))
	case reflect.String:
		fieldValue.Set(reflect.ValueOf(val))
	case reflect.Int8:
		vint8val, err := strconv.ParseInt(val, 10, 8)
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(int8(vint8val)))
	case reflect.Int16:
		vint16val, err := strconv.ParseInt(val, 10, 16)
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(int16(vint16val)))
	case reflect.Int32:
		vint32val, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(int32(vint32val)))
	case reflect.Int:
		vint32val, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(int(vint32val)))
	case reflect.Int64:
		vint64val, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(vint64val))
	case reflect.Float32:
		vfloatval, err := strconv.ParseFloat(val, 32)
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(float32(vfloatval)))
	case reflect.Float64:
		vfloatval, err := strconv.ParseFloat(val, 64)
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(vfloatval))
	case reflect.Slice, reflect.Array:
		if fieldValue.Type().Elem().Kind() == reflect.String {
			fieldValue.Set(reflect.ValueOf(strings.Split(val, ",")))
		}
	case reflect.Struct:
		if fieldValue.Type() == reflect.TypeOf(time.Time{}) {
			t, err := time.Parse(time.RFC3339, val)
			if err != nil {
				panic(fmt.Sprintf("[reflectx][SetFieldValue]%v", err))
			}
			fieldValue.Set(reflect.ValueOf(t))
		}
	}
}
