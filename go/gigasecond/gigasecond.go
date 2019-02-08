package gigasecond

import "time"

// AddGigasecond adjusts the given time by 1,000,000,000 seconds
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
