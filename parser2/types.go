package parser

import (
	"fmt"
	"strings"
	"time"
)

type Distance struct {
	Distance int
	Unit string
}

///////////////////////////////////////////////////////////////////////////////
type Power interface {
	String() string
}

type ZonePower struct {
	zone string
}
func (z *ZonePower) String() string { return z.zone; }

type RangePower struct {
	low, high int
}
func (r *RangePower) String() string { return fmt.Sprintf("%d%%-%d%%", r.low, r.high); }

///////////////////////////////////////////////////////////////////////////////
type Interval interface {
	String() string
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

type RepeatInterval struct {
	count int
	interval Interval
}
func (d *RepeatInterval) String() string {
	return fmt.Sprintf("%d * (%s)", d.count, d.interval)
}

type DurationInterval struct {
	dur time.Duration
	pow Power
}
func (d *DurationInterval) String() string {
	return fmt.Sprintf("%s %s", d.dur, d.pow)
}

type DistanceInterval struct {
	dist int
	unit string
	pow Power
}
func (d *DistanceInterval) String() string {
	return fmt.Sprintf("%d%s %s", d.dist, d.unit, d.pow)
}

