package main;

import (
	"fmt"
	"math"
	"sort"
	"atp_planner/parser"
)

//const typ = "bike";
const typ = "run";
//const test = "10m Z1, 6 * (30m Z4, 10m Z1), 10m Z1";
const test = "20m Z1";

const ftp = 352.0; // Watts
var zones = map[string]parser.Zone{
	"Z1": parser.Zone{45, 55},     // Active Recovery
	"Z2": parser.Zone{55, 73},     // Endurance
	"Z3": parser.Zone{73, 88},     // Tempo
	"Z4": parser.Zone{88, 104},    // Threshold
	"Z5a": parser.Zone{104, 119},  // VO2 Max
	"Z5b": parser.Zone{119, 150},  // Anarobic
	"Z5c": parser.Zone{150, 200},  // Neuromuscular
};

const thresholdPace = 10; // mph
var zonesRun = map[string]parser.Zone{
	"Z1": parser.Zone{60, 74},     // Active Recovery
	"Z2": parser.Zone{73, 82},     // Endurance
	"Z3": parser.Zone{83, 92},     // Tempo
	"Z4": parser.Zone{92, 100},    // Threshold
	"Z5a": parser.Zone{100, 103},  // VO2 Max
	"Z5b": parser.Zone{103, 110},  // Anarobic
	"Z5c": parser.Zone{111, 125},  // Neuromuscular
};

///////////////////////////////////////////////////////////////////////////////
func main() {
	mult := ftp
	tz := zones;
	if typ == "run" {
		mult = thresholdPace;
		tz = zonesRun;
	}
	// Parse
	toks, err := parser.Tokenize(test);
	if err != nil {
		fmt.Printf("Failed to tokenize: %v", err);
		return;
	}
	p := parser.New(toks, tz);
	i, err := p.Parse();
	if err != nil {
		fmt.Printf("Failed to parse: %v", err);
		return;
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
		fmt.Printf("%v^4=%v\n", tsum / 3.0, math.Pow(tsum / 3.0, 4.0));
		cnt += 1;
	}
	np := math.Round(math.Pow(sum / float64(cnt), 1.0/4.0)*100) / 100;
	intensity := np / mult;
	// (# of seconds of the workout x Normalized Power x Intensity Factor) / (FTP x 3600) x 100
	tss := (float64(len(w)) * np * intensity / (mult * 3600)) * 100;
	// Print results
	fmt.Printf("Name\tTotal (s)\t");
	zoneNames := []string{};
	for k := range zones {
		zoneNames = append(zoneNames, k);
	}
	sort.Strings(zoneNames);
	for _, k := range zoneNames {
		fmt.Printf("%s\t", k);
	}
	fmt.Printf("NP\tIF\tTSS\n");
	fmt.Printf("<NAME>\t%v\t", len(w))
	// Calculate Time In Zones
	timeInZones := map[string]int64{}
	for k := range zones {
		timeInZones[k] = 0;
	}
	for _, v := range w {
		timeInZones[p.ZoneToName(float64(v.PowerLow + v.PowerHigh) / 2.0)] += 1;
	}
	for _, v := range zoneNames {
		fmt.Printf("%d\t", timeInZones[v]);
	}
	fmt.Printf("%0.2f\t%0.2f\t%0.2f\n", np, intensity, tss)
}
