package main

import (
	"atp_planner/parser2"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"
)

///////////////////////////////////////////////////////////////////////////////
type ErrorListener struct {
	*antlr.DiagnosticErrorListener;
}
func NewErrorListener() *ErrorListener {
	return &ErrorListener{antlr.NewDiagnosticErrorListener(true)};
}
func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	fmt.Printf("SyntaxError\n")
	os.Exit(1)
}
func (l *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	fmt.Printf("ReportAmbiguity\n")
	os.Exit(1)
}
func (l *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	fmt.Printf("ReportAttemptingFullContext\n")
	os.Exit(1)
}
func (l *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	fmt.Printf("ReportContextSensitivity\n")
	os.Exit(1)
}

func main() {
	is := antlr.NewInputStream("5m Z1, 5 * (5m 70%-80%, 3 miles Z1)")
	lexer := parser.NewCalcLexer(is)
	lexer.AddErrorListener(NewErrorListener());
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCalcParser(stream)
	p.Start_()
	fmt.Printf("Final: %s\n", parser.Result)
}
