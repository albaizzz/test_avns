package helpers

import (
	"time"

	config "github.com/spf13/viper"
)

// Layout for time
const (
	DateTimeLayout = "2006-01-02 15:04:05"
)

func timeLoadLocation() *time.Location {
	timezone := config.GetString("app.timezone")
	location, _ := time.LoadLocation(timezone)
	return location
}

// GetCurrentDateTime gets current date and time
func GetCurrentDateTime() time.Time {
	currentDateTime := time.Now().In(timeLoadLocation())
	return currentDateTime
}

// TimeToString convert from time.Time into string
func TimeToString(dateInput time.Time) string {
	return dateInput.Format(DateTimeLayout)
}

// StringToTime - convert from date string into time.Time
func StringToTime(dateInput string) (time.Time, error) {
	result, err := time.Parse(DateTimeLayout, dateInput)
	return result, err
}

// TimestampToDateTime convert timestamp to date time
func TimestampToDateTime(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(DateTimeLayout)
}
