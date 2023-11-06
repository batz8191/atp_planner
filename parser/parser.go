package parser;

import (
	"fmt"
)

type Parser struct {
	zones ZoneMap
	tokens []Token
	i int
}

// New constructs a new Parser
func New(tokens []Token, zones ZoneMap) *Parser {
	return &Parser{zones, tokens, 0};
}

// ZoneToName finds the zone for which the given percentage (between 0 and 100) lies.
func (p *Parser) ZoneToName(pct float64) string {
	for k, z := range p.zones {
		if pct >= float64(z.PowerLow) && pct <= float64(z.PowerHigh) {
			return k;
		}
	}
	return "U";
}

// Parse parses the tokens or returns an error
func (p *Parser) Parse() (Interval, error) {
	return p.union();
}

func (p *Parser) advance() {
	p.i++
}

func (p *Parser) peek(kind tokenType) bool {
	if p.i >= len(p.tokens) {
		return false;
	}
	t := p.tokens[p.i];
	if t.kind != kind {
		return false;
	}
	return true;
}

func (p *Parser) consume(kind tokenType) bool {
	if p.i >= len(p.tokens) {
		return false;
	}
	t := p.tokens[p.i];
	if t.kind != kind {
		return false;
	}
	p.advance();
	return true;
}

func (p *Parser) union() (Interval, error) {
	n, err := p.repeat();
	if err != nil {
		return nil, err
	}
	for {
		if p.consume(commaTok) {
			t, err := p.repeat();
			if err != nil {
				return nil, err;
			}
			n = newUnionInterval(p, n, t);
		} else {
			return n, nil;
		}
	}
	return n, nil;
}

func (p *Parser) repeat() (Interval, error) {
	n, err := p.basic()
	if err != nil {
		return nil, err
	}
	for {
		if p.consume(multiplyTok) {
			t, err := p.basic();
			if err != nil {
				return nil, err;
			}
			if v, ok := n.(*countInterval); ok {
				n = newRepeatInterval(p, v.count, t);
			} else if v, ok := t.(*countInterval); ok {
				n = newRepeatInterval(p, v.count, n);
			} else {
				return n, fmt.Errorf("expected one token to be a constant: %v, %v", n, t);
			}
		} else {
			return n, nil;
		}
	}
	return n, nil;
}

func (p *Parser) basic() (Interval, error) {
	if p.consume(lparenTok) {
		i, err := p.union();
		if err != nil {
			return nil, err
		}
		if !p.consume(rparenTok) {
			return nil, fmt.Errorf("Missing ')'");
		}
		return i, nil;
	}
	d := p.tokens[p.i];
	if d.kind != durationTok && d.kind != distanceTok && d.kind != constantTok {
		return nil, fmt.Errorf("Unexpected token type: %v\n", d);
	}
	p.advance();
	if d.kind == constantTok && !p.peek(zoneTok) {
		return newCount(p, d.con), nil;
	}
	if p.i >= len(p.tokens) {
		return nil, fmt.Errorf("Could not find a zone after a interval");
	}
	z := p.tokens[p.i];
	if z.kind != zoneTok {
		return nil, fmt.Errorf("Expecting a zone after a duration: %v", z);
	}
	p.advance();
	v, ok := p.zones[z.zone];
	if !ok {
		return nil, fmt.Errorf("Unknown zone: %v", z.zone);
	}
	switch d.kind {
	case constantTok:
		return newDistanceInterval(p, float64(d.con), "", v), nil;
	case distanceTok:
		return newDistanceInterval(p, d.dist, d.unit, v), nil;
	case durationTok:
		return newDurationInterval(p, d.dur, v), nil;
	}
	return nil, fmt.Errorf("Unknown type: %v", d.kind);
}
