package utils

import "time"

func MeasureTime[T any](f func() (T, error)) (T, time.Duration, error) {
	start := time.Now()
	result, err := f()
	return result, time.Since(start), err
}
