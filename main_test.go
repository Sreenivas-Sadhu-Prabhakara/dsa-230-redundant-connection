package main

import "testing"

func TestFindRedundantConnection(t *testing.T) {
	cases := []struct {
		edges [][2]int
		want  [2]int
	}{
		{[][2]int{{1, 2}, {1, 3}, {2, 3}}, [2]int{2, 3}},
		{[][2]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}, [2]int{1, 4}},
		{[][2]int{{1, 2}, {2, 1}}, [2]int{2, 1}}, // parallel edge is the cycle
		{[][2]int{{3, 4}, {1, 2}, {2, 4}, {3, 1}}, [2]int{3, 1}},
	}
	for _, c := range cases {
		if got := FindRedundantConnection(c.edges); got != c.want {
			t.Errorf("FindRedundantConnection(%v) = %v, want %v", c.edges, got, c.want)
		}
	}
}
