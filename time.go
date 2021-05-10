package any

import (
	"math/rand"
	"time"
)

// Time generates a random time.
func Time() time.Time {
	now := time.Now().UTC()
	dif := time.Duration(rand.Int())
	start := now.Add(dif * -1)
	end := now.Add(dif)
	return TimeSpan(start, end)
}

// TimeSpan generates a random time between two times.
func TimeSpan(start, end time.Time) time.Time {
	if start.Equal(end) {
		return start
	}
	sy, sm, sd := start.Date()
	sh := start.Hour()
	ss := start.Second()
	sn := start.Nanosecond()

	ey, em, ed := end.Date()
	eh := end.Hour()
	es := end.Second()
	en := end.Nanosecond()

	year := Between(sy, ey)
	month := Between(int(sm), int(em))
	day := Between(sd, ed)
	hour := Between(sh, eh)
	second := Between(ss, es)
	nano := Between(sn, en)

	t := start.AddDate(year-1, month-1, day-1)
	t = t.Add(time.Hour * time.Duration(hour))
	t = t.Add(time.Second * time.Duration(second))
	t = t.Add(time.Duration(nano))
	return t
}
