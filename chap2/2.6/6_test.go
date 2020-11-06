package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSingleLinkedListPalindrome(t *testing.T) {
	tests := map[string]struct {
		list []int
		want bool
	}{
		"1": {
			list: []int{1, 2, 3},
			want: false,
		},
		"2": {
			list: []int{1, 2, 1},
			want: true,
		},
		"3": {
			list: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
			want: true,
		},
		"4": {
			list: []int{1, 2, 3, 4, 5, 6, 3, 2, 1},
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			l := newSingleLinkedList(tt.list...)
			l.print()
			got := l.isPalindrome()
			// // fmt.Println("value", sum1)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

type singleLinkedList struct {
	next  *singleLinkedList
	value int
}

func newSingleLinkedList(is ...int) *singleLinkedList {
	l := &singleLinkedList{value: is[0]}
	for i := 1; i < len(is); i++ {
		l.add(is[i])
	}
	return l
}

func (s *singleLinkedList) print() {
	if (*s == singleLinkedList{}) {
		return
	}
	fmt.Println(s.value)
	if s.next != nil {
		s.next.print()
	}
}

func (s *singleLinkedList) getValues() []int {
	if s == nil {
		return nil
	}
	// fmt.Println("s.value", s.value, )
	var l []int
	if (*s == singleLinkedList{}) {
		// fmt.Println("s.value", *s, singleLinkedList{})
		return l
	}
	return s.list(l)
}

func (s *singleLinkedList) list(l []int) []int {
	l = append(l, s.value)
	if s.next == nil {
		return l
	}
	return s.next.list(l)
}

func (s *singleLinkedList) add(target int) {
	if s.next == nil {
		s.next = &singleLinkedList{value: target}
		return
	}
	s.next.add(target)
}

func (s *singleLinkedList) search(target, cnt int) int {
	if s.value == target {
		return cnt
	}
	if s.next == nil {
		return -1
	}
	return s.next.search(target, cnt+1)
}

func (s *singleLinkedList) trim() {
	if s.isLast() {
		*s = singleLinkedList{}
		return
	}
	if s.next.isLast() {
		*s = singleLinkedList{value: s.value}
		return
	}
	s.next.trim()
}

func (s *singleLinkedList) delete(target int) {
	// del関数の中で、”次のNodeがTargetかどうか” という評価から始めて再帰を使いたいので、暫定的に適当なNodeを先頭にしている
	// これをすることで、次のNodeの評価から始められて、全部終わった後で先頭を消すという処理にした

	tmp := 0
	newS := &singleLinkedList{value: tmp, next: s}
	newS.del(target)
	// fmt.Println("!this", newS.getValues())
	if len(newS.getValues()) <= 1 {
		// fmt.Println("!this1", newS.getValues())
		*s = singleLinkedList{}
		return
	}
	// fmt.Println("!this2", newS.getValues())

	*s = singleLinkedList{value: newS.next.value, next: newS.next.next}
}
func (s *singleLinkedList) del(target int) {
	// fmt.Println("this", s.getValues())
	if s.isLast() {
		// fmt.Println("this-2", s.getValues())
		return
	}
	nexts := s.next // next Nodeを仮登録
	// 次のNodeがTargetの場合
	if s.next.value == target {
		// fmt.Println("this0", s.getValues())
		if s.next.isLast() {
			// fmt.Println("this1", s.getValues())
			// fmt.Println("this1-2", s.getValues(), s.value)
			*s = singleLinkedList{value: s.value}
			// fmt.Println("this1-3", s.getValues(), s.value)
			return
		}
		// fmt.Println("this2", s.getValues())
		// 次のNodeがTargetの時は”次の次のNode”を次のNodeとして設定する
		*s = singleLinkedList{value: s.value, next: s.next.next}

		nexts = s // 一つずらしたので、NextNodeを自分自身にする
	}
	// fmt.Println("this3", s.getValues())
	nexts.del(target)
}

func (s *singleLinkedList) deleteDuplicated() {
	// var m map[int]bool
	m := make(map[int]bool, len(s.getValues()))
	// m[10] = true

	s.deleteDup(m)
	fmt.Println("all after", s.getValues())
}

func (s *singleLinkedList) deleteDup(m map[int]bool) {
	m[s.value] = true
	if s.isLast() { // 最後のNodeなら終わり
		return
	}

	nextS := s.next // next Nodeを仮登録
	if _, ok := m[s.next.value]; ok {
		*s = singleLinkedList{value: s.value, next: s.next.next}
		nextS = s // 一つずらしたので、NextNodeを自分自身にする
	}
	nextS.deleteDup(m)
}

func (s *singleLinkedList) isLast() bool {
	if s.next != nil {
		return false
	}
	return true
}

// func (s *singleLinkedList) getKFromTail(k int) *singleLinkedList {
// 	var list []*singleLinkedList
// 	list = s.addList(list)
// 	// fmt.Println("!this", list)
// 	if len(list) < k {
// 		// fmt.Println("!this1")
// 		return nil
// 	}
// 	// fmt.Println("!this2")
// 	return list[len(list)-k]
// }

// func (s *singleLinkedList) addList(list []*singleLinkedList) []*singleLinkedList {
// 	// fmt.Println("this", s)
// 	list = append(list, s)
// 	// fmt.Println("this1", list[0])
// 	if s.isLast() {
// 		// fmt.Println("this2", list[0])
// 		return list
// 	}
// 	// fmt.Println("thi3")

// 	return s.next.addList(list)
// }

// func (s *singleLinkedList) summedUpList() int {
// 	return s.summedUpListInOrder(1, 0)
// }

// func (s *singleLinkedList) summedUpListInOrder(order, sum int) int {
// 	// fmt.Println("num", order*s.value)
// 	// sum += order * s.value
// 	if s.isLast() {
// 		// fmt.Println("order", order)
// 		return order*s.value + sum
// 	}
// 	return s.next.summedUpListInOrder(order*10, sum+order*s.value)
// }

// func parseSummedUpList(n int) *singleLinkedList {
// 	s := &singleLinkedList{}
// 	s.parseSummedUpListInOrder(n)
// 	// fmt.Println("value", s.getValues())
// 	return s
// }

// func (s *singleLinkedList) parseSummedUpListInOrder(n int) {
// 	if n < 10 {
// 		// fmt.Println("htis", n)
// 		*s = singleLinkedList{value: n % 10}
// 		return
// 	}
// 	// fmt.Println(n % 10)
// 	*s = singleLinkedList{value: n % 10, next: &singleLinkedList{}}
// 	// fmt.Println("v", s.getValues())
// 	s.next.parseSummedUpListInOrder(n / 10)
// }

func (s *singleLinkedList) isPalindrome() bool {

	fmt.Println("values", s.getValues())

	reversed := reverseSlice(s.getValues())
	fmt.Println("reverse values", reversed)

	return reflect.DeepEqual(s.getValues(), newSingleLinkedList(reversed...).getValues())
	// return true
}

func reverseSlice(l []int) []int {
	l2 := make([]int, 0, len(l))
	for i := len(l) - 1; i >= 0; i-- {
		l2 = append(l2, l[i])
	}
	return l2
}
