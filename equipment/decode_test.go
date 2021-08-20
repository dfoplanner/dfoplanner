package equipment

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/dfoplanner/dfoplanner/npconf"
	"strings"
	"testing"
)

func TestEquipment_Unmarshal(t *testing.T) {
	//data, _ := ioutil.ReadFile("../tests/equipment/100200085.equ")
	data :=
	"[name]\n\r" +
		"`<13::chn_flavor_text_490010012>`\n\r"+
		"[aa]\n\r"+
		"`d`\n\r"
	data = strings.Replace(data, "`","\"", -1)
	fmt.Printf("%s\n", data)
	npLexer := lexer.Must(lexer.NewSimple([]lexer.Rule{
		{"DateTime", `\d\d\d\d-\d\d-\d\dT\d\d:\d\d:\d\d(\.\d+)?(-\d\d:\d\d)?`, nil},
		{"Date", `\d\d\d\d-\d\d-\d\d`, nil},
		{"Time", `\d\d:\d\d:\d\d(\.\d+)?`, nil},
		{"Ident", `[a-zA-Z_][a-zA-Z_0-9]*`, nil},
		{"String", `"[^"]*"`, nil},
		{"StringPointer", `"<[^<]*>"`, nil},
		{"Number", `[-+]?[.0-9]+\b`, nil},
		{"comment", `#[^\n]+`, nil},
		{"whitespace", `\s+`, nil},
	}))
	var parser = participle.MustBuild(&npconf.NpConf{},
		participle.Lexer(npLexer),
	)
	test := &npconf.NpConf{}
	_ = parser.ParseString("",data, test)
	for k, v := range(test.Properties) {
		fmt.Printf("%+v %+v %+v\n", k, v.Key, v.Value)
	}

}
