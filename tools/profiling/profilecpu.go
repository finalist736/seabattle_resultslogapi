package profiling

import (
	"fmt"
	"time"

	"github.com/pkg/profile"
)

var cpuProfiler *profile.Profile

func ProfileCPU() {
	now := time.Now()
	profilePath := fmt.Sprintf("./profiles_%d_%d_%d_%d_%d", // store profiles in current directory
		now.Day(), now.Month(),
		now.Hour(), now.Minute(), now.Second())
	cpuProfiler = (profile.Start(profile.CPUProfile, profile.ProfilePath(profilePath), profile.NoShutdownHook)).(*profile.Profile)
}

func CloseCPU() {
	if cpuProfiler == nil {
		return
	}
	cpuProfiler.Stop()
}
