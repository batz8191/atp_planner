package parser;

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type tokenType int;

// 5m Z1, 6 * (30m Z2, 5m Z1), 5m Z1
// 5miles Z1
// 500meters
// 500yards
// 5h Z1
// 5s Z5c
// 100 Z1 (Swim)
const durOrDistChars = "aedhilmrsty";
const distUniqueChars = "adeilrty";

const (
	durationTok tokenType = iota
	distanceTok
	constantTok
	zoneTok
	multiplyTok
	commaTok
	lparenTok
	rparenTok
)

type Token struct {
	kind tokenType
	dur time.Duration
	dist float64
	unit string
	con int
	zone string
}

func newConst(d string) (Token, error) {
	o, err := strconv.ParseInt(d, 10, 64);
	if err != nil {
		return Token{}, fmt.Errorf("Failed to parse: [%v] as a constant %w", d, err);
	}
	return Token{
		kind: constantTok,
		con: int(o),
	}, nil;
}

func newDur(d string) (Token, error) {
	o, err := time.ParseDuration(d);
	if err != nil {
		return Token{}, fmt.Errorf("Failed to parse: [%v] as a duration %w", d, err);
	}
	if o.Seconds() < 1.0 {
		return Token{}, fmt.Errorf("Cannot support sub-second durations: %v", d);
	}
	return Token{
		kind: durationTok,
		dur: o,
	}, nil;
}

func newDist(d, u string) (Token, error) {
	o, err := strconv.ParseFloat(d, 64);
	if err != nil {
		return Token{}, fmt.Errorf("Failed to parse: [%v] as a distance %w", d, err);
	}
	return Token{
		kind: distanceTok,
		dist: o,
		unit: u,
	}, nil;
}

func isNum(c rune) bool {
	return c >= '0' && c <= '9';
}

// Tokenize splits string into a sequence of Tokens if possible
func Tokenize(input string) ([]Token, error) {
	chars := []rune(input);
	i := 0;
	n := len(chars);
	tokens := []Token{};
	for i < n {
		c := chars[i];
		if unicode.IsSpace(c) {
			i += 1;
			continue;
		}
		if c == '*' {
			i += 1;
			tokens = append(tokens, Token{kind: multiplyTok});
			continue;
		}
		if c == ',' {
			i += 1;
			tokens = append(tokens, Token{kind: commaTok});
			continue;
		}
		if c == '(' {
			i += 1;
			tokens = append(tokens, Token{kind: lparenTok});
			continue;
		}
		if c == ')' {
			i += 1;
			tokens = append(tokens, Token{kind: rparenTok});
			continue;
		}
		// zone
		if c == 'Z' {
			end := i+1;
			suffixStart := -1;
			for ; end < n; end += 1 {
				c := chars[end];
				if unicode.IsSpace(c) {
					break;
				}
				if isNum(c) {
					continue;
				}
				if suffixStart == -1 && (c == 'a' || c == 'b' || c == 'c') {
					suffixStart = end;
					continue;
				}
				break;
			}
			var err error
			if suffixStart != -1 {
				_, err = strconv.Atoi(string(chars[i+1:suffixStart]));
			} else {
				_, err = strconv.Atoi(string(chars[i+1:end]));
			}
			if err != nil {
				return []Token{}, fmt.Errorf("Could not parse zone: %s\n", string(chars[i+1:end]));
			}
			tokens = append(tokens, Token{
				kind: zoneTok,
				zone: string(chars[i:end]),
			})
			i = end;
			continue;
		}
		// constant, duration or distance
		if isNum(c) {
			end := i;
			kind := constantTok;
			dstStart := -1;
			for ; end < n; end += 1 {
				c := chars[end];
				if unicode.IsSpace(c) {
					break;
				}
				if isNum(c) {
					continue;
				}
				if strings.ContainsRune(durOrDistChars, c) {
					if dstStart == -1 {
						kind = durationTok;
						dstStart = end;
					}
					if strings.ContainsRune(distUniqueChars, c) {
						kind = distanceTok;
					}
					continue;
				}
				break;
			}
			switch kind {
			case constantTok:
				t, err := newConst(string(chars[i:end]));
				if err != nil {
					return []Token{}, err;
				}
				tokens = append(tokens, t);
			case durationTok:
				t, err := newDur(string(chars[i:end]));
				if err != nil {
					return []Token{}, err;
				}
				tokens = append(tokens, t);
			case distanceTok:
				t, err := newDist(string(chars[i:dstStart]), string(chars[dstStart:end]));
				if err != nil {
					return []Token{}, err;
				}
				tokens = append(tokens, t);
			}
			i = end;
			continue;
		}
		return []Token{}, fmt.Errorf("Unknown character: %s\n", string(c));
	}
	return tokens, nil;
}
