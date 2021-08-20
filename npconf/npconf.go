package npconf

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/dfoplanner/dfoplanner/gen"
)

// Neople uses a ini accent for their configuration
// This package implements a simple parser for those format

type NpConf struct {
	Section []*Section
}

type Section struct {
	Name *string
	Sections []*Section
	Properties []*string
}


func LoadNpConf(data []byte) *NpConf {

	is := antlr.NewInputStream(string(data))

	lexer := parser.NewnpconfLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)

	p := parser.NewnpconfParser(stream)
	var listener NpConfParserListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Npconf())
	return listener.npConf
}