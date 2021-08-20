// Code generated from /Users/umi/Desktop/dfoplanner/npconf/npconf.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // npconf

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasenpconfVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasenpconfVisitor) VisitNpconf(ctx *NpconfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasenpconfVisitor) VisitSection(ctx *SectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasenpconfVisitor) VisitSectionheader(ctx *SectionheaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasenpconfVisitor) VisitSectionfooter(ctx *SectionfooterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasenpconfVisitor) VisitStr(ctx *StrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasenpconfVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasenpconfVisitor) VisitStringlist(ctx *StringlistContext) interface{} {
	return v.VisitChildren(ctx)
}
