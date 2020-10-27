package main

import (
	"testing"
)

func TestURLify(t *testing.T) {
	tests := map[string]struct {
		str  string
		want string
	}{
		"1": {
			str:  "Mr John Smith ",
			want: "Mr%20John%20Smith",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := urlify(tt.str); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func urlify(s string) string {
	// fmt.Println(s)
	var ss []string
	origin := toSlice(s)
	// fmt.Println(origin, len(origin))
	origin = trim(reverse(origin))
	// fmt.Println(origin, len(origin))

	for _, v := range origin {
		ss = append(ss, replace(v))
	}

	return toString(reverse(ss))
}

func toSlice(s string) []string {
	var ss []string
	for i := 0; i < len(s); i++ {
		ss = append(ss, string(s[i]))
	}
	return ss
}

func trim(origin []string) []string {
	cnt := 0
	var ss []string
	for _, v := range origin {
		if v == " " && cnt == 0 {
			continue
		}
		ss = append(ss, v)
		cnt++
	}
	return ss
}

func replace(s string) string {
	if s == " " {
		return "%20"
	}
	return s
}

func reverse(origin []string) []string {
	var ss []string
	for i := len(origin) - 1; i >= 0; i-- {
		ss = append(ss, origin[i])
	}
	return ss
}

// 上のreverseは以下でもいい
func reverse2(ss []string) []string {
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
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
