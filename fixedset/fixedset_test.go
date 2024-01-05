package fixedset;

import (
	"fmt"
	"reflect"
	"testing"
)

type test struct {
	input []int
	want []int
}

func Test(t *testing.T) {
	tests := []test{
		{
			input: []int{2, 1, 3, 5, 4},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			input: []int{6, 5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			input: []int{2, 5, 4, 1, 7, 6, 3},
			want: []int{1, 2, 3, 4, 5},
		},
	};
	for _, input := range tests {
		t.Run(fmt.Sprintf("%v", input.input), func(t *testing.T) {
			s := New[int](5, func(a, b int) bool { return a < b; });
			for _, i := range input.input {
				s.Add(i);
			}
			if got,  want := s.items, input.want; !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want);
			}
		});
	}
}
