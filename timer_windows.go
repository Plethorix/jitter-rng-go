//go:build windows

package platform

/*
En Windows, time.Now() usa GetSystemTimeAsFileTime (resolución ~15.6ms por
defecto) o QueryPerformanceCounter (QPC, resolución sub-microsegundo) vía
runtime.

Go en Windows ya usa QPC internamente para time.Now() cuando es posible.
El TSC es estable en CPUs modernas (invariant TSC flag).
*/

import (
    "runtime"
    "time"
)

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
    // En Windows, Go suele usar QueryPerformanceCounter
    // El TSC es accesible y estable en x86_64 modernos
    if runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64" {
        return "nanoseconds (QPC / invariant TSC)"
    }
    return "nanoseconds (QPC)"
}
