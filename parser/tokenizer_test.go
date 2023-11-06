package parser;

import (
	"errors"
	"testing"
	"time"
)

type tokenResult struct {
	input string
	tokens []Token
	err error
}

func TestTokenizeBase(t *testing.T) {
	inputs := []tokenResult{
		// Basic loops
		{
			input: "*",
			tokens: []Token{{kind: multiplyTok}},
			err: nil,
		},
		{
			input: ",",
			tokens: []Token{{kind: commaTok}},
			err: nil,
		},
		{
			input: "(",
			tokens: []Token{{kind: lparenTok}},
			err: nil,
		},
		{
			input: ")",
			tokens: []Token{{kind: rparenTok}},
			err: nil,
		},
		{
			input: "5h30m30s",
			tokens: []Token{
				{
					kind: durationTok,
					dur: 5 * time.Hour + 30 * time.Minute + 30 * time.Second,
				},
			},
			err: nil,
		},
		{
			input: "5miles",
			tokens: []Token{
				{
					kind: distanceTok,
					dist: 5,
					unit: "miles",
				},
			},
			err: nil,
		},
		{
			input: "500meters",
			tokens: []Token{
				{
					kind: distanceTok,
					dist: 500,
					unit: "meters",
				},
			},
			err: nil,
		},
		{
			input: "500yards",
			tokens: []Token{
				{
					kind: distanceTok,
					dist: 500,
					unit: "yards",
				},
			},
			err: nil,
		},
		{
			input: "500",
			tokens: []Token{
				{
					kind: constantTok,
					con: 500,
				},
			},
			err: nil,
		},
		{
			input: "Z1",
			tokens: []Token{
				{
					kind: zoneTok,
					zone: "Z1",
				},
			},
			err: nil,
		},
		{
			input: "Z5a",
			tokens: []Token{
				{
					kind: zoneTok,
					zone: "Z5a",
				},
			},
			err: nil,
		},
		// Arrays
		{
			input: "5m Z1, 6 * ( 5m Z5a, 10m Z3), 5m Z1",
			tokens: []Token{
				{
					kind: durationTok,
					dur: 5 * time.Minute,
				},
				{
					kind: zoneTok,
					zone: "Z1",
				},
				{ kind: commaTok },
				{
					kind: constantTok,
					con: 6,
				},
				{ kind: multiplyTok },
				{ kind: lparenTok },
				{
					kind: durationTok,
					dur: 5 * time.Minute,
				},
				{
					kind: zoneTok,
					zone: "Z5a",
				},
				{ kind: commaTok },
				{
					kind: durationTok,
					dur: 10 * time.Minute,
				},
				{
					kind: zoneTok,
					zone: "Z3",
				},
				{ kind: rparenTok },
				{ kind: commaTok },
				{
					kind: durationTok,
					dur: 5 * time.Minute,
				},
				{
					kind: zoneTok,
					zone: "Z1",
				},
			},
			err: nil,
		},
		// Failures
		{
			input: "5us",
			err: errors.New(""),
		},
		{
			input: "5ms",
			err: errors.New(""),
		},
		{
			input: "Z5ac",
			err: errors.New(""),
		},
		{
			input: "Z5d",
			err: errors.New(""),
		},
	};
	for _, input := range inputs {
		t.Run(input.input, func(t *testing.T) {
			v, err := Tokenize(input.input)
			if err != nil && input.err == nil {
				t.Errorf("Unexpected error: %v", err);
				return;
			}
			if err == nil && input.err != nil {
				t.Errorf("Expected an error");
				return;
			}
			if err != nil {
				return;
			}
			if len(v) != len(input.tokens) {
				t.Errorf("got %v, want %v", len(v), len(input.tokens));
				return;
			}
			for i := range v {
				if v[i] != input.tokens[i] {
					t.Errorf("%d: got %v, want %v", i, v[i], input.tokens[i]);
				}
			}
		})
	}
}
