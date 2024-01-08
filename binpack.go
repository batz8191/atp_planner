package main;

import (
	//"atp_planner/binpack2d"
	"atp_planner/fixedset"
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	//"sort"
	"strconv"
	"strings"
	"time"
)

var input = flag.String("input", "", "the file to read workouts from")
var typ = flag.String("type", "bike", "which type of workout to plan for")
var phases = flag.String("phase", "Base", "the phase of the ATP")
var hoursByWeek = flag.String("hours", "14h", "the goal hours for the week")
var TSSByWeek = flag.String("tss", "870", "the goal tss for the week")
var maxWorkouts = flag.Int("max_workouts", 4, "how many workouts to find")

// TODO:
// Take a file with phase, hours
const maxNonLongWorkout = 1 * time.Hour + 45 * time.Minute;

// Breakdown of percentages (50 == 50%) of time for each sport
var WeeklyBreakdown = map[string]map[string]int{
	"Base": {
		"bike": 50,
		"run": 30,
		"swim": 20,
	},
	"Build": {
		"bike": 52,
		"run": 31,
		"swim": 17,
	},
	"Peak": {
		"bike": 52,
		"run": 31,
		"swim": 17,
	},
	"Race": {
		"bike": 52,
		"run": 31,
		"swim": 17,
	},
	"Transition": {
		"bike": 50,
		"run": 30,
		"swim": 20,
	},
};

// Distribution represents the percentage of time spent in each zoen
type Distribution struct {
	Zone1 int
	Zone2 int
	Zone3 int
	Zone4 int
	Zone5a int
	Zone5b int
	Zone5c int
};

// The distribution expected for each phase.  Zone1 and Zone2 are
// combined.
var distribution = map[string]Distribution{
	"Base": Distribution{
		Zone1: 80,
		Zone2: 0,
		Zone3: 10,
		Zone4: 7,
		Zone5a: 2,
		Zone5b: 1,
		Zone5c: 0,
	},
	"Build": Distribution{
		Zone1: 66,
		Zone2: 0,
		Zone3: 23,
		Zone4: 8,
		Zone5a: 2,
		Zone5b: 1,
		Zone5c: 0,
	},
	"Peak": Distribution{
		Zone1: 66,
		Zone2: 0,
		Zone3: 20,
		Zone4: 10,
		Zone5a: 3,
		Zone5b: 1,
		Zone5c: 0,
	},
	"Taper": Distribution{
		Zone1: 90,
		Zone2: 0,
		Zone3: 0,
		Zone4: 7,
		Zone5a: 3,
		Zone5b: 0,
		Zone5c: 0,
	},
	"Race": Distribution{
		Zone1: 90,
		Zone2: 0,
		Zone3: 0,
		Zone4: 7,
		Zone5a: 3,
		Zone5b: 0,
		Zone5c: 0,
	},
};


///////////////////////////////////////////////////////////////////////////////
// Zones represents the time in the zones
type Zones struct {
	Z1, Z2, Z3, Z4, Z5a, Z5b, Z5c time.Duration
}

func (z *Zones) Add(o *Zones) Zones {
	return Zones{
		Z1: z.Z1 + o.Z1,
		Z2: z.Z2 + o.Z2,
		Z3: z.Z3 + o.Z3,
		Z4: z.Z4 + o.Z4,
		Z5a: z.Z5a + o.Z5a,
		Z5b: z.Z5b + o.Z5b,
		Z5c: z.Z5c + o.Z5c,
	};
}

func (z *Zones) Sub(o *Zones) Zones {
	return Zones{
		Z1: z.Z1 - o.Z1,
		Z2: z.Z2 - o.Z2,
		Z3: z.Z3 - o.Z3,
		Z4: z.Z4 - o.Z4,
		Z5a: z.Z5a - o.Z5a,
		Z5b: z.Z5b - o.Z5b,
		Z5c: z.Z5c - o.Z5c,
	};
}

func (z *Zones) Iterate() []time.Duration {
	return []time.Duration{z.Z1, z.Z2, z.Z3, z.Z4, z.Z5a, z.Z5b, z.Z5c};
}

