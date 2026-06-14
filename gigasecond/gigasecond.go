package gigasecond

import (
	"math"
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(math.Pow10(9)) * time.Second)
}
