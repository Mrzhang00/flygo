package reflectx

import (
	"fmt"
	"reflect"
	"strconv"
)

//SetFieldValue
func SetFieldValue(val reflect.Value, fieldValue reflect.Value) {
	switch fieldValue.Kind() {
	case reflect.Bool:
		vboolval, err := strconv.ParseBool(val.String())
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(vboolval))
	case reflect.String:
		fieldValue.Set(val)
	case reflect.Int8:
		vint8val, err := strconv.ParseInt(val.String(), 10, 8)
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(int8(vint8val)))
	case reflect.Int16:
		vint16val, err := strconv.ParseInt(val.String(), 10, 16)
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(int16(vint16val)))
	case reflect.Int32:
		vint32val, err := strconv.ParseInt(val.String(), 10, 32)
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(int32(vint32val)))
	case reflect.Int:
		vint32val, err := strconv.ParseInt(val.String(), 10, 32)
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(int(vint32val)))
	case reflect.Int64:
		vint64val, err := strconv.ParseInt(val.String(), 10, 64)
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(vint64val))
	case reflect.Float32:
		vfloatval, err := strconv.ParseFloat(val.String(), 32)
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(float32(vfloatval)))
	case reflect.Float64:
		vfloatval, err := strconv.ParseFloat(val.String(), 64)
		if err != nil {
			panic(fmt.Sprintf("[Validator]%v", err))
		}
		fieldValue.Set(reflect.ValueOf(vfloatval))
	default:
		if val.Kind() == fieldValue.Kind() {
			fieldValue.Set(val)
		}
	}
}
