package tools

import "time"

// Time2String 时间转string
func Time2String(times time.Time, showHour bool) string {
	var cstZone = time.FixedZone("CST", 8*3600)
	if showHour {
		return times.In(cstZone).Format("2006-01-02 15:04:05")
	}
	return times.In(cstZone).Format("2006-01-02")
}

// Str2Time string转时间
func Str2Time(times string, hour bool) time.Time {
	if loc, err := time.LoadLocation("Local"); err == nil {
		var theTime time.Time
		// 是否为带时间的转换
		if hour {
			theTime, err = time.ParseInLocation("2006-01-02 15:04:05", times, loc)
		} else {
			theTime, err = time.ParseInLocation("2006-01-02", times, loc)
		}
		// 判断是否转换成功
		if err == nil {
			return theTime
		}
	}
	return time.Now()
}
