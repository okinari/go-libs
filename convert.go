package golibs

import "time"

func ConvertDatetimeFromString(strDatetime string) time.Time {
	retTime, _ := time.Parse("2006/01/02 15:04:05", strDatetime)
	return retTime
}
