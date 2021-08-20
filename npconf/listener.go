package npconf

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/dfoplanner/dfoplanner/gen"
)

type NpConfParserListener struct {
	parser.BasenpconfListener
	npConf *NpConf
	stack []*Section
	currentSection *Section
}

func (s *NpConfParserListener) push(i *Section) {
	if s.stack == nil {
		s.stack = make([]*Section,0)
	}
	s.stack = append(s.stack, i)
}

func (s *NpConfParserListener) pop() *Section {
	if len(s.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := s.stack[len(s.stack)-1]

	// Remove the last element from the stack.
	s.stack = s.stack[:len(s.stack)-1]

	return result
}


func (s *NpConfParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *NpConfParserListener) VisitErrorNode(node antlr.ErrorNode) {
	panic(node)
}

// EnterEveryRule is called when any rule is entered.
func (s *NpConfParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *NpConfParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNpconf is called when production npconf is entered.
func (s *NpConfParserListener) EnterNpconf(ctx *parser.NpconfContext) {
	s.npConf = &NpConf{
		Section: make([]*Section, 0),
	}
}

// ExitNpconf is called when production npconf is exited.
func (s *NpConfParserListener) ExitNpconf(ctx *parser.NpconfContext) {
	for len(s.stack) != 0 {
		stackEle := s.pop()
		s.npConf.Section = append(s.npConf.Section, stackEle)
	}
}
// EnterSection is called when production section is entered.
func (s *NpConfParserListener) EnterSection(ctx *parser.SectionContext) {
}

// ExitSection is called when production section is exited.
func (s *NpConfParserListener) ExitSection(ctx *parser.SectionContext) {}

// EnterSectionheader is called when production sectionheader is entered.
func (s *NpConfParserListener) EnterSectionheader(ctx *parser.SectionheaderContext) {
	if s.currentSection != nil {
		s.push(s.currentSection)
	}
	s.currentSection = &Section{}
	currentSectionName := ctx.GetText()
	currentSectionNamePtr := &currentSectionName
	s.currentSection.Name = currentSectionNamePtr
	s.currentSection.Properties = make([]*string, 0)
}

// ExitSectionheader is called when production sectionheader is exited.
func (s *NpConfParserListener) ExitSectionheader(ctx *parser.SectionheaderContext) {}

// EnterSectionfooter is called when production sectionfooter is entered.
func (s *NpConfParserListener) EnterSectionfooter(ctx *parser.SectionfooterContext) {}

// ExitSectionfooter is called when production sectionfooter is exited.
func (s *NpConfParserListener) ExitSectionfooter(ctx *parser.SectionfooterContext) {
	parentSection := s.pop()
	if parentSection.Sections == nil {
		parentSection.Sections = make([]*Section, 0)
	}
	parentSection.Sections = append(parentSection.Sections,s.currentSection)
	s.currentSection = parentSection
}

// EnterStr is called when production str is entered.
func (s *NpConfParserListener) EnterStr(ctx *parser.StrContext) {}

// ExitStr is called when production str is exited.
func (s *NpConfParserListener) ExitStr(ctx *parser.StrContext) {}

// EnterLine is called when production line is entered.
func (s *NpConfParserListener) EnterLine(ctx *parser.LineContext) {}

// ExitLine is called when production line is exited.
func (s *NpConfParserListener) ExitLine(ctx *parser.LineContext) {
	line := ctx.GetText()
	linePtr := &line
	v := linePtr
	s.currentSection.Properties = append(s.currentSection.Properties, v)
}

// EnterStringlist is called when production stringlist is entered.
func (s *NpConfParserListener) EnterStringlist(ctx *parser.StringlistContext) {}

// ExitStringlist is called when production stringlist is exited.
func (s *NpConfParserListener) ExitStringlist(ctx *parser.StringlistContext) {}
