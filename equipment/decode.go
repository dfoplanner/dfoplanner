package equipment
import (
	"fmt"
	"gopkg.in/ini.v1"
)
func Unmarshal(data []byte, v *Equipment) error {
	cfg, err := ini.Load(data)
	if err != nil {
		panic(err)
	}
	for k, v := range cfg.Sections() {
		fmt.Printf("%d, %+v\n", k, v)
	}
	return nil
}
