package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := map[string]struct {
		str  string
		want bool
	}{
		"1": {
			str:  "Tact Coa",
			want: true,
		},
		"2": {
			str:  "Nurses run",
			want: true,
		},
		"3": {
			str:  "Some men interpret nine memos",
			want: true,
		},
		"4": {
			str:  "Cigar Toss it in a can It is so tragic",
			want: true,
		},
		"5": {
			str:  "Cigar? Toss it in a can It is so tragic",
			want: false,
		},
		"6": {
			str:  "A",
			want: true,
		},
		"7": {
			str:  "Ab",
			want: false,
		},
		"8": {
			str:  "Aba",
			want: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := isPalindrome(tt.str); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// 回文かどうかを判定するためには、ある文字列に含まれる文字のうち、
// 文字列の長さが偶数の場合：含まれる文字がすべて偶数である必要がある
// 文字列の長さが奇数の場合：含まれる文字が１つを除いてすべて偶数である必要がある
func isPalindrome(s string) bool {
	ss := toLowerCaseSlice(s)
	m := toMap(ss)
	fmt.Println(m)

	return areAllEvenExceptOne(m)
}

// すべて小文字のSliceにする
func toLowerCaseSlice(s string) []string {
	var ss []string
	for i := 0; i < len(s); i++ {
		ss = append(ss, strings.ToLower(string(s[i])))
	}
	return ss
}

func toMap(ss []string) map[string]int {
	m := make(map[string]int)
	for _, s := range ss {
		if s != " " {
			// fmt.Println("word", s)
			m[s]++
		}
	}
	return m
}

func areAllEvenExceptOne(m map[string]int) bool {
	foundOdd := false
	for _, v := range m {
		if v%2 != 0 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}
	return true
}

// func areAllEvenExceptOne(m map[string]int) bool {
// 	oddCnt := 0
// 	for _, v := range m {
// 		if v%2 != 0 {
// 			oddCnt++
// 			if oddCnt > 1 {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
