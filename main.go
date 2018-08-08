package main

// BallClock indicates which balls are in which time tracks
type BallClock struct {
	minuteTrack []int
	fiveTrack   []int
	hourTrack   []int
	mainTrack   []int
}

const minuteCap = 4
const fiveCap = 11
const hourCap = 11

func newBallClock(numOfBalls int) BallClock {
	mainTrack := []int{}

	for i := 1; i <= numOfBalls; i++ {
		mainTrack = append(mainTrack, i)
	}

	return BallClock{
		[]int{},
		[]int{},
		[]int{},
		mainTrack,
	}
}

func main() {

}

func (c *BallClock) runForMinutes(minutes int) BallClock {
	for i := 0; i < minutes; i++ {
		ball := c.mainTrack[0]
		c.mainTrack = c.mainTrack[1:]

		if len(c.minuteTrack) == minuteCap {
			c.fiveTrack = append(c.fiveTrack, ball)
			for i := minuteCap - 1; i >= 0; i-- {
				minBall := c.minuteTrack[i]
				c.minuteTrack = c.minuteTrack[:i]
				c.mainTrack = append(c.mainTrack, minBall)
			}
		} else {
			c.minuteTrack = append(c.minuteTrack, ball)
		}
	}

	return *c
}
