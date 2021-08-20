// Code generated from /Users/umi/Desktop/dfoplanner/npconf/npconf.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // npconf

import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by npconfParser.
type npconfVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by npconfParser#npconf.
	VisitNpconf(ctx *NpconfContext) interface{}

	// Visit a parse tree produced by npconfParser#section.
	VisitSection(ctx *SectionContext) interface{}

	// Visit a parse tree produced by npconfParser#sectionheader.
	VisitSectionheader(ctx *SectionheaderContext) interface{}

	// Visit a parse tree produced by npconfParser#sectionfooter.
	VisitSectionfooter(ctx *SectionfooterContext) interface{}

	// Visit a parse tree produced by npconfParser#str.
	VisitStr(ctx *StrContext) interface{}

	// Visit a parse tree produced by npconfParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by npconfParser#stringlist.
	VisitStringlist(ctx *StringlistContext) interface{}

}