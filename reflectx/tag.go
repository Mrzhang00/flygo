package reflectx

import (
	"reflect"
	"strings"
)

//CreateFromTag
func CreateFromTag(structPtr, distPtr interface{}, alias, tag string) []reflect.StructField {
	if reflect.TypeOf(distPtr).Kind() != reflect.Ptr {
		panic("[CreateFromTag]distPtr of non-ptr type")
	}
	if reflect.TypeOf(distPtr).Elem().Kind() != reflect.Slice {
		panic("[CreateFromTag]distPtr of non-slice type")
	}
	if reflect.TypeOf(distPtr).Elem().Elem().Kind() != reflect.Ptr {
		panic("[CreateFromTag]distPtr of non-ptr type")
	}
	if reflect.TypeOf(distPtr).Elem().Elem().Elem().Kind() != reflect.Struct {
		panic("[CreateFromTag]distPtr of non-struct type")
	}
	if reflect.TypeOf(structPtr).Kind() != reflect.Ptr {
		panic("[CreateFromTag]structPtr of non-pointer type")
	}
	distType := reflect.TypeOf(distPtr).Elem().Elem().Elem()
	structType := reflect.TypeOf(structPtr).Elem()
	aliasMap := make(map[string]string, 0)
	for i := 0; i < distType.NumField(); i++ {
		field := distType.Field(i)
		alias := strings.TrimSpace(field.Tag.Get(alias))
		if alias != "" {
			aliasMap[alias] = field.Name
		} else {
			aliasMap[field.Name] = field.Name
		}
	}
	distValues := reflect.ValueOf(distPtr).Elem()
	fields := make([]reflect.StructField, 0)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		validateTag, have := field.Tag.Lookup(tag)
		if !have || validateTag == "" {
			continue
		}
		tags := strings.Split(validateTag, " ")
		if len(tags) <= 0 {
			continue
		}
		distValue := reflect.New(distType)
		for _, tag := range tags {
			tag = strings.TrimSpace(tag)
			if tag == "" {
				continue
			}
			sindex := strings.Index(tag, "(")
			eindex := strings.LastIndex(tag, ")")
			if sindex < 2 {
				continue
			}
			if eindex == -1 || eindex <= sindex {
				continue
			}
			vname := tag[:sindex]
			vval := tag[sindex+1 : eindex]
			fieldName, have := aliasMap[vname]
			if have {
				fieldValue := distValue.Elem().FieldByName(fieldName)
				SetFieldValue(reflect.ValueOf(vval), fieldValue)
			}
		}
		fields = append(fields, field)
		distValues.Set(reflect.Append(distValues, distValue))
	}
	return fields
}
