package main

import (
	"strconv"
	"testing"
)

func Test_findCh(t *testing.T) {
	tests := []struct {
		lenRes int
		want   int
	}{
		{lenRes: 10, want: 5},
		{lenRes: 11, want: 6},
		{lenRes: 0, want: 0},
	}
	for ttI, tt := range tests {
		t.Run(strconv.Itoa(ttI), func(t *testing.T) {
			if got := findCh(tt.lenRes); got != tt.want {
				t.Errorf("findCh() = %v, want %v", got, tt.want)
			}
		})
	}
}
