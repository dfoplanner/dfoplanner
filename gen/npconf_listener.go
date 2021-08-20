// Code generated from /Users/umi/Desktop/dfoplanner/npconf/npconf.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // npconf

import "github.com/antlr/antlr4/runtime/Go/antlr"

// npconfListener is a complete listener for a parse tree produced by npconfParser.
type npconfListener interface {
	antlr.ParseTreeListener

	// EnterNpconf is called when entering the npconf production.
	EnterNpconf(c *NpconfContext)

	// EnterSection is called when entering the section production.
	EnterSection(c *SectionContext)

	// EnterSectionheader is called when entering the sectionheader production.
	EnterSectionheader(c *SectionheaderContext)

	// EnterSectionfooter is called when entering the sectionfooter production.
	EnterSectionfooter(c *SectionfooterContext)

	// EnterStr is called when entering the str production.
	EnterStr(c *StrContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterStringlist is called when entering the stringlist production.
	EnterStringlist(c *StringlistContext)

	// ExitNpconf is called when exiting the npconf production.
	ExitNpconf(c *NpconfContext)

	// ExitSection is called when exiting the section production.
	ExitSection(c *SectionContext)

	// ExitSectionheader is called when exiting the sectionheader production.
	ExitSectionheader(c *SectionheaderContext)

	// ExitSectionfooter is called when exiting the sectionfooter production.
	ExitSectionfooter(c *SectionfooterContext)

	// ExitStr is called when exiting the str production.
	ExitStr(c *StrContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitStringlist is called when exiting the stringlist production.
	ExitStringlist(c *StringlistContext)
}
