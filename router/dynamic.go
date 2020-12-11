package router

//Define Dynamic struct
type Dynamic struct {
	Pos     map[int]string //dynamic params pos index
	*Simple                //contains Simple router
}
