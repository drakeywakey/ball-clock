package main

// BallClock indicates which balls are in which time tracks
type BallClock struct {
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
		[4]int{},
		[11]int{},
		[11]int{},
		mainTrack,
	}
}

func main() {

}

func (c *BallClock) runForMinutes(minutes int) BallClock {
	return BallClock{}
}
