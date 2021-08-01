package helper

import "time"

const DefaultTimeZone string = "Asia/Jakarta"

func GetCurrentTimeUTC() time.Time {
	return time.Now().In(time.UTC)
}


func GenerateCurrentTimeZone(timezone string) time.Time {
	loc, _ := time.LoadLocation(timezone)
	return time.Now().In(loc).Add(7 * time.Hour)
}

func GetCurrentTimeAsiaJakarta() time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(loc).Add(7 * time.Hour)
}

func GenerateDateFormatReturn(inputedTime *time.Time) string {

	stringResult := ""
	if inputedTime != nil {
		stringResult = inputedTime.Format("2006-01-02")
	}

	return stringResult
}
