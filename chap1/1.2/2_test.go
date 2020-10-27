package main

import (
	"sort"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	tests := map[string]struct {
		str1 string
		str2 string
		want bool
	}{
		"1": {
			str1: "abc",
			str2: "cab",
			want: true,
		},
		"2": {
			str1: "? abc",
			str2: "c ?ba",
			want: true,
		},
		"3": {
			str1: "? abc",
			str2: "c ?bac",
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := isSortedSliceSame(tt.str1, tt.str2); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if got := isSortedSliceSame2(tt.str1, tt.str2); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if got := isSortedSliceSame3(tt.str1, tt.str2); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func isSortedSliceSame(s1, s2 string) bool {
	ss1 := toSlice(s1)
	sort.Strings(ss1)
	s1 = toString(ss1)

	ss2 := toSlice(s2)
	sort.Strings(ss2)
	s2 = toString(ss2)

	// fmt.Println(s1, s2)
	return s1 == s2
}

func toSlice(s string) []string {
	var ss []string
	for i := 0; i < len(s); i++ {
		ss = append(ss, string(s[i]))
	}
	return ss
}

func toString(ss []string) string {
	var s string
	for _, v := range ss {
		s = s + v
	}
	return s
}

func isSortedSliceSame2(s1, s2 string) bool {
	ss1 := strings.Split(s1, "")
	sort.Strings(ss1)
	s1 = strings.Join(ss1, "")

	ss2 := strings.Split(s2, "")
	sort.Strings(ss2)
	s2 = strings.Join(ss2, "")
	return s1 == s2
}

func isSortedSliceSame3(s1, s2 string) bool {
	ss1 := strings.Split(s1, "")
	sort.Slice(ss1, func(i int, j int) bool { return ss1[i] < ss1[j] })
	s1 = strings.Join(ss1, "")

	ss2 := strings.Split(s2, "")
	sort.Slice(ss2, func(i int, j int) bool { return ss2[i] < ss2[j] })
	s2 = strings.Join(ss2, "")
	return s1 == s2
}