func (z *Zones) Less(o *Zones) bool {
	if z.Z5c < o.Z5c { return true; }
	if z.Z5b < o.Z5b { return true; }
	if z.Z5a < o.Z5a { return true; }
	if z.Z4 < o.Z4 { return true; }
	if z.Z3 < o.Z3 { return true; }
	if z.Z2 < o.Z2 { return true; }
	if z.Z1 < o.Z1 { return true; }
	return false;
}

func (z *Zones) String() string {
	return strings.Join([]string{z.Z1.String(), z.Z2.String(), z.Z3.String(), z.Z4.String(), z.Z5a.String(), z.Z5b.String(), z.Z5c.String()}, "\t");
}

///////////////////////////////////////////////////////////////////////////////
// Workout represents a single workout
type Workout struct {
	Type string
	Name string
	Details string
	TotalDur time.Duration
	NP float64
	IF float64
	TSS float64
	Zones *Zones
};

func (w Workout) String() string {
	return fmt.Sprintf("%s\t%s\t%s", w.Name, w.TotalDur, w.Zones)
}

///////////////////////////////////////////////////////////////////////////////
// Workouts represents a combination of workouts
type Workouts struct {
	w []Workout;
	eff float64;
	tss float64;
}
func NewEmptyPlan() *Workouts {
	return &Workouts{
		w: make([]Workout, 0),
		eff: 0.0,
	}
}
func NewPlan(w []Workout, eff float64) *Workouts {
	tss := 0.0
	for _, t := range w {
		tss += t.TSS
	}
	return &Workouts{
		w: w,
		eff: eff,
		tss: tss,
	}
}
func (w *Workouts) Len() int { return len(w.w); }
func (w *Workouts) Efficiency() float64 { return w.eff; }
func (w *Workouts) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Eff: %0.2f, tss: %0.2f\n", w.eff, w.tss))
	for _, o := range w.w {
		sb.WriteString(o.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

// Unused, will be used when optimized
//type WeeklyPlan []*Workouts
//func (w WeeklyPlan) Len() int      { return len(w) }
//func (w WeeklyPlan) Swap(i, j int) { w[i], w[j] = w[j], w[i] }
//func (w WeeklyPlan) Less(i, j int) bool {
	//if math.Abs(w[i].eff - w[j].eff) < 1e-9 {
		//return w[i].tss > w[j].tss
	//}
	//return w[i].eff < w[j].eff
//}

///////////////////////////////////////////////////////////////////////////////
// WorkoutMap maps a type to a set of workouts
type WorkoutMap map[string]*Workouts;

///////////////////////////////////////////////////////////////////////////////
// loadWorkouts reads the set of workouts from a file (TSV).
// <type, name, description, duration, np, if, tss, zones (in seconds)>
func loadWorkouts(input string) (WorkoutMap, error) {
	ret := WorkoutMap{};
	readFile, err := os.Open(input)
	if err != nil {
		return WorkoutMap{}, fmt.Errorf("Could not read file [%s]: %v\n", input, err);
	}
	scanner := bufio.NewScanner(readFile)
	scanner.Scan();  // First line is header
	for scanner.Scan() {
		line := scanner.Text();
		data := strings.Split(line, "\t");
		if len(data) != 14 {
			return WorkoutMap{}, fmt.Errorf("Could not parse line: [%s]\n", line);
		}
		dur, err := strconv.Atoi(data[3]);
		if err != nil {
			return WorkoutMap{}, fmt.Errorf("Failed to parse duration: %s due to %w", line, err);
		}
		np, err := strconv.ParseFloat(data[4], 64);
		if err != nil {
			return WorkoutMap{}, fmt.Errorf("Failed to parse np: %s due to %w", line, err);
		}
		iF, err := strconv.ParseFloat(data[5], 64);
		if err != nil {
			return WorkoutMap{}, fmt.Errorf("Failed to parse if: %s due to %w", line, err);
		}
		tss, err := strconv.ParseFloat(data[6], 64);
		if err != nil {
			return WorkoutMap{}, fmt.Errorf("Failed to parse tss: %s due to %w", line, err);
		}
		zones := make([]time.Duration, 7);
		for i := 0; i < len(zones); i++ {
			dur, err := strconv.Atoi(data[7+i]);
			if err != nil {
				return WorkoutMap{}, fmt.Errorf("Failed to parse Z%d: %s due to %w", (i+1), line, err);
			}
			zones[i] = time.Duration(dur) * time.Second;
		}
		w, ok := ret[data[0]];
		if !ok {
			w = NewEmptyPlan();
			ret[data[0]] = w;
		}
		w.w = append(w.w, Workout{
			Type: data[0],
			Name: data[1],
			Details: data[2],
			TotalDur: time.Duration(dur) * time.Second,
			NP: np,
			IF: iF,
			TSS: tss,
			Zones: &Zones{
				Z1: zones[0],
				Z2: zones[1],
				Z3: zones[2],
				Z4: zones[3],
				Z5a: zones[4],
				Z5b: zones[5],
				Z5c: zones[6],
			},
		});
		ret[data[0]] = w;
	}
	return ret, nil
}

func multDur(d time.Duration, frac int) time.Duration {
	return time.Duration(float64(d.Seconds()) * float64(frac) / 100.0) * time.Second;
}

func diff(observed, expected float64) float64 {
	f := math.Max(expected, 60.0)
	return math.Abs(observed - expected) / math.Abs(f)
}

func efficiency(observed, desired Zones) float64 {
	return math.Pow((observed.Z1.Seconds() + observed.Z2.Seconds()) - (desired.Z1.Seconds() + desired.Z2.Seconds()), 2) +
		math.Pow(observed.Z3.Seconds() - desired.Z3.Seconds(), 2) +
		math.Pow(observed.Z4.Seconds() - desired.Z4.Seconds(), 2) +
		math.Pow(observed.Z5a.Seconds() - desired.Z5a.Seconds(), 2) +
		math.Pow(observed.Z5b.Seconds() - desired.Z5b.Seconds(), 2) +
		math.Pow(observed.Z5c.Seconds() - desired.Z5c.Seconds(), 2)
}

///////////////////////////////////////////////////////////////////////////////
// enumerate finds the best workouts matching the given goal of Zones
var best *fixedset.FixedSet[*Workouts]
func enumerate(foundLong bool, maxWorkouts int, goal Zones, all, list *Workouts) {
	// TODO:
	// optimize
	// make recursive
	totalCount := 0
	maxIterations := math.Pow(float64(len(all.w)), float64(maxWorkouts))
	toPrint := int(maxIterations / 100.0)
	nextPrint := toPrint
	fmt.Fprintf(os.Stderr, "maxIterations: %0.2f, toPrint: %d, nextPrint: %d\n", maxIterations, toPrint, nextPrint)
	for i := 0; i < len(all.w); i++ {
		a := all.w[i].Zones
		foundLong := all.w[i].TotalDur > maxNonLongWorkout
		for j := i+1; j < len(all.w); j++ {
			if all.w[j].TotalDur > maxNonLongWorkout && foundLong {
				totalCount++
				continue
			}
			foundLong = foundLong || all.w[j].TotalDur > maxNonLongWorkout
			b := a.Add(all.w[j].Zones)
			for k := j+1; k < len(all.w); k++ {
				if all.w[k].TotalDur > maxNonLongWorkout && foundLong {
					totalCount++
					continue
				}
				foundLong = foundLong || all.w[k].TotalDur > maxNonLongWorkout
				c := b.Add(all.w[k].Zones)
				for l := k+1; l < len(all.w); l++ {
					if all.w[l].TotalDur > maxNonLongWorkout && foundLong {
						totalCount++
						continue
					}
					d := c.Add(all.w[l].Zones)
					list := NewPlan([]Workout{all.w[i], all.w[j], all.w[k], all.w[l]}, efficiency(d, goal))
					best.Add(list);
					totalCount++
					if totalCount > nextPrint {
						nextPrint += toPrint
						fmt.Fprintf(os.Stderr, "nextPrint: %d\n", nextPrint)
						fmt.Fprintf(os.Stderr, "Done: %0.2f\n", 100.0 * float64(totalCount) / float64(maxIterations));
					}
				}
			}
		}
	}
}

func main() {
	flag.Parse()
	workoutMap, err := loadWorkouts(*input);
	if err != nil {
		fmt.Printf("%v", err);
		return;
	}
	workouts, ok := workoutMap[*typ];
	if !ok {
		fmt.Printf("Unknown type: %v\n", *typ);
		return;
	}
	fmt.Printf("Have: %d workouts (enumerating %0.2f times)\n", workouts.Len(), math.Pow(float64(workouts.Len()), float64(*maxWorkouts)))
	p := strings.Split(*phases, ",")
	h := strings.Split(*hoursByWeek, ",")
	t := strings.Split(*TSSByWeek, ",")
	if len(p) != len(h) || len(h) != len(p) {
		fmt.Fprintf(os.Stderr, "Arrays must be the same length: phase=%d hours=%d tss=%d\n", len(p), len(h), len(t))
		os.Exit(1)
	}
	for i := 0; i < len(p); i++ {
		phase := p[i]
		hours, err := time.ParseDuration(h[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not parse hours: %s\n", err)
		}
		tss, err := strconv.ParseFloat(t[i], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not parse tss: %s\n", err)
		}
		dist, ok := distribution[phase];
		if !ok {
			fmt.Printf("Unknown phase: %v\n", phase);
			return;
		}
		weekDist, ok := WeeklyBreakdown[phase];
		if !ok {
			fmt.Printf("Unknown phase (weekly): %v\n", phase);
			return;
		}
		frac := float64(weekDist[*typ]) / 100.0;
		goal := Zones{
			Z1: multDur(time.Duration(float64(hours)*frac), dist.Zone1),
			Z2: multDur(time.Duration(float64(hours)*frac), dist.Zone2),
			Z3: multDur(time.Duration(float64(hours)*frac), dist.Zone3),
			Z4: multDur(time.Duration(float64(hours)*frac), dist.Zone4),
			Z5a: multDur(time.Duration(float64(hours)*frac), dist.Zone5a),
			Z5b: multDur(time.Duration(float64(hours)*frac), dist.Zone5b),
			Z5c: multDur(time.Duration(float64(hours)*frac), dist.Zone5c),
		}
		goalTSS := tss * frac
		best = fixedset.New[*Workouts](10, func(a, b *Workouts) bool {
			return a.Efficiency() < b.Efficiency()
		})
		enumerate(false, *maxWorkouts, goal, workouts, NewEmptyPlan())
		// Find the closest to the goal TSS
		var bestWorkouts []Workout
		bestDiff := math.MaxFloat64
		for _, b := range best.Iterate() {
			tss := 0.0;
			for _, w := range b.w {
				tss += w.TSS;
			}
			d := diff(tss, goalTSS)
			if d < bestDiff {
				bestWorkouts = b.w
				bestDiff = d
			}
		}
		fmt.Printf("Week[%d]:\n", i)
		bestTSS := 0.0;
		total := time.Duration(0)
		delta := goal
		for _, w := range bestWorkouts {
			star := ""
			if w.TotalDur > maxNonLongWorkout {
				star = "*"
			}
			fmt.Printf("\t%s%s\t%s\t%s\n", star, w.Name, w.TotalDur, w.Zones.String());
			bestTSS += w.TSS;
			total += w.TotalDur
			delta = delta.Sub(w.Zones)
		}
		fmt.Printf("total: %s (goal: %s)\ntss: %0.2f (goal: %0.2f)\ndiff: %0.2f\n", total, time.Duration(float64(hours)*frac), bestTSS, goalTSS, bestDiff)
		fmt.Printf("Remaining\t%s\nGoal\t%s\n\n", delta, goal)
	}
}
