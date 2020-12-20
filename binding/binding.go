package binding

import (
	"encoding/json"
	"github.com/billcoding/flygo/log"
	"github.com/billcoding/flygo/reflectx"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

//Define Binding struct
type Binding struct {
	structPtr interface{}
	typ       *btype
	items     []*Item
	fields    []*reflect.StructField
	logger    log.Logger
	dataMap   map[string]interface{}
}

//New
func New(structPtr interface{}, typ *btype) *Binding {
	items := make([]*Item, 0)
	fields := reflectx.CreateFromTag(structPtr, &items, "alias", "binding")
	if len(items) != len(fields) {
		panic("[New]invalid len both items and fields")
	}
	return &Binding{
		structPtr: structPtr,
		typ:       typ,
		items:     items,
		fields:    fields,
		logger:    log.New("[Binding]"),
		dataMap:   make(map[string]interface{}, 0),
	}
}

//setMap
func setMap(m map[string]interface{}, k string, v []string) {
	if len(v) > 0 {
		if len(v) == 1 {
			m[k] = v[0]
		} else {
			m[k] = v
		}
	}
}

//Bind
func (b *Binding) Bind(req *http.Request) {
	b.initMap(req)
	b.setVal()
}

//initMap
func (b *Binding) initMap(req *http.Request) {
	switch b.typ {
	case Header:
		for k, v := range req.Header {
			setMap(b.dataMap, k, v)
		}
	case Param:
		for k, v := range req.Form {
			setMap(b.dataMap, k, v)
		}
		switch req.Header.Get("Content-Type") {
		case "multipart/form-data":
			err := req.ParseMultipartForm(0)
			if err == nil {
				for k, v := range req.MultipartForm.Value {
					setMap(b.dataMap, k, v)
				}
			}
		default:
			err := req.ParseForm()
			if err == nil {
				for k, v := range req.PostForm {
					setMap(b.dataMap, k, v)
				}
			}
		}
	case Body:
		cts := strings.Split(req.Header.Get("Content-Type"), ";")
		if len(cts) > 0 {
			contentType := cts[0]
			if contentType == "application/json" {
				bytes, err := ioutil.ReadAll(req.Body)
				if err == nil {
					if json.Valid(bytes) {
						err := json.Unmarshal(bytes, &b.dataMap)
						if err != nil {
							b.logger.Warn("[Bind]%v", err)
						}
					}
				}
			}
		}
	}
}

//setVal
func (b *Binding) setVal() {
	for pos, item := range b.items {
		field := b.fields[pos]
		value := reflect.ValueOf(b.structPtr).Elem().FieldByName(field.Name)
		item.Bind(field, value, b.dataMap)
	}
}
