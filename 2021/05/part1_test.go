package main

import "testing"

func TestPart1(t *testing.T) {
	want := 6007
	got := Part1()
	if got != want {
		t.Errorf("Part1() = %d, want = %d", got, want)
	}
}
