package uploadfile

import (
	"github.com/billcoding/flygo"
	"testing"
)

func TestUploadfile(t *testing.T) {
	uploadfile := New()
	flygo.GetApp().Use(uploadfile).UseNotFound().UseRecovery().Run()
}
