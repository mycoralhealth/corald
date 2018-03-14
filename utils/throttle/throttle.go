package throttle

import "time"

type Throttle struct {
	maxBumps  uint64
	period    time.Duration
	users     map[string]uint64
	lastReset time.Time
}

func NewThrottle(maxBumps uint64, period time.Duration) *Throttle {
	return &Throttle{
		maxBumps: maxBumps,
		period:   period,
	}
}

// Bump asks the throttler whether the given user can make a request.
// If they may, it returns true and increments the counter for that user.
// If they may not, it returns false
func (th *Throttle) Bump(name string) bool {
	if time.Now().Sub(th.lastReset) > th.period {
		// Reset the counters
		th.users = make(map[string]uint64)
	}
	if th.users[name] < th.maxBumps {
		th.users[name]++
		return true
	}

	return false
}
