package easygolang

import (
	"fmt"
	"time"
)

type Time time.Time

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

func TimeStr(t Time) string {
	t2 := time.Time(t)
	s := fmt.Sprintf("%d.%02d.%02d %02d:%02d",
		t2.Year(), t2.Month(), t2.Day(),
		t2.Hour(), t2.Minute()) //, t.Second())
	return s
}

func TimeNowStr() string {
	return TimeStr(TimeNow())
}

func TimeAddMS(t Time, ms int) Time {
	t2 := time.Time(t)
	return Time(t2.Add(time.Millisecond * time.Duration(ms)))
}
