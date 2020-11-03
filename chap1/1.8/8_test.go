package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPadding(t *testing.T) {
	tests := map[string]struct {
		mn   [][]int
		want [][]int
	}{
		"1": {
			mn: [][]int{
				{0, 2, 3},
				{4, 5, 6},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 5, 6},
			},
		},
		"2": {
			mn: [][]int{
				{1, 0, 3},
				{4, 5, 6},
			},
			want: [][]int{
				{0, 0, 0},
				{4, 0, 6},
			},
		},
		"3": {
			mn: [][]int{
				{1, 2, 3},
				{4, 5, 0},
			},
			want: [][]int{
				{1, 2, 0},
				{0, 0, 0},
			},
		},
		"4": {
			mn: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 0, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
			},
			want: [][]int{
				{1, 2, 3, 0, 5},
				{0, 0, 0, 0, 0},
				{11, 12, 13, 0, 15},
				{16, 17, 18, 0, 20},
			},
		},
		"5": {
			mn: [][]int{
				{1, 2, 3, 4, 5, 6},
				{7, 8, 9, 10, 11, 12},
				{13, 14, 15, 16, 17, 18},
				{19, 20, 21, 22, 23, 24},
				{25, 26, 27, 28, 29, 30},
				{31, 32, 33, 34, 35, 0},
			},
			want: [][]int{
				{1, 2, 3, 4, 5, 0},
				{7, 8, 9, 10, 11, 0},
				{13, 14, 15, 16, 17, 0},
				{19, 20, 21, 22, 23, 0},
				{25, 26, 27, 28, 29, 0},
				{0, 0, 0, 0, 0, 0},
			},
		},
		"6": {
			mn: [][]int{
				{1, 2, 3, 4, 5, 6},
				{7, 8, 9, 10, 11, 12},
				{13, 0, 15, 16, 17, 18},
				{19, 20, 21, 22, 23, 24},
				{25, 26, 27, 28, 29, 30},
				{31, 32, 33, 34, 35, 0},
			},
			want: [][]int{
				{1, 0, 3, 4, 5, 0},
				{7, 0, 9, 10, 11, 0},
				{0, 0, 0, 0, 0, 0},
				{19, 0, 21, 22, 23, 0},
				{25, 0, 27, 28, 29, 0},
				{0, 0, 0, 0, 0, 0},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// fmt.Println(tt.mn)
			printMatrix(tt.mn)
			// printMatrix(ninetyDegreeRotation(tt.mn))
			nmn := padding(tt.mn)
			printMatrix(nmn)
			if !reflect.DeepEqual(nmn, tt.want) {
				t.Errorf("got: %v, want: %v", nmn, tt.want)
			}

			mn := padding2(tt.mn)
			printMatrix(mn)
			if !reflect.DeepEqual(mn, tt.want) {
				t.Errorf("got: %v, want: %v", mn, tt.want)
			}
		})
	}
}

func printMatrix(mn [][]int) {
	for _, v1 := range mn {
		for _, v2 := range v1 {
			fmt.Printf("%d, ", v2)
		}
		fmt.Println()
	}
}

func padding(mn [][]int) [][]int {
	var row []int
	var column []int
	for i, v1 := range mn {
		for j, v2 := range v1 {
			if v2 == 0 {
				row = append(row, i)
				column = append(column, j)
			}
		}
	}

	nmn := make([][]int, len(mn)) // new m*n
	for i := 0; i < len(mn); i++ {
		nmn[i] = make([]int, len(mn[0]))
		for j := 0; j < len(mn[0]); j++ {
			if isIn(row, i) || isIn(column, j) {
				nmn[i][j] = 0
				continue
			}
			nmn[i][j] = mn[i][j]
		}
	}
	return nmn
}

func isIn(is []int, i int) bool {
	for _, v := range is {
		if v == i {
			return true
		}
	}
	return false
}

func padding2(mn [][]int) [][]int {
	row := make([]bool, len(mn))
	column := make([]bool, len(mn[0]))
	for i, v1 := range mn {
		for j, v2 := range v1 {
			if v2 == 0 {
				row[i] = true
				column[j] = true
			}
		}
	}

	for i := 0; i < len(mn); i++ {
		for j := 0; j < len(mn[0]); j++ {
			if row[i] || column[j] {
				mn[i][j] = 0
				continue
			}
		}
	}
	return mn
}
