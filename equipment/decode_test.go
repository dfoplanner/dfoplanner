package equipment

import (
	"io/ioutil"
	"testing"
)

func TestEquipment_Unmarshal(t *testing.T) {
	data, _ := ioutil.ReadFile("../tests/equipment/100200085.equ")
	Unmarshal(data, nil)
}
