package main

import "testing"
import "reflect"

func TestNewBallClock(t *testing.T) {
	got := newBallClock(30)
	want := BallClock{
		[]int{},
		[]int{},
		[]int{},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("creating a new ball clock should have made %v, but instead made %v", want, got)
	}
}

func TestRunTimeFor30Balls(t *testing.T) {

	t.Run("it moves 4 balls onto the minute track", func(t *testing.T) {
		clock := newBallClock(30)
		got := clock.runForMinutes(4)
		want := BallClock{
			[]int{1, 2, 3, 4},
			[]int{},
			[]int{},
			[]int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("running the clock should have made %v, but instead made %v", want, got)
		}

	})

	t.Run("it moves the fifth ball to the fiveTrack, and returns others to the mainTrack", func(t *testing.T) {
		clock := newBallClock(30)
		got := clock.runForMinutes(5)
		want := BallClock{
			[]int{},
			[]int{5},
			[]int{},
			[]int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 4, 3, 2, 1},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("running the clock should have made %v, but instead made %v", want, got)
		}
	})
}
