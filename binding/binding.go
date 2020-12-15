package binding

//Define Binding struct
type Binding struct {
	Name      string `alias:"name"`
	Default   string `alias:"default"`
	Split     bool   `alias:"Split"`
	SplitChar string `alias:"splitChar"`
	Join      bool   `alias:"Join"`
	JoinChar  string `alias:"joinChar"`
}

type userModel struct {
	Id int `binding:"name(id) default(100) split(true) splitSp(,) join(true) joinSp(,)"`
}
