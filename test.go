package main

import (
	"atp_planner/parser"
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
	//is := antlr.NewInputStream("5m Z1, 5 * (5m 70%-80%, 3 miles 50%)")
	is := antlr.NewInputStream("5m Z1, 5 * (5m 70%-80%, 3m 50%)")
	lexer := parser.NewCalcLexer(is)
	lexer.AddErrorListener(NewErrorListener());
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCalcParser(stream)
	p.Start_()
	fmt.Printf("Final: %s\n", parser.Result)
	z := parser.PowerData{
		FTP: 352,
		Zones: parser.ZoneMap{
			"Z1": parser.Zone{45, 55},     // Active Recovery
			"Z2": parser.Zone{55, 73},     // Endurance
			"Z3": parser.Zone{73, 88},     // Tempo
			"Z4": parser.Zone{88, 104},    // Threshold
			"Z5a": parser.Zone{104, 119},  // VO2 Max
			"Z5b": parser.Zone{119, 150},  // Anarobic
			"Z5c": parser.Zone{150, 200},  // Neuromuscular
		},
	}
	fmt.Printf("Zones:\n")
	w1s := parser.Result.Window1s(z)
	cnt := 0
	v := float64(0)
	for _, w := range w1s {
		if w != v {
			fmt.Printf("%d * %.0f\n", cnt, v)
			cnt = 0
			v = w
		}
		cnt++
	}
	fmt.Printf("%d * %.0f\n", cnt, v)
}
