package binding

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestBinding(t *testing.T) {
	type userModel struct {
		Var string `binding:"name(id) default(xx) split(F) join(T) joinsp(',') prefix(') suffix(')"`
	}
	um := &userModel{}

	binding := New(um, Param)
	var buffer bytes.Buffer
	buffer.WriteString(`{"Id":100,"Name":"zhangsan"}`)
	req, _ := http.NewRequest(http.MethodGet, "http://localhost", &buffer)
	req.Header = make(http.Header, 0)
	req.Form = make(url.Values, 0)
	req.PostForm = make(url.Values, 0)

	req.Form.Add("id", "aaa")
	req.Form.Add("id", "bbb")
	req.Form.Add("id", "ccc")

	//req.PostForm.Set("Idx", "1000")
	//req.PostForm.Set("Namex", "lisu")

	//req.Header.Set("Id", "10000")
	//req.Header.Set("Name", "wangwu")
	//req.Header.Set("Content-Type", "application/json;chaset=utf-8")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	binding.Bind(req)

	fmt.Println(um)
}
