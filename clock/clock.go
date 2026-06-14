package clock

import (
	"fmt"
)

type Clock struct {
	Hour   int
	Minute int
}

const (
	MinutesPerDay  = 1440
	MinutesPerHour = 60
)

func New(hour, minute int) Clock {
	minutesInDay := (hour*MinutesPerHour + minute) % MinutesPerDay
	if minutesInDay < 0 {
		minutesInDay += MinutesPerDay
	}
	return Clock{Hour: minutesInDay / MinutesPerHour, Minute: minutesInDay % MinutesPerHour}
}

func (clock Clock) Add(minute int) Clock {
	return New(clock.Hour, clock.Minute+minute)
}

func (clock Clock) Subtract(minute int) Clock {
	return New(clock.Hour, clock.Minute-minute)
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.Hour, clock.Minute)
}
