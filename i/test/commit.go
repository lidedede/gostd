package test

import "time"

func CheckCommit(ts []time.Time, checkTime time.Time) bool {
	return len(getLastTwoDayCommit(ts, checkTime)) > 0
}

func getLastTwoDayCommit(ts []time.Time, checkTime time.Time) []time.Time {
	startTime := getLastTwoDayStartTime(checkTime)
	var twoDayCommits []time.Time
	for _, t := range ts {
		if t.After(startTime) {
			//if t.After(startTime) && isWorkDay(t) {
			twoDayCommits = append(twoDayCommits, t)
		}
	}
	return twoDayCommits
}

func isWorkDay(t time.Time) bool {
	if t.Weekday() == time.Sunday ||
		t.Weekday() == time.Saturday {
		return false
	}
	return true
}

func getLastTwoDayStartTime(cur time.Time) time.Time {
	zeroTime := time.Date(cur.Year(), cur.Month(), cur.Day(), 0, 0, 0, 0, time.Local)
	num := 2
	for num > 0 {
		zeroTime = zeroTime.Add(-time.Hour*24)
		if !isWorkDay(zeroTime) {
			continue
		}
		num--
	}
	return zeroTime
}