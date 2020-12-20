package binding

var (
	Header = &btype{"header"}
	Param  = &btype{"param"}
	Body   = &btype{"body"}
)

//Define btype struct
type btype struct {
	t string
}
