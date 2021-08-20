package m_priest

import (
	"encoding/json"
	"github.com/dfoplanner/dfoplanner/status"
)

type Crusader struct {

}

func (c *Crusader) GetStatus(inDungeon bool) interface{} {
	s := status.Status{HP: "1000/1000"}
	b, _ := json.Marshal(&s)
	return string(b)
}
