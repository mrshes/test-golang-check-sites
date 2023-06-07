package pkg

import "time"

// Stopwatch Секундомер
func Stopwatch(f func()) (timeDiffMilli int64) {
	timeStart := time.Now().UnixMilli()
	f()
	timeEnd := time.Now().UnixMilli()
	return timeEnd - timeStart
}
