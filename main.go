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

func (c *BallClock) runUntilRepeat() int {
	return 0
}

func runClockOneMinute(c *BallClock) {
	ball := c.mainTrack[0]
	c.mainTrack = c.mainTrack[1:]

	if len(c.minuteTrack) == minuteCap {
		// move the 4 balls on the minute track back to the main track
		for i := minuteCap - 1; i >= 0; i-- {
			minBall := c.minuteTrack[i]
			c.minuteTrack = c.minuteTrack[:i]
			c.mainTrack = append(c.mainTrack, minBall)
		}

		if len(c.fiveTrack) == fiveCap {
			// move the 11 balls on the five track back to the main track
			for i := fiveCap - 1; i >= 0; i-- {
				fiveBall := c.fiveTrack[i]
				c.fiveTrack = c.fiveTrack[:i]
				c.mainTrack = append(c.mainTrack, fiveBall)
			}

			if len(c.hourTrack) == hourCap {
				// move the 11 balls on the hour track back to the main track
				for i := hourCap - 1; i >= 0; i-- {
					hourBall := c.hourTrack[i]
					c.hourTrack = c.hourTrack[:i]
					c.mainTrack = append(c.mainTrack, hourBall)
				}

				//and move the ball to the main track
				c.mainTrack = append(c.mainTrack, ball)
			} else {
				//add this ball to the hour track
				c.hourTrack = append(c.hourTrack, ball)
			}

		} else {
			// add this ball to the five track
			c.fiveTrack = append(c.fiveTrack, ball)
		}
	} else {
		// add this ball to the minute track
		c.minuteTrack = append(c.minuteTrack, ball)
	}
}

func (c *BallClock) runForMinutes(minutes int) {
	for i := 0; i < minutes; i++ {
		runClockOneMinute(c)
	}
}
