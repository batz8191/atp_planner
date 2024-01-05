package parser

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type IntervalType int

const (
	BikeInterval IntervalType = iota
	RunInterval
	//RunPowerInterval
	//SwimInterval
)

func (i IntervalType) String() string {
	switch i {
	case BikeInterval: return "bike"
	case RunInterval: return "run"
	}
	panic(fmt.Sprintf("Unknown type! [%v]", i))
	return ""
}

type Zone struct {
	PowerLow int
	PowerHigh int
};

type ZoneMap map[string]Zone

type Distance struct {
	Distance int
	Unit string
}

type PowerData struct {
	BikeFTP float64
	BikeZones ZoneMap
	RunPace float64 // min / mile
	RunZones ZoneMap
	// TODO: Swim, RunPower
}

// Very rough
func BikePowerToMph(pow float64) float64 {
	// From: https://www.gribble.org/cycling/power_v_speed.html
	// Drivetrain loss, as percent
	L_d := 2.0
	// Total Weight, in kg
	W := 83.91 + 9.07
	// Frontal area, in m^2
	A := 0.51
	// Coefficient of drag
	C_d := 0.63
	// Coefficient of rolling resistance
	C_rr := 0.005
	// Air Density, in \frac{kg}{m^3}
	rho := 1.23
	// Velocity Headwind, in \frac{m}{s}
	V_h := 0.0
	// Grade Percent
	G := 0.0
	// Intermediaries
	a := 1.0/2.0 * C_d * A * rho
	b := V_h * C_d * A * rho
	c := 9.8067 * W * (math.Sin(math.Atan(G / 100.0)) + C_rr * math.Cos(math.Atan(G/100.0))) + (1.0/2.0 * C_d * A * rho * V_h*V_h)
	d := -1.0 * (1.0 - (L_d/100.0)) * pow
	Q := (3.0*a*c) / (9*a*a)
	R := (9.0*a*b*c - 27*a*a*d - 2*b*b*b) / (54*a*a*a)
	S := math.Cbrt(R + math.Sqrt(Q*Q*Q + R*R))
	T := math.Cbrt(R - math.Sqrt(Q*Q*Q + R*R))
	meterPerSecToMPH := 2.23694
	return meterPerSecToMPH * (S + T - b / (3.0*a))
}

//func runPowerToMph(pow float64) float64 {
	//// TODO
	//return 0
//}

func minToMileToMph(runPace float64) float64 {
	return 60.0 / runPace
}

func Multiplier(typ IntervalType, p PowerData) float64 {
	switch typ {
	case BikeInterval: return p.BikeFTP
	case RunInterval: return minToMileToMph(p.RunPace)
	}
	panic(fmt.Sprintf("Unknown type! [%v]", typ))
	return 0.0
}

func zones(typ IntervalType, p PowerData) ZoneMap {
	switch typ {
	case BikeInterval: return p.BikeZones
	case RunInterval: return p.RunZones
	}
	panic(fmt.Sprintf("Unknown type! [%v]", typ))
	return ZoneMap{}
}

func (p PowerData) ZoneToName(typ IntervalType, pow float64) string {
	mult := Multiplier(typ, p)
	zones := zones(typ, p)
	for k, z := range zones {
		low := mult * float64(z.PowerLow) / 100.0
		high := mult * float64(z.PowerHigh) / 100.0
		if pow >= low && pow <= high {
			return k;
		}
	}
	return "Z1";
}

///////////////////////////////////////////////////////////////////////////////
type Power interface {
	String() string
	Mean(typ IntervalType, p PowerData) float64
}

type ZonePower struct {
	zone string
}
func (z *ZonePower) String() string { return z.zone; }
func (z *ZonePower) Mean(typ IntervalType, p PowerData) float64 {
	mult := Multiplier(typ, p)
	zone := zones(typ, p)[z.zone]
	return mult * float64(zone.PowerLow + zone.PowerHigh) / 2.0 / 100.0
}

type RangePower struct {
	low, high int
}
func (r *RangePower) String() string {
	if r.low == r.high {
		return fmt.Sprintf("%d%%", r.low);
	}
	return fmt.Sprintf("%d%%-%d%%", r.low, r.high);
}
func (r *RangePower) Mean(typ IntervalType, p PowerData) float64 {
	mult := Multiplier(typ, p)
	return mult * float64(r.low + r.high) / 2.0 / 100.0
}

///////////////////////////////////////////////////////////////////////////////
type Interval interface {
	String() string
	Window1s(typ IntervalType, p PowerData) []float64
}

type UnionInterval struct {
	interval []Interval
}
func (u *UnionInterval) String() string {
	s := []string{}
	for _, i := range u.interval {
		s = append(s, i.String())
	}
	return strings.Join(s, ", ")
}
func (u *UnionInterval) Window1s(typ IntervalType, p PowerData) []float64 {
	r := make([]float64, 0)
	for _, i := range u.interval {
		r = append(r, i.Window1s(typ, p)...)
	}
	return r
}

type RepeatInterval struct {
	count int
	interval Interval
}
func (r *RepeatInterval) String() string {
	return fmt.Sprintf("%d * (%s)", r.count, r.interval)
}
func (r *RepeatInterval) Window1s(typ IntervalType, p PowerData) []float64 {
	ret := make([]float64, 0)
	iv := r.interval.Window1s(typ, p)
	for i := 0; i < r.count; i++ {
		ret = append(ret, iv...)
	}
	return ret
}

type DurationInterval struct {
	dur time.Duration
	pow Power
}
func (d *DurationInterval) String() string {
	return fmt.Sprintf("%s %s", d.dur, d.pow)
}
func (d *DurationInterval) Window1s(typ IntervalType, p PowerData) []float64 {
	sec := int64(d.dur.Seconds())
	r := make([]float64, 0)
	for i := int64(0); i < sec; i++ {
		r = append(r, d.pow.Mean(typ, p))
	}
	return r
}

type DistanceInterval struct {
	dist int
	unit string
	pow Power
}
func (d *DistanceInterval) String() string {
	return fmt.Sprintf("%d%s %s", d.dist, d.unit, d.pow)
}
func (d *DistanceInterval) Window1s(typ IntervalType, p PowerData) []float64 {
	var mph float64
	switch typ {
	case BikeInterval:
		mph = BikePowerToMph(d.pow.Mean(typ, p))
	case RunInterval:
		mph = minToMileToMph(p.RunPace)
	}
	var sec int
	switch d.unit {
	case "mile", "miles":
		sec = int(math.Round(3600.0 * float64(d.dist) / mph))
	case "yard", "yards":
		sec = int(math.Round(3600.0 * float64(d.dist) / 1760.0 / mph))
	case "meter", "meters":
		sec = int(math.Round(3600.0 * float64(d.dist) / 1609.34 / mph))
	case "km":
		sec = int(math.Round(3600.0 * float64(d.dist) / 1.60934 / mph))
	}
	r := make([]float64, 0)
	for i := 0; i < sec; i++ {
		r = append(r, d.pow.Mean(typ, p))
	}
	return r
}
