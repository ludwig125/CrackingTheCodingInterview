package main

import (
	"strings"
	"testing"
)

func TestCanOnetimeConvert(t *testing.T) {
	tests := map[string]struct {
		str  string
		str2 string
		want bool
	}{
		"1": {
			str:  "pale",
			str2: "ple",
			want: true,
		},
		"2": {
			str:  "pales",
			str2: "pale",
			want: true,
		},
		"3": {
			str:  "pale",
			str2: "bale",
			want: true,
		},
		"4": {
			str:  "pale",
			str2: "bake",
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := canOnetimeConvert(tt.str, tt.str2); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if got := canOnetimeConvert2(tt.str, tt.str2); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// 一方の文字列がもう一方の文字列に一回の操作で変換可能か
// 操作とは、文字の挿入、文字の削除、文字の置き換えの３つ
// 考えることを減らすために、まず長いほうの文字列を基準に短い方の文字列に変換可能かを考える
// そうすると、文字の削除もしくは置き換えで済むかが考えられる
// このとき、文字列の長さが２以上異なっていたらその時点で無理

func canOnetimeConvert(s1, s2 string) bool {
	if s1 == s2 {
		return true // ２つの文字列が最初から同じだったらOK
	}
	ls1 := len(s1)
	ls2 := len(s2)
	if ls1 == ls2 {
		return isSameExceptOneWord(s1, s2)
	} else if ls1 == ls2+1 { // s1の方がs2より一文字だけ長い場合
		return isSameWhenDeleteOneWord(s1, s2)
	} else if ls1+1 == ls2 { // s2の方がs1より一文字だけ長い場合
		return isSameWhenDeleteOneWord(s2, s1)
	}
	return false
}

func canOnetimeConvert2(s1, s2 string) bool {
	if s1 == s2 {
		return true // ２つの文字列が最初から同じだったらOK
	}
	ls1 := len(s1)
	ls2 := len(s2)
	if ls1 == ls2 {
		return isSameExceptOneWord2(s1, s2)
	} else if ls1 == ls2+1 { // s1の方がs2より一文字だけ長い場合
		return isSameWhenDeleteOneWord2(s1, s2)
	} else if ls1+1 == ls2 { // s2の方がs1より一文字だけ長い場合
		return isSameWhenDeleteOneWord2(s2, s1)
	}
	return false
}

//　一文字を削除すると同じになるか調べる関数
// ２つの引数の文字列は常に１つ目の方が大きい必要がある
func isSameWhenDeleteOneWord(s1, s2 string) bool {
	for i := 0; i < len(s1); i++ {
		if strDeletedOneWord(s1, i) == s2 {
			return true
		}
	}
	return false
}

// 指定の要素番号の分を削除した文字列を返す関数
func strDeletedOneWord(s string, i int) string {
	ss := strings.Split(s, "")
	return strings.Join(append(ss[:i], ss[i+1:]...), "")
}

// 一文字を置き換えると同じになるか調べる関数
// ２つの文字列が１文字以外同じかどうかを調べればいい
func isSameExceptOneWord(s1, s2 string) bool {
	for i := 0; i < len(s1); i++ {
		if strDeletedOneWord(s1, i) == strDeletedOneWord(s2, i) {
			return true
		}
	}
	return false
}

// 以下、上とは別のやり方で解いたもの

//　一文字を削除すると同じになるか調べる関数
// ２つの引数の文字列は常に１つ目の方が大きい必要がある
func isSameWhenDeleteOneWord2(s1, s2 string) bool {
	foundDifference := false
	index := 0
	for i := 0; i < len(s2); i++ {
		// fmt.Println(string(s1[i]), string(s2[i+plus]), i, i+plus, foundDifference)
		// s1もs2も最初は同じ要素番号で比較するが、
		// 途中で文字の違いがあったら次のループからはs2の方が１つ前の要素番号で比較する必要がある
		if string(s1[i]) != string(s2[i+index]) {
			if foundDifference {
				return false
			}
			foundDifference = true
			index = -1
		}
	}
	return true
}

func isSameExceptOneWord2(s1, s2 string) bool {
	foundDifference := false
	for i := 0; i < len(s1); i++ {
		if string(s1[i]) != string(s2[i]) {
			if foundDifference {
				return false
			}
			foundDifference = true
		}
	}
	return true
}
