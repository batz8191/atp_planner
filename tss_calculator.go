package main

import (
	"atp_planner/parser"
	"bufio"
	"flag"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"math"
	"os"
	"sort"
	"strings"
)

var input = flag.String("input", "", "the file to read workouts from")
var output = flag.String("output", "", "the file to write workouts to")

const ftp = 352.0; // Watts
var zones = parser.ZoneMap{
	"Z1": parser.Zone{45, 55},     // Active Recovery
	"Z2": parser.Zone{55, 73},     // Endurance
	"Z3": parser.Zone{73, 88},     // Tempo
	"Z4": parser.Zone{88, 104},    // Threshold
	"Z5a": parser.Zone{104, 119},  // VO2 Max
	"Z5b": parser.Zone{119, 150},  // Anarobic
	"Z5c": parser.Zone{150, 200},  // Neuromuscular
}

const thresholdPace = 6 // min / mile
var zonesRun = parser.ZoneMap{
	"Z1": parser.Zone{60, 74},     // Active Recovery
	"Z2": parser.Zone{73, 82},     // Endurance
	"Z3": parser.Zone{83, 92},     // Tempo
	"Z4": parser.Zone{92, 100},    // Threshold
	"Z5a": parser.Zone{100, 103},  // VO2 Max
	"Z5b": parser.Zone{103, 110},  // Anarobic
	"Z5c": parser.Zone{111, 125},  // Neuromuscular
}

///////////////////////////////////////////////////////////////////////////////
type ErrorListener struct {
	*antlr.DiagnosticErrorListener
}
func NewErrorListener() *ErrorListener {
	return &ErrorListener{antlr.NewDiagnosticErrorListener(true)}
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

///////////////////////////////////////////////////////////////////////////////
func parseInput(name, input string) parser.Interval {
	fmt.Fprintf(os.Stderr, "Parsing: %s\n", name);
	is := antlr.NewInputStream(input)
	lexer := parser.NewCalcLexer(is)
	lexer.AddErrorListener(NewErrorListener())
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCalcParser(stream)
	p.SetErrorHandler(antlr.NewBailErrorStrategy())
	p.Start_()
	if p.HasError() {
		fmt.Fprintf(os.Stderr, "Parse failed: %v\n", p.GetError())
		os.Exit(1)
	}
	return parser.Result
}

func analyzeWorkout(typ parser.IntervalType, name, input string, p parser.PowerData, sortedZoneNames []string) (string, error) {
	// Compute statistics
	w := parseInput(name, input).Window1s(typ, p)
	cnt := 0
	sum := 0.0
	for i := 0; i < len(w)-30; i++ {
		tsum := 0.0
		for j := i; j < i+30 && j < len(w); j++ {
			tsum += w[j]
		}
		sum += math.Pow(tsum / 30.0, 4.0)
		cnt++
	}
	mult := parser.Multiplier(typ, p)
	np := math.Round(math.Pow(sum / float64(cnt), 1.0/4.0)*100) / 100
	intensity := np / mult
	// (# of seconds of the workout x Normalized Power x Intensity Factor) / (FTP x 3600) x 100
	tss := (float64(len(w)) * np * intensity / (mult * 3600)) * 100
	// Print results
	body := []string{typ.String(), name, input, fmt.Sprintf("%d", len(w))}
	body = append(body, fmt.Sprintf("%0.2f", np), fmt.Sprintf("%0.2f", intensity), fmt.Sprintf("%0.2f", tss))
	timeInZones := map[string]int64{}
	for k := range zones {
		timeInZones[k] = 0
	}
	for _, v := range w {
		timeInZones[p.ZoneToName(typ, v)] += 1
	}
	for _, v := range sortedZoneNames {
		body = append(body, fmt.Sprintf("%d", timeInZones[v]))
	}
	return strings.Join(body, "\t"), nil
}

///////////////////////////////////////////////////////////////////////////////
func main() {
	flag.Parse()
	zoneNames := []string{}
	for k := range zones {
		zoneNames = append(zoneNames, k)
	}
	sort.Strings(zoneNames)
	hdr := []string{"Type", "Name", "Details", "Total (s)", "NP", "IF", "TSS"}
	hdr = append(hdr, zoneNames...)
	fmt.Printf("%s\n", strings.Join(hdr, "\t"))
	readFile, err := os.Open(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read file [%s]: %v\n", *input, err)
		return
	}
	defer readFile.Close()
	outFile, err := os.Create(*output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not write file [%s]: %v\n", *output, err)
		return
	}
	defer outFile.Close()
	scanner := bufio.NewScanner(readFile)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, "\t")
		if len(data) != 3 {
			fmt.Fprintf(os.Stderr, "Could not parse line: [%s]\n", line)
			return
		}
		p := parser.PowerData {
			BikeFTP: ftp,
			BikeZones: zones,
			RunPace: thresholdPace,
			RunZones: zonesRun,
		}
		var typ parser.IntervalType
		switch data[0] {
		case parser.BikeInterval.String():
			typ = parser.BikeInterval
		case parser.RunInterval.String():
			typ = parser.RunInterval
		default:
			fmt.Fprintf(os.Stderr, "unknown type: %v in [%s]\n", typ, line)
			os.Exit(1)
		}
		body, err := analyzeWorkout(typ, data[1], data[2], p, zoneNames)
		if err != nil {
			fmt.Fprintf(outFile, "%v\n", err)
			return
		}
		fmt.Fprintf(outFile, "%s\n", body)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
}
