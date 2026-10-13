//go:build !linux && !windows

package platform

import "time"

type HighResTimer struct {
    last time.Time
}

func NewTimer() *HighResTimer {
    return &HighResTimer{last: time.Now()}
}

func (t *HighResTimer) Now() time.Time {
    return time.Now()
}

func (t *HighResTimer) DeltaSince(prev time.Time) uint64 {
    return uint64(time.Since(prev).Nanoseconds())
}

func (t *HighResTimer) Resolution() string {
    return "nanoseconds (platform dependent)"
}
