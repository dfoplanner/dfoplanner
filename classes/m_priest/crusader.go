package m_priest

import (
	"encoding/json"
	"github.com/dfoplanner/dfoplanner/status"
)

type Crusader struct {

}

func (c *Crusader) GetStatus(inDungeon bool) interface{} {
	var s status.Status
	if !inDungeon {
		s = status.Status{HP: "1000/1000"}
	} else {
		s = status.Status{HP: "8888/8888"}
	}
	b, _ := json.Marshal(&s)
	return string(b)
}
