package parser;

import (
	"fmt"
	"testing"
	"time"
)

var zones = map[string]Zone{
	"Z1": Zone{45, 55},     // Active Recovery
	"Z2": Zone{55, 73},     // Endurance
	"Z3": Zone{73, 88},     // Tempo
	"Z4": Zone{88, 104},    // Threshold
	"Z5a": Zone{104, 119},  // VO2 Max
	"Z5b": Zone{119, 150},  // Anarobic
	"Z5c": Zone{150, 200},  // Neuromuscular
};

type parserResult struct {
	input string
	want string
	err error
}

func TestParser(t *testing.T) {
	fiveM := 5 * time.Minute;
	tests := []parserResult{
		{
			input: "5m Z1",
			want: fmt.Sprintf("%s Z1", fiveM),
			err: nil,
		},
		{
			input: "5m Z1, 5m Z2, 5m Z3",
			want: fmt.Sprintf("%s Z1 + %s Z2 + %s Z3", fiveM, fiveM, fiveM),
			err: nil,
		},
		{
			input: "5m Z1, 6 * (5m Z4, 5m Z5a), 5m Z3",
			want: fmt.Sprintf("%s Z1 + 6 * (%s Z4 + %s Z5a) + %v Z3", fiveM, fiveM, fiveM, fiveM),
			err: nil,
		},
		{
			input: "500 Z1",
			want: fmt.Sprintf("500.00 Z1"),
			err: nil,
		},
		{
			input: "500meters Z1",
			want: fmt.Sprintf("500.00 meters Z1"),
			err: nil,
		},
	};
	for _, input := range tests {
		t.Run(input.input, func(t *testing.T) {
			toks, err := Tokenize(input.input);
			if err != nil {
				t.Errorf("Failed to tokenize: %v", err);
				return;
			}
			p := New(toks, zones);
			i, err := p.Parse();
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
			if got, want := i.String(), input.want; got != want {
				t.Errorf("got: %v, want: %v", got, want);
			}
		});
	}
}
