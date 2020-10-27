package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCompress(t *testing.T) {
	tests := map[string]struct {
		str  string
		want string
	}{
		"1": {
			str:  "aabcccccaaa",
			want: "a2b1c5a3",
		},
		"2": {
			str:  "abca",
			want: "abca",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := compress(tt.str); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if got := compress2(tt.str); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// 新しく作成する文字列をns(new string)として空文字で定義する
// 文字列をスライスに格納して、1文字ずつループを回す
// 最初の文字として””（空文字）を定義しておき、prevとする
// 同じ文字が何回出現したか数えるためにcntというint変数を０で初期値定義しておく
// ループ内の処理：
// prevと同じ文字であれば、cntに１加算
// prevと異なる文字であれば、「prevの後ろにcntをつけた文字列」をnsに追加して、cntを１に書き換えて、prevをその文字で書き換える
func compress(s string) string {
	ns := ""
	prev := ""
	cnt := 0
	for _, v := range targetStrSlice(s) {
		if v == prev {
			cnt++
			continue
		}
		if prev != "" { // prevが最初に定義した空文字の時は何もしない
			ns = fmt.Sprintf("%s%s%d", ns, prev, cnt)
		}
		cnt = 1
		prev = v
	}
	fmt.Println("ns", ns)
	if len(ns) > len(s) {
		return s
	}
	return ns
}

// こちらのほうが、nssとしてスライスを用意して入れていくだけなので速い
func compress2(s string) string {
	var nss []string // new string slice
	prev := ""
	cnt := 0
	for _, v := range targetStrSlice(s) {
		if v == prev {
			cnt++
			continue
		}
		if prev != "" { // prevが最初に定義した空文字の時は何もしない
			//ns = fmt.Sprintf("%s%s%d", ns, prev, cnt)
			nss = append(nss, fmt.Sprintf("%s%d", prev, cnt))
		}
		cnt = 1
		prev = v
	}
	ns := strings.Join(nss, "")
	fmt.Println("ns", ns)
	if len(ns) > len(s) {
		return s
	}
	return ns
}

func targetStrSlice(s string) []string {
	// 普通に見ていくと一番最後の文字列だけ処理されないので、最後に別の文字列として空文字を入れておく
	// aabbbccdd - > a2b3c2で終わってしまう
	// aabbbccdd"" - > a2b3c2d2 最後が空文字列なので、その前のdの数を評価してくれる
	return append(strings.Split(s, ""), []string{""}...)
}
