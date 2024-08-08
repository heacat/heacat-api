package utils

import "time"

func SetInterval(duration time.Duration, fn func() (any, string)) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for range ticker.C {
		fn()
	}
}
