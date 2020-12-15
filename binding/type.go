package binding

var (
	Header = &Type{"header"}
	Param  = &Type{"param"}
	Body   = &Type{"body"}
)

//Define Type struct
type Type struct {
	t string
}
