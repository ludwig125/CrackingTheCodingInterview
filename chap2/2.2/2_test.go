package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSingleLinkedListGetKFromTail(t *testing.T) {
	tests := map[string]struct {
		list []int
		k    int
		want *singleLinkedList
	}{
		"0": {
			list: []int{10},
			k:    1,
			want: newSingleLinkedList([]int{10}...),
		},
		"1": {
			list: []int{10, 11},
			k:    1,
			want: newSingleLinkedList([]int{11}...),
		},
		"2": {
			list: []int{10, 11, 12},
			k:    1,
			want: newSingleLinkedList([]int{12}...),
		},
		"3": {
			list: []int{10, 11, 12},
			k:    2,
			want: newSingleLinkedList([]int{11, 12}...),
		},
		"4": {
			list: []int{10, 11, 12},
			k:    3,
			want: newSingleLinkedList([]int{10, 11, 12}...),
		},
		"5": {
			list: []int{10, 11, 12},
			k:    4,
			want: nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			l := newSingleLinkedList(tt.list...)
			l.print()

			got := l.getKFromTail(tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
			got2 := l.getKFromTail2(tt.k)
			if !reflect.DeepEqual(got2, tt.want) {
				t.Errorf("got: %v, want: %v", got2, tt.want)
			}
			got3 := l.getKFromTail3(tt.k)
			if !reflect.DeepEqual(got3, tt.want) {
				t.Errorf("got: %v, want: %v", got3, tt.want)
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

func (s *singleLinkedList) getKFromTail(k int) *singleLinkedList {
	var list []*singleLinkedList
	list = s.addList(list)
	// fmt.Println("!this", list)
	if len(list) < k {
		// fmt.Println("!this1")
		return nil
	}
	// fmt.Println("!this2")
	return list[len(list)-k]
}

func (s *singleLinkedList) addList(list []*singleLinkedList) []*singleLinkedList {
	// fmt.Println("this", s)
	list = append(list, s)
	// fmt.Println("this1", list[0])
	if s.isLast() {
		// fmt.Println("this2", list[0])
		return list
	}
	// fmt.Println("thi3")

	return s.next.addList(list)
}

// 再帰を使わないで書いてみたもの
func (s *singleLinkedList) getKFromTail2(k int) *singleLinkedList {
	var list []*singleLinkedList
	// list := make([]*singleLinkedList, k)

	thisS := s
	for {
		list = append(list, thisS)
		// fmt.Println("s value", thisS.value, len(list))
		if thisS.isLast() {
			break
		}
		thisS = thisS.next
	}
	if len(list) < k {
		return nil
	}
	return list[len(list)-k]
}

// ランナーテクニックで上の関数を書き直したもの
func (s *singleLinkedList) getKFromTail3(k int) *singleLinkedList {
	s1 := s
	s2 := s
	// fmt.Println("begin s1 value s2 value", s1.value, s2.value)
	for i := 1; i < k; i++ {
		if s1.isLast() {
			// fmt.Println("last s1 value s2 value", s1.value, s2.value)
			return nil
		}
		s1 = s1.next
	}

	for {
		// fmt.Println("s1 value s2 value", s1.value, s2.value)
		if s1.isLast() {
			break
		}
		s1 = s1.next
		s2 = s2.next
	}

	return s2
}
