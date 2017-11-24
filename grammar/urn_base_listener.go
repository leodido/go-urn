// Generated from /home/leodido/workspaces/go/src/github.com/leodido/go-urn/grammar/Urn.g4 by ANTLR 4.7.

package grammar // Urn
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseUrnListener is a complete listener for a parse tree produced by UrnParser.
type BaseUrnListener struct{}

var _ UrnListener = &BaseUrnListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseUrnListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseUrnListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseUrnListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseUrnListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterUrn is called when production urn is entered.
func (s *BaseUrnListener) EnterUrn(ctx *UrnContext) {}

// ExitUrn is called when production urn is exited.
func (s *BaseUrnListener) ExitUrn(ctx *UrnContext) {}

// EnterID is called when production iD is entered.
func (s *BaseUrnListener) EnterID(ctx *IDContext) {}

// ExitID is called when production iD is exited.
func (s *BaseUrnListener) ExitID(ctx *IDContext) {}

// EnterSS is called when production sS is entered.
func (s *BaseUrnListener) EnterSS(ctx *SSContext) {}

// ExitSS is called when production sS is exited.
func (s *BaseUrnListener) ExitSS(ctx *SSContext) {}
