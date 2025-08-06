// Package gigasecond provides functionality to calculate times one gigasecond apart.
package gigasecond

import "time"

// AddGigasecond returns the moment one gigasecond (1×10^9 seconds) after t.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
