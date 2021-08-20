package web

import (
	"github.com/dfoplanner/dfoplanner/classes/m_priest"
    "syscall/js"
)
func getStatus(this js.Value, args []js.Value) interface{} {
	inDungeon := args[0].Bool()
	c := &m_priest.Crusader{}
	return js.ValueOf(c.GetStatus(inDungeon))
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("GetStatus", js.FuncOf(getStatus))
	<-done
}