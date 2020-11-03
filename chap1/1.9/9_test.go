package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestIsRotatedString(t *testing.T) {
	tests := map[string]struct {
		str1 string
		str2 string
		want bool
	}{
		"1": {
			str1: "waterbottle",
			str2: "erbottlewat",
			want: true,
		},
		"2": {
			str1: "Alfa Bravo Charlie Delta Echo Foxtrot Golf",
			str2: "lta Echo Foxtrot GolfAlfa Bravo Charlie De",
			want: true,
		},
		"3": {
			str1: "Alfa Bravo Charlie Delta Echo Foxtrot Golf",
			str2: " lta Echo Foxtrot GolfAlfa Bravo Charlie De",
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := isRotatedString(tt.str1, tt.str2)
			fmt.Println(got)
			if got != tt.want {
				t.Errorf("got: %t, want: %t", got, tt.want)
			}
		})
	}
}

func isRotatedString(str1, str2 string) bool {
	return str(str1 + str1).isSubstring(str2)
}

type str string

func (s str) isSubstring(target string) bool {
	return strings.Contains(string(s), target)
}
