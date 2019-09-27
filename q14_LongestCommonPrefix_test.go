package leecode

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	for _, unit := range []struct {
		input []string
		expect string
	} {
		{[]string{"flower","flow","flight"}, "fl"},
		{[]string{"dog","racecar","car"}, ""},
		{[]string{""}, ""},
		{[]string{"a"}, "a"},
		{[]string{"c","c"}, "c"},
		{[]string{"aa", "aa"}, "aa"},
		{[]string{"a", "b"}, ""},
	} {
		if actually := LongestCommonPrefix(unit.input); actually != unit.expect {
			t.Errorf("\ninput: %v, expect: %v , \nactually: %v", unit.input, unit.expect, actually)
		}
	}
}
