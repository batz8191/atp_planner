// Code generated from ./parser2/Calc.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calc

import "github.com/antlr4-go/antlr/v4"

// CalcListener is a complete listener for a parse tree produced by CalcParser.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRepeat is called when entering the repeat production.
	EnterRepeat(c *RepeatContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterZone is called when entering the zone production.
	EnterZone(c *ZoneContext)

	// EnterRange is called when entering the range production.
	EnterRange(c *RangeContext)

	// EnterDur is called when entering the dur production.
	EnterDur(c *DurContext)

	// EnterDist is called when entering the dist production.
	EnterDist(c *DistContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRepeat is called when exiting the repeat production.
	ExitRepeat(c *RepeatContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitZone is called when exiting the zone production.
	ExitZone(c *ZoneContext)

	// ExitRange is called when exiting the range production.
	ExitRange(c *RangeContext)

	// ExitDur is called when exiting the dur production.
	ExitDur(c *DurContext)

	// ExitDist is called when exiting the dist production.
	ExitDist(c *DistContext)
}
