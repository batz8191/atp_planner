package main;

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
	"sort"
	"atp_planner/parser"
)

var input = flag.String("input", "", "the file to read workouts from")

const ftp = 352.0; // Watts
var zones = parser.ZoneMap{
	"Z1": parser.Zone{45, 55},     // Active Recovery
	"Z2": parser.Zone{55, 73},     // Endurance
	"Z3": parser.Zone{73, 88},     // Tempo
	"Z4": parser.Zone{88, 104},    // Threshold
	"Z5a": parser.Zone{104, 119},  // VO2 Max
	"Z5b": parser.Zone{119, 150},  // Anarobic
	"Z5c": parser.Zone{150, 200},  // Neuromuscular
};

const thresholdPace = 10; // mph
var zonesRun = parser.ZoneMap{
	"Z1": parser.Zone{60, 74},     // Active Recovery
	"Z2": parser.Zone{73, 82},     // Endurance
	"Z3": parser.Zone{83, 92},     // Tempo
	"Z4": parser.Zone{92, 100},    // Threshold
	"Z5a": parser.Zone{100, 103},  // VO2 Max
	"Z5b": parser.Zone{103, 110},  // Anarobic
	"Z5c": parser.Zone{111, 125},  // Neuromuscular
};

func analyzeWorkout(typ, name, input string, mult float64, zones parser.ZoneMap, sortedZoneNames []string) (string, error) {
	toks, err := parser.Tokenize(input);
	if err != nil {
		return "", fmt.Errorf("Failed to tokenize: %v", err);
	}
	p := parser.New(toks, zones);
	i, err := p.Parse();
	if err != nil {
		return "", fmt.Errorf("Failed to parse: %v", err);
	}
	// Compute statistics
	w := i.Window1s();
	w30 := []float64{}
	for i := 0; i < len(w); i += 30 {
		sum := 0.0;
		for j := 0; j < 30; j += 1 {
			sum += mult * float64(w[i+j].PowerLow + w[i+j].PowerHigh) / 200.0;
		}
		w30 = append(w30, sum/30.0)
	}
	cnt := 0
	sum := 0.0;
	for i := 2; i < len(w30); i += 1 {
		tsum := 0.0
		for j := 2; j >= 0; j -= 1 {
			tsum += w30[i-j];
		}
		sum += math.Pow(tsum / 3.0, 4.0);
		cnt += 1;
	}
	np := math.Round(math.Pow(sum / float64(cnt), 1.0/4.0)*100) / 100;
	intensity := np / mult;
	// (# of seconds of the workout x Normalized Power x Intensity Factor) / (FTP x 3600) x 100
	tss := (float64(len(w)) * np * intensity / (mult * 3600)) * 100;
	// Print results
	body := []string{typ, name, input, fmt.Sprintf("%d", len(w))};
	body = append(body, fmt.Sprintf("%0.2f", np), fmt.Sprintf("%0.2f", intensity), fmt.Sprintf("%0.2f", tss));
	timeInZones := map[string]int64{}
	for k := range zones {
		timeInZones[k] = 0;
	}
	for _, v := range w {
		timeInZones[p.ZoneToName(float64(v.PowerLow + v.PowerHigh) / 2.0)] += 1;
	}
	for _, v := range sortedZoneNames {
		body = append(body, fmt.Sprintf("%d", timeInZones[v]));
	}
	return strings.Join(body, "\t"), nil;
}

///////////////////////////////////////////////////////////////////////////////
func main() {
	flag.Parse()
	zoneNames := []string{};
	for k := range zones {
		zoneNames = append(zoneNames, k);
	}
	sort.Strings(zoneNames);
	hdr := []string{"Type", "Name", "Details", "Total (s)", "NP", "IF", "TSS"}
	hdr = append(hdr, zoneNames...);
	fmt.Printf("%s\n", strings.Join(hdr, "\t"));
	// Read file
	readFile, err := os.Open(*input)
	if err != nil {
		fmt.Printf("Could not read file [%s]: %v\n", *input, err);
		return;
	}
	scanner := bufio.NewScanner(readFile)
	for scanner.Scan() {
		line := scanner.Text();
		data := strings.Split(line, "\t");
		if len(data) != 3 {
			fmt.Printf("Could not parse line: [%s]\n", line);
			return;
		}
		mult := ftp
		tz := zones;
		if data[0] == "run" {
			mult = thresholdPace;
			tz = zonesRun;
		}
		body, err := analyzeWorkout(data[0], data[1], data[2], mult, tz, zoneNames);
		if err != nil {
			fmt.Printf("%v\n", err);
			return;
		}
		fmt.Printf("%s\n", body);
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// TODO test a distance workout
}
