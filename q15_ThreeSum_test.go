package leecode

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	for _, unit := range []struct {
		input []int
		expect [][]int
	} {
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1,-1,2},{-1,0,1}}},
		{[]int{0, 0, 0, 0,}, [][]int{{0,0,0}}},
		{[]int{0, 0, 0}, [][]int{{0,0,0}}},
		{[]int{-1,0,1,0}, [][]int{{-1,0,1}}},
		{[]int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6},
			[][]int{{-4,-2,6},{-4,0,4},{-4,1,3},{-4,2,2},{-2,-2,4},{-2,0,2}}},
	} {
		if actually := ThreeSum(unit.input); ! reflect.DeepEqual(actually, unit.expect) {
			t.Errorf("\ninput: %v, expect: %v, \nactually: %v", unit.input, unit.expect, actually)
		}
	}
}
