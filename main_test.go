package main

import "testing"
import "reflect"

func TestNewBallClock(t *testing.T) {
	got := newBallClock(30)
	want := BallClock{
		[4]int{},
		[11]int{},
		[11]int{},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("creating a new ball clock should have made %v, but instead made %v", want, got)
	}
}

func TestRunTimeFor30Balls(t *testing.T) {
	clock := newBallClock(30)

	got := clock.runForMinutes(1)
	want := BallClock{
		[4]int{1},
		[11]int{},
		[11]int{},
		[]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("running the clock should have made %v, but instead made %v", want, got)
	}
}
