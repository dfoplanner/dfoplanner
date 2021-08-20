package equipment

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	gen "github.com/dfoplanner/dfoplanner/gen"
	"github.com/dfoplanner/dfoplanner/npconf"
	"io/ioutil"
	"testing"
)

func TestEquipment_Unmarshal(t *testing.T) {
	data, _ := ioutil.ReadFile("../tests/equipment/100200085.equ")
	is := antlr.NewInputStream(string(data))

	lexer := gen.NewnpconfLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)

	p := gen.NewnpconfParser(stream)
	p.BuildParseTrees = true
	tree := p.Npconf()

	antlr.ParseTreeWalkerDefault.Walk(&npconf.TreeShapeListener{},tree)
	fmt.Printf("%+v\n\n", tree)
}
