package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
)

// BallClock indicates which balls are in which time tracks
type BallClock struct {
	MinuteTrack []int `json:"Min"`
	FiveTrack   []int `json:"FiveMin"`
	HourTrack   []int `json:"Hour"`
	MainTrack   []int `json:"Main"`
}

const minuteCap = 4
const fiveCap = 11
const hourCap = 11

func newBallClock(numOfBalls int) BallClock {
	MainTrack := []int{}

	for i := 1; i <= numOfBalls; i++ {
		MainTrack = append(MainTrack, i)
	}

	return BallClock{
		[]int{},
		[]int{},
		[]int{},
		MainTrack,
	}
}

func main() {
	// probably a naive decision, but I'm just going to assume they've at least provided a number of balls
	numOfBalls, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic("Sorry, I couldn't understand how many balls you want in the clock")
	}

	clock := newBallClock(numOfBalls)

	if len(os.Args) == 2 {
		fmt.Printf("%s balls cycle after %d days", os.Args[1], clock.runUntilRepeat()/(24*60))
	} else {
		minutes, err := strconv.Atoi(os.Args[2])

		if err != nil {
			panic("Sorry, I couldn't understand how many minutes you want to run the clock")
		}

		clock.runForMinutes(minutes)

		enc := json.NewEncoder(os.Stdout)
		enc.Encode(clock)
	}
}

func (c *BallClock) runUntilRepeat() int {
	init := newBallClock(len(c.MainTrack))
	minutes := 1
	runClockOneMinute(c)

	// run until the configuration is back to the start, OR to prevent "minutes" from overflowing, until minutes would be the max int32 value
	for !reflect.DeepEqual(init, *c) && minutes < math.MaxInt32 {
		minutes++
		runClockOneMinute(c)
	}

	return minutes
}

func (c *BallClock) runForMinutes(minutes int) {
	for i := 0; i < minutes; i++ {
		runClockOneMinute(c)
	}
}

func runClockOneMinute(c *BallClock) {
	ball := c.MainTrack[0]
	c.MainTrack = c.MainTrack[1:]

	if len(c.MinuteTrack) == minuteCap {

		clearMinuteTrack(c)

		if len(c.FiveTrack) == fiveCap {

			clearFiveTrack(c)

			if len(c.HourTrack) == hourCap {

				clearHourTrack(c)

				//move the last ball back to the main track
				c.MainTrack = append(c.MainTrack, ball)
			} else {
				//add this ball to the hour track
				c.HourTrack = append(c.HourTrack, ball)
			}

		} else {
			// add this ball to the five track
			c.FiveTrack = append(c.FiveTrack, ball)
		}
	} else {
		// add this ball to the minute track
		c.MinuteTrack = append(c.MinuteTrack, ball)
	}
}

func clearMinuteTrack(c *BallClock) {
	for i := minuteCap - 1; i >= 0; i-- {
		minBall := c.MinuteTrack[i]
		c.MinuteTrack = c.MinuteTrack[:i]
		c.MainTrack = append(c.MainTrack, minBall)
	}
}

func clearFiveTrack(c *BallClock) {
	for i := fiveCap - 1; i >= 0; i-- {
		fiveBall := c.FiveTrack[i]
		c.FiveTrack = c.FiveTrack[:i]
		c.MainTrack = append(c.MainTrack, fiveBall)
	}
}

func clearHourTrack(c *BallClock) {
	for i := hourCap - 1; i >= 0; i-- {
		hourBall := c.HourTrack[i]
		c.HourTrack = c.HourTrack[:i]
		c.MainTrack = append(c.MainTrack, hourBall)
	}
}
