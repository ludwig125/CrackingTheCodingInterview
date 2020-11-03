package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNinetyDegreeRotation(t *testing.T) {
	tests := map[string]struct {
		nn   [][]int
		want [][]int
	}{
		"1": {
			nn: [][]int{
				{1, 2},
				{3, 4},
			},
			want: [][]int{
				{3, 1},
				{4, 2},
			},
		},
		"2": {
			nn: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		"3": {
			nn: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: [][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
		"5": {
			nn: [][]int{
				{1, 2, 3, 4, 5, 6},
				{7, 8, 9, 10, 11, 12},
				{13, 14, 15, 16, 17, 18},
				{19, 20, 21, 22, 23, 24},
				{25, 26, 27, 28, 29, 30},
				{31, 32, 33, 34, 35, 36},
			},
			want: [][]int{
				{31, 25, 19, 13, 7, 1},
				{32, 26, 20, 14, 8, 2},
				{33, 27, 21, 15, 9, 3},
				{34, 28, 22, 16, 10, 4},
				{35, 29, 23, 17, 11, 5},
				{36, 30, 24, 18, 12, 6},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// fmt.Println(tt.nn)
			printMatrix(tt.nn)
			// printMatrix(ninetyDegreeRotation(tt.nn))
			ninetyDegreeRotation2(tt.nn)
			printMatrix(tt.nn)
			if !reflect.DeepEqual(tt.nn, tt.want) {
				t.Errorf("got: %v, want: %v", tt.nn, tt.want)
			}
		})
	}
}

func printMatrix(nn [][]int) {
	for _, v1 := range nn {
		for _, v2 := range v1 {
			fmt.Printf("%d, ", v2)
		}
		fmt.Println()
	}
}

// 以下の配列の場合、
// {1, 2, 3, 4},
// {5, 6, 7, 8},
// {9, 10, 11, 12},
// {13, 14, 15, 16},
func ninetyDegreeRotation(nn [][]int) [][]int {
	if len(nn) == 2 {
		return ninetyDegreeRotation2And2(nn)
	}
	if len(nn) == 3 {
		return ninetyDegreeRotation3And3(nn)
	}
	if len(nn) == 4 {
		return ninetyDegreeRotation4And4(nn)
	}
	return nn
}

// 2*2の行列の90度回転
func ninetyDegreeRotation2And2(nn [][]int) [][]int {
	begin := 0 // 0始まり
	end := 1   // 1で終わり
	tmp := nn[begin][begin]
	nn[begin][begin] = nn[end][begin]
	nn[end][begin] = nn[end][end]
	nn[end][end] = nn[begin][end]
	nn[begin][end] = tmp
	return nn
}

// 3*3の行列の90度回転
func ninetyDegreeRotation3And3(nn [][]int) [][]int {
	begin := 0 // 0始まり
	end := 2   // 2で終わり
	// このループは下のコメントにした処理を一般化したもの
	for offset := 0; offset < end; offset++ {
		tmp := nn[begin][begin+offset]
		nn[begin][begin+offset] = nn[end-offset][begin]
		nn[end-offset][begin] = nn[end][end-offset]
		nn[end][end-offset] = nn[begin+offset][end]
		nn[begin+offset][end] = tmp
	}

	// tmp := nn[begin][begin]
	// nn[begin][begin] = nn[end][begin]
	// nn[end][begin] = nn[end][end]
	// nn[end][end] = nn[begin][end]
	// nn[begin][end] = tmp

	// offset := 1
	// tmp = nn[begin][begin+offset]
	// nn[begin][begin+offset] = nn[end-offset][begin]
	// nn[end-offset][begin] = nn[end][end-offset]
	// nn[end][end-offset] = nn[begin+offset][end]
	// nn[begin+offset][end] = tmp

	return nn
}

// 4*4の行列の90度回転
func ninetyDegreeRotation4And4(nn [][]int) [][]int {
	// これは下のコメントにした処理を一般化したもの
	begin := 0         // 0始まり
	end := len(nn) - 1 //  3で終わり
	for layer := 0; layer < len(nn)/2; layer++ {
		begin = begin + layer
		end = end - layer
		// 外周の一周の数字を９０度回転
		for offset := begin; offset < end; offset++ {
			tmp := nn[begin][begin+offset]
			nn[begin][begin+offset] = nn[end-offset][begin]
			nn[end-offset][begin] = nn[end][end-offset]
			nn[end][end-offset] = nn[begin+offset][end]
			nn[begin+offset][end] = tmp
		}
	}

	// begin := 0         // 0始まり
	// end := len(nn) - 1 // 3で終わり
	// // 外周の一周の数字を９０度回転
	// for offset := begin; offset < end; offset++ {
	// 	tmp := nn[begin][begin+offset]
	// 	nn[begin][begin+offset] = nn[end-offset][begin]
	// 	nn[end-offset][begin] = nn[end][end-offset]
	// 	nn[end][end-offset] = nn[begin+offset][end]
	// 	nn[begin+offset][end] = tmp
	// }

	// layer := 1
	// // 内側の一周の数字を９０度回転
	// begin = begin + layer // 1始まり
	// end = end - layer     // 2で終わり
	// tmp := nn[begin][begin]
	// nn[begin][begin] = nn[end][begin]
	// nn[end][begin] = nn[end][end]
	// nn[end][end] = nn[begin][end]
	// nn[begin][end] = tmp

	return nn
}

// 上のを完全に一般化すると以下になる
func ninetyDegreeRotation2(nn [][]int) {
	for layer := 0; layer < len(nn)/2; layer++ {
		begin := 0 + layer
		end := len(nn) - 1 - layer
		// 外周の一周の数字を９０度回転
		for offset := 0; offset < (end - begin); offset++ {
			tmp := nn[begin][begin+offset]
			nn[begin][begin+offset] = nn[end-offset][begin]
			nn[end-offset][begin] = nn[end][end-offset]
			nn[end][end-offset] = nn[begin+offset][end]
			nn[begin+offset][end] = tmp
		}
	}
}
