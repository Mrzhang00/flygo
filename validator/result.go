package validator

//Define Result struct
type Result struct {
	StructPtr interface{}
	Passed    bool
	FieldName string
	Message   string
	Code      int
}
