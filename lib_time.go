package easygolang

import (
	"fmt"
	"time"
)

type Time time.Time

func TimeZero() Time {
	return Time(time.Time{})
}

func TimeNow() Time {
	return Time(time.Now())
}

func TimeSeconds(start Time) float64 {
	elapsed := time.Since(time.Time(start))
	return float64(elapsed / time.Second)
}

func TimeSecondsSub(start Time, finished Time) float64 {
	f := time.Time(finished)
	delta := f.Sub(time.Time(start))
	return delta.Seconds()
}

func TimeStr(t Time, seconds bool) string {
	t2 := time.Time(t)
	s := ""
	if seconds {
		s = fmt.Sprintf("%d.%02d.%02d %02d:%02d:%02d",
			t2.Year(), t2.Month(), t2.Day(),
			t2.Hour(), t2.Minute(), t2.Second()) //, t.Second())
	} else {
		s = fmt.Sprintf("%d.%02d.%02d %02d:%02d",
			t2.Year(), t2.Month(), t2.Day(),
			t2.Hour(), t2.Minute()) //, t.Second())
	}
	return s
}

func TimeNowStr() string {
	return TimeStr(TimeNow(), true)
}

func TimeAddMS(t Time, ms int) Time {
	t2 := time.Time(t)
	return Time(t2.Add(time.Millisecond * time.Duration(ms)))
}

func TimeAddDays(t Time, days int) Time {
	t2 := time.Time(t)
	return Time(t2.Add(time.Hour * time.Duration(days*24)))
}

func TimeWeekday(t Time) int {
	t2 := time.Time(t)
	day := t2.Weekday()
	switch day {
	case time.Monday:
		return 1
	case time.Tuesday:
		return 2
	case time.Wednesday:
		return 3
	case time.Thursday:
		return 4
	case time.Friday:
		return 5
	case time.Saturday:
		return 6
	case time.Sunday:
		return 7
	default:
		return 0
	}
}

func TimeGolang(t Time) time.Time {
	return time.Time(t)
}
