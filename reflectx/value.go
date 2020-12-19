package reflectx

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

var logger = log.New(os.Stdout, "[reflectx]", log.LstdFlags)

//SetValue
func SetValue(sourceValue reflect.Value, distValue reflect.Value) {
	switch distValue.Kind() {
	case reflect.Bool:
		accept(sourceValue.Type().Kind() == reflect.String, func() {
			vboolval, err := strconv.ParseBool(sourceValue.String())
			if err != nil {
				logger.Println(fmt.Sprintf("[SetValue]%v", err))
			}
			distValue.SetBool(vboolval)
		})
		accept(sourceValue.Type().Kind() == reflect.Bool, func() {
			distValue.Set(sourceValue)
		})
	case reflect.String:
		distValue.SetString(fmt.Sprintf("%v", sourceValue.Interface()))
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
		accept(sourceValue.Type().Kind() == reflect.String, func() {
			vintval, err := strconv.ParseInt(sourceValue.String(), 10, 8)
			if err != nil {
				logger.Println(fmt.Sprintf("[SetValue]%v", err))
			}
			distValue.SetInt(vintval)
		})
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
		accept(sourceValue.Type().Kind() == reflect.String, func() {
			vuintval, err := strconv.ParseUint(sourceValue.String(), 10, 8)
			if err != nil {
				logger.Println(fmt.Sprintf("[SetValue]%v", err))
			}
			distValue.SetUint(vuintval)
		})
	case reflect.Float32, reflect.Float64:
		accept(sourceValue.Type().Kind() == reflect.String, func() {
			vfloatval, err := strconv.ParseFloat(sourceValue.String(), 32)
			if err != nil {
				logger.Println(fmt.Sprintf("[SetValue]%v", err))
			}
			distValue.SetFloat(vfloatval)
		})
	case reflect.Slice, reflect.Array:
		accept(sourceValue.Type().Kind() == reflect.Array ||
			sourceValue.Type().Kind() == reflect.Slice, func() {
			distValue.Set(sourceValue)
		})
	case reflect.Struct:
		switch distValue.Type() {
		case reflect.TypeOf(time.Time{}):
			accept(sourceValue.Type().Kind() == reflect.String, func() {
				distValue.Set(reflect.ValueOf(ParseTime(sourceValue.String())))
			})
		case reflect.TypeOf(time.Second):
			accept(sourceValue.Type().Kind() == reflect.String, func() {
				distValue.Set(reflect.ValueOf(ParseDuration(sourceValue.String())))
			})
		}
	}
}

func accept(ok bool, callFunc func()) {
	if ok {
		callFunc()
	}
}
