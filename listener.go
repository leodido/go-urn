package urn

import (
	"github.com/leodido/go-urn/grammar"
)

// Listener is ...
type Listener struct {
	*grammar.BaseUrnListener
	target *URN
}

// NewListener is ...
func NewListener(u *URN) *Listener {
	return &Listener{
		target: u,
	}
}

// ExitID is called when rule iD is exited
func (l *Listener) ExitID(ctx *grammar.IDContext) {
	l.target.ID = ctx.GetText()
}

// ExitSS is called when rule sS is exited
func (l *Listener) ExitSS(ctx *grammar.SSContext) {
	l.target.SS = ctx.GetText()
}

// ExitUrn is called when rule urn is exited
func (l *Listener) ExitUrn(ctx *grammar.UrnContext) {
	l.target.tree = ctx.ToStringTree(nil, ctx.GetParser())
}
