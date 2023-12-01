// Code generated from ./parser2/Calc.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calc

import "github.com/antlr4-go/antlr/v4"

// BaseCalcListener is a complete listener for a parse tree produced by CalcParser.
type BaseCalcListener struct{}

var _ CalcListener = &BaseCalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseCalcListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseCalcListener) ExitStart(ctx *StartContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCalcListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCalcListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRepeat is called when production repeat is entered.
func (s *BaseCalcListener) EnterRepeat(ctx *RepeatContext) {}

// ExitRepeat is called when production repeat is exited.
func (s *BaseCalcListener) ExitRepeat(ctx *RepeatContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseCalcListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseCalcListener) ExitAtom(ctx *AtomContext) {}

// EnterZone is called when production zone is entered.
func (s *BaseCalcListener) EnterZone(ctx *ZoneContext) {}

// ExitZone is called when production zone is exited.
func (s *BaseCalcListener) ExitZone(ctx *ZoneContext) {}

// EnterRange is called when production range is entered.
func (s *BaseCalcListener) EnterRange(ctx *RangeContext) {}

// ExitRange is called when production range is exited.
func (s *BaseCalcListener) ExitRange(ctx *RangeContext) {}

// EnterDur is called when production dur is entered.
func (s *BaseCalcListener) EnterDur(ctx *DurContext) {}

// ExitDur is called when production dur is exited.
func (s *BaseCalcListener) ExitDur(ctx *DurContext) {}

// EnterDist is called when production dist is entered.
func (s *BaseCalcListener) EnterDist(ctx *DistContext) {}

// ExitDist is called when production dist is exited.
func (s *BaseCalcListener) ExitDist(ctx *DistContext) {}
