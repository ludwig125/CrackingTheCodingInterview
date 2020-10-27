package main

import (
	"fmt"
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
		// "2": {
		// 	mn: [][]int{
		// 		{1, 2, 3},
		// 		{4, 5, 6},
		// 		{7, 8, 9},
		// 	},
		// 	want: [][]int{
		// 		{7, 4, 1},
		// 		{8, 5, 2},
		// 		{9, 6, 3},
		// 	},
		// },
		// "3": {
		// 	mn: [][]int{
		// 		{1, 2, 3, 4},
		// 		{5, 6, 7, 8},
		// 		{9, 10, 11, 12},
		// 		{13, 14, 15, 16},
		// 	},
		// 	want: [][]int{
		// 		{13, 9, 5, 1},
		// 		{14, 10, 6, 2},
		// 		{15, 11, 7, 3},
		// 		{16, 12, 8, 4},
		// 	},
		// },
		// "5": {
		// 	mn: [][]int{
		// 		{1, 2, 3, 4, 5, 6},
		// 		{7, 8, 9, 10, 11, 12},
		// 		{13, 14, 15, 16, 17, 18},
		// 		{19, 20, 21, 22, 23, 24},
		// 		{25, 26, 27, 28, 29, 30},
		// 		{31, 32, 33, 34, 35, 36},
		// 	},
		// 	want: [][]int{
		// 		{31, 25, 19, 13, 7, 1},
		// 		{32, 26, 20, 14, 8, 2},
		// 		{33, 27, 21, 15, 9, 3},
		// 		{34, 28, 22, 16, 10, 4},
		// 		{35, 29, 23, 17, 11, 5},
		// 		{36, 30, 24, 18, 12, 6},
		// 	},
		// },
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// fmt.Println(tt.mn)
			printMatrix(tt.mn)
			// printMatrix(ninetyDegreeRotation(tt.mn))
			mn := padding(tt.mn)
			printMatrix(mn)
			// if !reflect.DeepEqual(tt.mn, tt.want) {
			// 	t.Errorf("got: %v, want: %v", tt.mn, tt.want)
			// }
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
				// fmt.Println(i, j, mn[i][j])
			}
		}
	}
	fmt.Println(row)
	fmt.Println(column)

	// var nmn [][]int // new m*n
	// for i := 0; i < len(mn); i++ {
	// 	for j := 0; j < len(mn[0]); j++ {

	// 	}
	// }
	return mn
}
