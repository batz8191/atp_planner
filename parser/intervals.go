package parser;

import (
	"fmt"
	"time"
)

type Zone struct {
	PowerLow int
	PowerHigh int
};

type ZoneMap map[string]Zone

///////////////////////////////////////////////////////////////////////////////
type Interval interface {
	// Returns the overall duration of the interval
	Duration() time.Duration
	// Returns the minimum and maximum zones
	Zone() Zone
	// Returns the array of 1s Zones for this interval.
	Window1s() []Zone
	// Returns a string representation of the interval.
	String() string
}

///////////////////////////////////////////////////////////////////////////////
type durationInterval struct {
	p *Parser
	d time.Duration
	power Zone
}

func newDurationInterval(p *Parser, d time.Duration, power Zone) *durationInterval {
	return &durationInterval{p, d, power};
}
func (n *durationInterval) Duration() time.Duration { return n.d; }
func (n *durationInterval) Zone() Zone { return n.power; }

func (n *durationInterval) Window1s() []Zone {
	d := int64(n.d.Seconds());
	r := make([]Zone, 0, d);
	for i := int64(0); i < d; i += 1 {
		r = append(r, n.power);
	}
	return r;
}

func (n *durationInterval) String() string {
	mid := float64(n.power.PowerLow + n.power.PowerHigh) / 2.0;
	return fmt.Sprintf("%s %s", n.d, n.p.ZoneToName(mid));
}

///////////////////////////////////////////////////////////////////////////////
type distanceInterval struct {
	p *Parser
	d float64
	unit string
	power Zone
}

func newDistanceInterval(p *Parser, d float64, u string, power Zone) *distanceInterval {
	return &distanceInterval{p, d, u, power};
}
func (n *distanceInterval) Duration() time.Duration {
	// TODO: convert to duration (distance * speed)
	return time.Duration(0);
}
func (n *distanceInterval) Zone() Zone { return n.power; }

func (n *distanceInterval) Window1s() []Zone {
	// TODO: Figure out how to handle this one
	d := int64(n.d);
	r := make([]Zone, 0, d);
	for i := int64(0); i < d; i += 1 {
		r = append(r, n.power);
	}
	return r;
}

func (d *distanceInterval) String() string {
	mid := float64(d.power.PowerLow + d.power.PowerHigh) / 2.0;
	t := ""
	if d.unit != "" {
		t = d.unit + " ";
	}
	return fmt.Sprintf("%0.2f %s%s", d.d, t, d.p.ZoneToName(mid));
}

///////////////////////////////////////////////////////////////////////////////
type unionInterval struct {
	p *Parser
	left, right Interval
}

func newUnionInterval(p *Parser, left, right Interval) *unionInterval {
	return &unionInterval{p, left, right};
}

func (u *unionInterval) Duration() time.Duration {
	return u.left.Duration() + u.right.Duration();
}
func (u *unionInterval) Zone() Zone {
	lz := u.left.Zone()
	rz := u.right.Zone()
	low := lz.PowerLow
	high := lz.PowerHigh
	if rz.PowerLow < low { low = rz.PowerLow; }
	if rz.PowerHigh > high { high = rz.PowerHigh; }
	return Zone{low, high};
}

func (u *unionInterval) Window1s() []Zone {
	return append(u.left.Window1s(), u.right.Window1s()...);
}

func (u *unionInterval) String() string {
	return fmt.Sprintf("%s + %s", u.left.String(), u.right.String());
}

///////////////////////////////////////////////////////////////////////////////
type countInterval struct {
	p *Parser
	count int
}

func newCount(p *Parser, c int) *countInterval {
	return &countInterval{p, c};
}

func (r *countInterval) Duration() time.Duration {
	return time.Duration(0);
}

func (r *countInterval) Zone() Zone {
	return Zone{0, 0};
}

func (r *countInterval) Window1s() []Zone {
	return []Zone{};
}

func (r *countInterval) String() string {
	return fmt.Sprintf("Count: %v", r.count);
}

///////////////////////////////////////////////////////////////////////////////
type repeatInterval struct {
	p *Parser
	count int
	interval Interval
}

func newRepeatInterval(p *Parser, count int, interval Interval) *repeatInterval {
	return &repeatInterval{p, count, interval};
}

func (r *repeatInterval) Duration() time.Duration {
	return time.Duration(int64(r.count) * int64(r.interval.Duration()));
}

func (r *repeatInterval) Zone() Zone {
	return r.interval.Zone();
}

func (r *repeatInterval) Window1s() []Zone {
	ret := []Zone{};
	iv := r.interval.Window1s();
	for i := 0; i < r.count; i += 1 {
		ret = append(ret, iv...)
	}
	return ret;
}

func (r *repeatInterval) String() string {
	return fmt.Sprintf("%d * (%s)", r.count, r.interval.String());
}
