package flygo

//Define Response struct
type Response struct {
	data        []byte //Response data
	contentType string //content type
	done        bool   //handled
}

//Set content type
func (response *Response) SetContentType(contentType string) *Response {
	response.contentType = contentType
	return response
}

//Get content type
func (response *Response) GetContentType() string {
	return response.contentType
}

//Get data
func (response *Response) GetData() []byte {
	return response.data
}

//Set done
func (c *Context) SetDone(done bool) *Context {
	c.Response.SetDone(done)
	return c
}

//Set done
func (response *Response) SetDone(done bool) *Response {
	response.done = done
	return response
}
