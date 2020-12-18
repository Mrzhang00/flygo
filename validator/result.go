package validator

import "reflect"

//Define Result struct
type Result struct {
	StructPtr interface{}
	Passed    bool
	Items     []*ResultItem
}

//Define ResultItem struct
type ResultItem struct {
	Field   *reflect.StructField
	Passed  bool
	Message string
}
