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

func TestRunForMinutes30Balls(t *testing.T) {

	t.Run("it moves 4 balls onto the minute track", func(t *testing.T) {
		clock := newBallClock(30)
		clock.runForMinutes(4)
		want := BallClock{
			[]int{1, 2, 3, 4},
			[]int{},
			[]int{},
			[]int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		}

		if !reflect.DeepEqual(clock, want) {
			t.Errorf("running the clock should have made %v, but instead made %v", want, clock)
		}

	})

	t.Run("it moves the fifth ball to the fiveTrack, and returns others to the mainTrack", func(t *testing.T) {
		clock := newBallClock(30)
		clock.runForMinutes(5)
		want := BallClock{
			[]int{},
			[]int{5},
			[]int{},
			[]int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 4, 3, 2, 1},
		}

		if !reflect.DeepEqual(clock, want) {
			t.Errorf("running the clock should have made %v, but instead made %v", want, clock)
		}
	})

	t.Run("it is correct up to 30 minutes from start", func(t *testing.T) {
		clock := newBallClock(30)
		clock.runForMinutes(30)
		want := BallClock{
			[]int{},
			[]int{5, 10, 15, 20, 25, 30},
			[]int{},
			[]int{4, 3, 2, 1, 9, 8, 7, 6, 14, 13, 12, 11, 19, 18, 17, 16, 24, 23, 22, 21, 29, 28, 27, 26},
		}

		if !reflect.DeepEqual(clock, want) {
			t.Errorf("running the clock should have made %v, but instead made %v", want, clock)
		}
	})

	t.Run("it holds 11 balls in the fiveTrack and 4 balls in the minuteTrack", func(t *testing.T) {
		clock := newBallClock(30)
		clock.runForMinutes(59)
		want := BallClock{
			[]int{2, 3, 4, 14},
			[]int{5, 10, 15, 20, 25, 30, 9, 13, 17, 21, 1},
			[]int{},
			[]int{6, 7, 8, 18, 19, 11, 12, 22, 23, 24, 16, 26, 27, 28, 29},
		}

		if !reflect.DeepEqual(clock, want) {
			t.Errorf("running the clock should have made %v, but instead made %v", want, clock)
		}
	})

	t.Run("it moves a ball to the hourTrack after 60 minutes", func(t *testing.T) {
		clock := newBallClock(30)
		clock.runForMinutes(60)
		want := BallClock{
			[]int{},
			[]int{},
			[]int{6},
			[]int{7, 8, 18, 19, 11, 12, 22, 23, 24, 16, 26, 27, 28, 29, 14, 4, 3, 2, 1, 21, 17, 13, 9, 30, 25, 20, 15, 10, 5},
		}

		if !reflect.DeepEqual(clock, want) {
			t.Errorf("running the clock should have made %v, but instead made %v", want, clock)
		}
	})

	t.Run("it holds 11 balls in the hourTrack, 11 in the fiveTrack and 4 in the minuteTrack", func(t *testing.T) {
		clock := newBallClock(30)
		clock.runForMinutes(719)
		want := BallClock{
			[]int{29, 21, 7, 27},
			[]int{11, 10, 20, 1, 8, 9, 3, 24, 16, 14, 30},
			[]int{6, 12, 17, 4, 15, 23, 22, 2, 26, 18, 28},
			[]int{25, 5, 19, 13},
		}

		if !reflect.DeepEqual(clock, want) {
			t.Errorf("running the clock should have made %v, but instead made %v", want, clock)
		}
	})

	t.Run("it returns all balls to the mainTrack after 12 hours", func(t *testing.T) {
		clock := newBallClock(30)
		clock.runForMinutes(720)
		want := BallClock{
			[]int{},
			[]int{},
			[]int{},
			[]int{5, 19, 13, 27, 7, 21, 29, 30, 14, 16, 24, 3, 9, 8, 1, 20, 10, 11, 28, 18, 26, 2, 22, 23, 15, 4, 17, 12, 6, 25},
		}

		if !reflect.DeepEqual(clock, want) {
			t.Errorf("running the clock should have made %v, but instead made %v", want, clock)
		}
	})

	t.Run("it should match given test criteria after 325 minutes", func(t *testing.T) {
		clock := newBallClock(30)
		clock.runForMinutes(325)
		want := BallClock{
			[]int{},
			[]int{22, 13, 25, 3, 7},
			[]int{6, 12, 17, 4, 15},
			[]int{11, 5, 26, 18, 2, 30, 19, 8, 24, 10, 29, 20, 16, 21, 28, 1, 23, 14, 27, 9},
		}

		if !reflect.DeepEqual(clock, want) {
			t.Errorf("running the clock should have made %v, but instead made %v", want, clock)
		}
	})

}
