package util

import (
	"strconv"
	"time"
)

func Now() int {
	return int(time.Now().Unix())
}
func NowStr() string {
	return strconv.Itoa(int(time.Now().Unix()))
}
func ThisDay() (time.Time, time.Time) {
	now := time.Now()
	currentYear, currentMonth, currentDay := now.Date()
	st := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, now.Location())
	et := st.AddDate(0, 0, 1)
	return st, et
}
func ThisMonth(month int) (time.Time, time.Time) {
	now := time.Now()
	currentYear, _, _ := now.Date()
	firstOfMonth := time.Date(currentYear, time.Month(month), 1, 0, 0, 0, 0, now.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, 0)
	return firstOfMonth, lastOfMonth
}

func ParseTimeLocal(format, date string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	return time.ParseInLocation(format, date, loc)
}

//format is 2006-01-02
func ParseDate(date string) (time.Time, error) {
	return ParseTimeLocal("2006-01-02", date)
}

//format is 2006-01-02 15:04
func ParseDateTimeShort(date string) (time.Time, error) {
	return ParseTimeLocal("2006-01-02 15:04", date)
}

func FormatTimeWith(unix int, format string) string {
	return time.Unix(int64(unix), 0).Format(format)
}

func FormatDatetime(unix int) string {
	return FormatTimeWith(unix, "2006-01-02 15:04:05")
}
func FormatDate(unix int) string {
	return FormatTimeWith(unix, "2006-01-02")
}

func FormatTime(unix int) string {
	return FormatTimeWith(unix, "15:04:05")
}
func FormatDatetimeShort(unix int) string {
	return FormatTimeWith(unix, "2006-01-02 15:04")
}
