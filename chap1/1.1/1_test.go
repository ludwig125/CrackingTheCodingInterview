package main

import (
	"testing"
)

func TestCheckDuplicate(t *testing.T) {
	tests := map[string]struct {
		str  string
		want bool
	}{
		"abc": {
			str:  "abc",
			want: true,
		},
		"abb": {
			str:  "abb",
			want: false,
		},
		"abcdefghijklmnopa": {
			str:  "abcdefghijklmnopa",
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := checkDuplicate(tt.str); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if got := checkDuplicate2(tt.str); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			// fmt.Println(checkDuplicate(tt.str))
		})
	}
}

func checkDuplicate(s string) bool {
	m := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		w := string(s[i])
		if _, ok := m[w]; ok {
			return false
		}
		m[w] = true
	}
	return true
}

func checkDuplicate2(s string) bool {
	m := make(map[rune]bool)
	for _, v := range s {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = true
	}
	return true
}
