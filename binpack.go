package main;

import (
	"atp_planner/binpack2d"
	"fmt"
)

type Area struct {
	width, height int
}

func New(w, h int) Area { return Area{w, h}; }

// Basic idea assume 5 zones:
// i := Z5
// goal = []int{Z_1, ..., Z_5}
// for each workout that satisfies Z_i:
//    goal -= workout
//    if goal == 0 || at max_sport: break
//    recurse(goal, i-1)
//    add to array
// repeat for other sports, then binpack on the
// cartesian product of the three

func main() {
	packer := binpack2d.Create(12, 12)
	workouts := []Area{
		New(5, 5),
		New(6, 5),
		New(7, 5),
		New(8, 5),
		New(9, 5),
		New(10, 5),
	};
	// TODO: Sort by area (or height), then try to fit in order
	for _, w := range workouts {
		rect, ok := packer.Insert(w.width, w.height, binpack2d.RULE_BEST_AREA_FIT);
		fmt.Printf("Added: %v %v\n", rect, ok);
		if !ok {
			break
		}
	}
	packer.ShrinkBin(false)
	for i := 0; i < packer.GetUsedRectanglesLength(); i += 1 {
		fmt.Printf("Rect[%d]=%v\n", i, packer.GetUsedRectangle(i));
	}
}
