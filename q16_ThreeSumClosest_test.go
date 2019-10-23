package leecode

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	for _,unit := range []struct {
		input1 []int
		input2 int
		expect int
	} {
		{[]int{-1,2,1,-4}, 1, 2},
		{[]int{1,1,1,1}, 3, 3},
		{[]int{1,2,5,10,11}, 12, 13},
		{[]int{1,2,4,8,16,32,64,128},82,82},
	} {
		if actually := ThreeSumClosest(unit.input1, unit.input2); actually != unit.expect {
			t.Errorf("\ninput1: %v, input2: %v, expect: %v, \nactually: %v",
				unit.input1, unit.input2, unit.expect, actually)
		}
	}
}
