package flygo

import (
	"testing"
)

const (
	PORT1 = 9090 + iota
	PORT2
	PORT3
)

func TestGroups(t *testing.T) {
	NewAppWithId("zhangsan").Config.Flygo.Server.Port = PORT1
	NewAppWithId("lisi").Config.Flygo.Server.Port = PORT2
	NewAppWithId("wangwu").Config.Flygo.Server.Port = PORT3
	AppGroup.IgnoreMaster()
	AppGroup.Start()
}
