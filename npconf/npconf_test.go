package npconf

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestNpConf_Load(t *testing.T) {
	data, _ := ioutil.ReadFile("../tests/equipment/100200085.equ")
	n := LoadNpConf(data)
	fmt.Printf("%+v\n\n", n)
}
