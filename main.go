package main

// BallClock knows how many balls it has to use, and maintains its time tracks
type BallClock struct {
	numOfBalls  int
	minuteTrack [4]int
	fiveTrack   [11]int
	hourTrack   [11]int
	mainTrack   []int
}

func newBallClock(numOfBalls int) BallClock {
	mainTrack := []int{}

	for i := 1; i <= numOfBalls; i++ {
		mainTrack = append(mainTrack, i)
	}

	return BallClock{
		numOfBalls,
		[4]int{},
		[11]int{},
		[11]int{},
		mainTrack,
	}
}

func main() {

}
