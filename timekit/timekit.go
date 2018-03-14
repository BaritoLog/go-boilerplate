package timekit

import "time"

func Sleep(durationFormat string) {
	sleepDuration, _ := time.ParseDuration(durationFormat)
	time.Sleep(sleepDuration)
}
