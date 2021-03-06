package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSingleLinkedListTrim(t *testing.T) {
	tests := map[string]struct {
		list []int
		want []int
	}{
		"0": {
			list: []int{10},
			want: nil,
		},
		"1": {
			list: []int{10, 11},
			want: []int{10},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// var l singleLinkedList
			l := newSingleLinkedList(tt.list...)
			l.trim()

			l.print()
			// fmt.Printf("after trim: %+v %t\n", l, (*l == singleLinkedList{}))
			if !reflect.DeepEqual(l.getValues(), tt.want) {
				t.Errorf("got: %v, want: %v", l.getValues(), tt.want)
			}
		})
	}
}

func TestSingleLinkedListDelete(t *testing.T) {
	tests := map[string]struct {
		list   []int
		target int
		want   []int
	}{
		"0": {
			list:   []int{10},
			target: 10,
			want:   nil,
		},
		"1": {
			list:   []int{11},
			target: 10,
			want:   []int{11},
		},
		"2": {
			list:   []int{10, 11},
			target: 10,
			want:   []int{11},
		},
		"3": {
			list:   []int{11, 10},
			target: 10,
			want:   []int{11},
		},
		"4": {
			list:   []int{11, 10, 10},
			target: 10,
			want:   []int{11},
		},
		"5": {
			list:   []int{10, 11, 10},
			target: 10,
			want:   []int{11},
		},
		"6": {
			list:   []int{10, 10, 11},
			target: 10,
			want:   []int{11},
		},
		"7": {
			list:   []int{10, 10, 10},
			target: 10,
			want:   nil,
		},
		"8": {
			list:   []int{10, 11, 12},
			target: 10,
			want:   []int{11, 12},
		},
		"9": {
			list:   []int{10, 11, 12, 13},
			target: 9,
			want:   []int{10, 11, 12, 13},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			l := newSingleLinkedList(tt.list...)
			l.delete(tt.target)

			l.print()
			if !reflect.DeepEqual(l.getValues(), tt.want) {
				t.Errorf("got: %v, want: %v", l.getValues(), tt.want)
			}
		})
		// t.Run(name, func(t *testing.T) {
		// 	l := newSingleLinkedList(tt.list...)
		// 	l.delete2(tt.target)
		// 	l.print()
		// 	if !reflect.DeepEqual(l.getValues(), tt.want) {
		// 		t.Errorf("got: %v, want: %v", l.getValues(), tt.want)
		// 	}
		// })
	}
}

func TestSingleLinkedListDeleteDup(t *testing.T) {
	tests := map[string]struct {
		list []int
		want []int
	}{
		"0": {
			list: []int{10},
			want: []int{10},
		},
		"1": {
			list: []int{10, 10},
			want: []int{10},
		},
		"2": {
			list: []int{10, 11},
			want: []int{10, 11},
		},
		"3": {
			list: []int{10, 11, 12, 13},
			want: []int{10, 11, 12, 13},
		},
		"4": {
			list: []int{10, 11, 12, 13, 11},
			want: []int{10, 11, 12, 13},
		},
		"5": {
			list: []int{10, 10, 11, 12, 13},
			want: []int{10, 11, 12, 13},
		},
		"6": {
			list: []int{10, 10, 10, 10},
			want: []int{10},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			l := newSingleLinkedList(tt.list...)
			l.print()

			l.deleteDuplicated()
			if !reflect.DeepEqual(l.getValues(), tt.want) {
				t.Errorf("got: %v, want: %v", l.getValues(), tt.want)
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

// // 上のdelete関数の再帰をつかわない書き直し
// func (s *singleLinkedList) delete2(target int) {

// 	current := s
// 	previous := &singleLinkedList{next: current}
// 	for {
// 		fmt.Println("this", current.getValues(), previous.getValues())
// 		if current.isLast() {
// 			break
// 		}
// 		// if current.value == target {
// 		// 	previous.next = current.next
// 		// }
// 		previous = current
// 		current = current.next

// 		*s.next = current
// 	}
// 	// *s = *previous.next
// }

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

// func (s *singleLinkedList) isBeforeLast() bool {
// 	if s.next == nil {
// 		return false
// 	}
// 	if s.next.next == nil {
// 		return true
// 	}
// 	return false
// }
