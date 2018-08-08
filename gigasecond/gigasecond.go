package gigasecond

import "time"

// AddGigasecond its return add a Gigasecond.
func AddGigasecond(t time.Time) time.Time {
	return t.Local().Add(time.Second * time.Duration(1000000000))
}
