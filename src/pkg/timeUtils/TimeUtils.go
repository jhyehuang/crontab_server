package timeUtils

import (
	"github.com/jhyehuang/crontab_server/src/global/constant"
	"github.com/jhyehuang/crontab_server/src/result"
	"time"
)

func DaysAdd(date time.Time, days int64) time.Time {
	return date.Add(time.Duration(days) * 24 * time.Hour)
}

func DaysAddStr(date time.Time, days int64) string {
	return date.Add(time.Duration(days) * 24 * time.Hour).Format(constant.TimeStart)
}

func GetTimeStart(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, time.Local)
}

func GetTimeStrStart(tm time.Time) string {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, time.Local).Format(constant.TimeStart)
}

func GetHourStartStr(date time.Time, hours int64) string {
	tm := date.Add(time.Duration(hours) * time.Hour)
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), 0, 0, 0, time.Local).Format(constant.TimeStart)
}

func GetHourEndStr(date time.Time, hours int64) string {
	tm := date.Add(time.Duration(hours) * time.Hour)
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), 59, 59, 0, time.Local).Format(constant.TimeStart)
}

func GetTimeEnd(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), 23, 59, 59, 0, time.Local)
}

func GetTimeStrEnd(tm time.Time) string {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), 23, 59, 59, 0, time.Local).Format(constant.TimeStart)
}

func TimeStringFormat(tm time.Time) string {
	return tm.Format(constant.TimeStart)
}

// TransferPeriodToDays 1w/1m/1y转化天数
func TransferPeriodToDays(period string) (int, string) {
	if period == "" {
		panic(result.ParamError)
	}
	switch period {
	case "1w":
		return 7, "weekly"
	case "1m":
		return 30, "monthly"
	case "1y":
		return 365, "yearly"
	default:
		panic(result.ParamError)
	}
}

func GetMonthStart(d time.Time) time.Time {
	return GetTimeStart(d.AddDate(0, 0, -d.Day()+1))
}

func GetMonthEnd(d time.Time) time.Time {
	return GetTimeEnd(GetMonthStart(d).AddDate(0, 1, -1))
}

func GetYesterday() string {
	currentTime := time.Now()
	yesterday := currentTime.AddDate(0, 0, -1)
	yesterdayStr := yesterday.Format("2006-01-02")
	return yesterdayStr
}

func GetToday() string {
	currentTime := time.Now()
	yesterday := currentTime.AddDate(0, 0, 0)
	yesterdayStr := yesterday.Format("2006-01-02")
	return yesterdayStr
}

func ConvertToSeconds(dateStr string) (uint64, error) {
	// 解析日期字符串
	date, err := time.Parse("2006-01-02 15:04:05-07:00", dateStr)
	if err != nil {
		return 0, err
	}
	// 转换为秒数（Unix 时间戳）
	seconds := date.Unix()
	return uint64(seconds), nil
}

func GetEpochByTime(dateStr string, base uint64) (uint64, error) {
	targetTimestamp, err := ConvertToSeconds(dateStr)
	if err != nil {
		return 0, err
	}
	epoch := uint64((targetTimestamp - base) / 30)

	return epoch, nil
}

func GetTimeByEpoch(epoch uint64, base uint64) (uint64, error) {
	targetTimestamp := epoch*30 + base
	return targetTimestamp, nil
}
