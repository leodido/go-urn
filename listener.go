package urn

import (
	"github.com/leodido/go-urn/grammar"
)

// Listener is ...
type Listener struct {
	*grammar.BaseUrnListener
	target *Urn
}

// NewListener is ...
func NewListener(target *Urn) *Listener {
	return &Listener{
		target: target,
	}
}

// ExitID is called when production id is exited
func (l *Listener) ExitID(ctx *grammar.IDContext) {
	l.target.ID = ctx.GetText()
}

// ExitSS is called when production ss is exited
func (l *Listener) ExitSS(ctx *grammar.SSContext) {
	l.target.SS = ctx.GetText()
}
