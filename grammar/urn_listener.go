// Generated from /home/leodido/workspaces/go/src/github.com/leodido/go-urn/grammar/Urn.g4 by ANTLR 4.7.

package grammar // Urn
import "github.com/antlr/antlr4/runtime/Go/antlr"

// UrnListener is a complete listener for a parse tree produced by UrnParser.
type UrnListener interface {
	antlr.ParseTreeListener

	// EnterUrn is called when entering the urn production.
	EnterUrn(c *UrnContext)

	// EnterID is called when entering the iD production.
	EnterID(c *IDContext)

	// EnterSS is called when entering the sS production.
	EnterSS(c *SSContext)

	// ExitUrn is called when exiting the urn production.
	ExitUrn(c *UrnContext)

	// ExitID is called when exiting the iD production.
	ExitID(c *IDContext)

	// ExitSS is called when exiting the sS production.
	ExitSS(c *SSContext)
}
