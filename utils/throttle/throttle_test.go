package throttle_test

import (
	"testing"
	"time"

	"github.com/mycoralhealth/corald/utils/throttle"
)

var throttleTests = []struct {
	user      string
	expected  bool
	postDelay uint64 // Millisecond
}{
	{"a", true, 0},
	{"b", true, 0},
	{"b", true, 0},
	{"a", true, 0},
	{"a", false, 0},
	{"b", false, 10},

	{"a", true, 0},
	{"b", true, 0},
}

func TestThrottle(t *testing.T) {
	th := throttle.NewThrottle(2, 10*time.Millisecond)
	for i, tt := range throttleTests {
		actual := th.Bump(tt.user)
		if tt.expected != actual {
			t.Errorf("%d:Bump(%s) = %v, want %v", i, tt.user, actual, tt.expected)
		}
		time.Sleep(time.Duration(tt.postDelay) * time.Millisecond)

	}

}
