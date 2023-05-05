package feature

import (
	"time"
	"errors"
)

func GetDay(dateString string) (string, error) {
	dateFormats := []string{"02/01/2006", "01-02-2006", "2006-01-02", "2 January 2006", "2006/01/02", "02/1/2006", "2/01/2006", "2 Jan 2006", "1-02-2006", "01-2-2006", "1-2-2006", "1/2/2006"}
	var date time.Time
	var err error
	for _, format := range dateFormats{
		date, err = time.Parse(format, dateString)
		if(err == nil){
			break
		}
	}
	if(err != nil){
		return "", errors.New("Invalid Date Format");
	}
	day := date.Weekday().String()
	return day, nil
}