package main

const (
	SecondsPerMinute = 60
	SecondsPerHour = SecondsPerMinute * 60
	SecondsPerDay = SecondsPerHour * 24
)

func resoveTime(seconds int) (day int, hour int, minute int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute
	return
}