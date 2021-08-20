// Code generated from /Users/umi/Desktop/dfoplanner/npconf/npconf.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // npconf

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasenpconfListener is a complete listener for a parse tree produced by npconfParser.
type BasenpconfListener struct{}

var _ npconfListener = &BasenpconfListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasenpconfListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasenpconfListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasenpconfListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasenpconfListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNpconf is called when production npconf is entered.
func (s *BasenpconfListener) EnterNpconf(ctx *NpconfContext) {}

// ExitNpconf is called when production npconf is exited.
func (s *BasenpconfListener) ExitNpconf(ctx *NpconfContext) {}

// EnterSection is called when production section is entered.
func (s *BasenpconfListener) EnterSection(ctx *SectionContext) {}

// ExitSection is called when production section is exited.
func (s *BasenpconfListener) ExitSection(ctx *SectionContext) {}

// EnterSectionheader is called when production sectionheader is entered.
func (s *BasenpconfListener) EnterSectionheader(ctx *SectionheaderContext) {}

// ExitSectionheader is called when production sectionheader is exited.
func (s *BasenpconfListener) ExitSectionheader(ctx *SectionheaderContext) {}

// EnterSectionfooter is called when production sectionfooter is entered.
func (s *BasenpconfListener) EnterSectionfooter(ctx *SectionfooterContext) {}

// ExitSectionfooter is called when production sectionfooter is exited.
func (s *BasenpconfListener) ExitSectionfooter(ctx *SectionfooterContext) {}

// EnterStr is called when production str is entered.
func (s *BasenpconfListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production str is exited.
func (s *BasenpconfListener) ExitStr(ctx *StrContext) {}

// EnterLine is called when production line is entered.
func (s *BasenpconfListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BasenpconfListener) ExitLine(ctx *LineContext) {}

// EnterStringlist is called when production stringlist is entered.
func (s *BasenpconfListener) EnterStringlist(ctx *StringlistContext) {}

// ExitStringlist is called when production stringlist is exited.
func (s *BasenpconfListener) ExitStringlist(ctx *StringlistContext) {}
