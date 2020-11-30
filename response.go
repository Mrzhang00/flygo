package flygo

//Define Response struct
type Response struct {
	data        []byte //Response data
	contentType string //content type
	done        bool   //handled
}

//Set content type
func (r *Response) SetContentType(contentType string) *Response {
	r.contentType = contentType
	return r
}

//Get content type
func (r *Response) GetContentType() string {
	return r.contentType
}

//Get data
func (r *Response) GetData() []byte {
	return r.data
}

//Set done
func (c *Context) SetDone(done bool) *Context {
	c.Response.SetDone(done)
	return c
}

//Set done
func (r *Response) SetDone(done bool) *Response {
	r.done = done
	return r
}
