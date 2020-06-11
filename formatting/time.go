package formatting

import (
	"strings"
	"time"
)

const ZeroDate = "0001-01-01 00:00:00"

//ConvertFormat function that accept java format then replace with golang fromat
// yyMMddHHmmss return 060102150405
// dd-MMM-yyyy HH:mm:ss return 02-Jan-2006 15:04:05
// dd/MMMM/yyyy HH:mm:ssSSS return 02/January/2006 15:04:05.000
func ConvertFormat(javaFormat string) string {
	golangFormat := strings.ReplaceAll(javaFormat, "MMMM", "January")
	golangFormat = strings.ReplaceAll(golangFormat, "MMM", "Jan")
	golangFormat = strings.ReplaceAll(golangFormat, "yyyy", "2006")
	golangFormat = strings.ReplaceAll(golangFormat, "yy", "06")
	golangFormat = strings.ReplaceAll(golangFormat, "dd", "02")
	golangFormat = strings.ReplaceAll(golangFormat, "MM", "01")
	golangFormat = strings.ReplaceAll(golangFormat, "HH", "15")
	golangFormat = strings.ReplaceAll(golangFormat, "mm", "04")
	golangFormat = strings.ReplaceAll(golangFormat, "ss", "05")
	golangFormat = strings.ReplaceAll(golangFormat, "SSS", ".000")
	return golangFormat
}

const (
	Timestamp               = "060102150405"
	TimestampMillis         = "060102150405.000"
	TimestampFullYear       = "20060102150405"
	TimestampFullYearMillis = "20060102150405.000"
)

//FormatTimestamp return string with format yyMMddHHmmss
// 5 March 2020 09:37:10 -> "200305093710"
func FormatTimestamp(t time.Time) string {
	return t.Format(Timestamp)
}

//FormatTimestamp return string with format yyMMddHHmmssSSS
// 5 March 2020 09:37:10.123 -> "200305093710123"
func FormatTimestampMillis(t time.Time) string {
	return strings.ReplaceAll(t.Format(TimestampMillis), ".", "")
}

//FormatTimestampFullYear return string with format yyyyMMddHHmmss
// 5 March 2020 09:37:10 -> "20200305093710"
func FormatTimestampFullYear(t time.Time) string {
	return t.Format(TimestampFullYear)
}

//FormatTimestampFullYear return string with format yyyyMMddHHmmssSSS
// 5 March 2020 09:37:10.123 -> "20200305093710123"
func FormatTimestampFullYearMillis(t time.Time) string {
	return strings.ReplaceAll(t.Format(TimestampFullYearMillis), ".", "")
}

const (
	DateWithSlash          = "02/01/2006"
	DateWithDash           = "02-01-2006"
	DateShortMonth         = "02 Jan 2006"
	DateShortMonthWithDash = "02-Jan-2006"
	DateLongMonth          = "02 January 2006"
	DateLongMonthWithDash  = "02-January-2006"
)

//FormatDateWithDash return string with format dd-MM-yyyy
// 5 March 2020 09:37:10 -> "05-03-2020"
func FormatDateWithDash(t time.Time) string {
	return t.Format(DateWithDash)
}

//FormatDateWithSlash return string with format dd/MM/yyyy
// 5 March 2020 09:37:10 -> "05/03/2020"
func FormatDateWithSlash(t time.Time) string {
	return t.Format(DateWithSlash)
}

//FormatDateShortMonth return string with format dd MMM yyyy
// 5 March 2020 09:37:10 -> "05 Mar 2020"
func FormatDateShortMonth(t time.Time) string {
	return t.Format(DateShortMonth)
}

//FormatDateShortMonthWithDash return string with format dd-MMM-yyyy
// 5 March 2020 09:37:10 -> "05-Mar-2020"
func FormatDateShortMonthWithDash(t time.Time) string {
	return t.Format(DateShortMonthWithDash)
}

//FormatDateLongMonth return string with format dd MMMM yyyy
// 5 March 2020 09:37:10 -> "05 March 2020"
func FormatDateLongMonth(t time.Time) string {
	return t.Format(DateLongMonth)
}

//FormatDateLongMonthWithDash return string with format dd-MMMM-yyyy
// 5 March 2020 09:37:10 -> "05-March-2020"
func FormatDateLongMonthWithDash(t time.Time) string {
	return t.Format(DateLongMonthWithDash)
}

const (
	DateISO8601                  = "20060102"
	DateISO8601ShortYear         = "060102"
	DateISO8601WithDash          = "2006-01-02"
	DateISO8601ShortYearWithDash = "06-01-02"
)

//FormatDateISO8601 return string with format yyyyMMdd
// 5 March 2020 09:37:10 -> "20200305"
func FormatDateISO8601(t time.Time) string {
	return t.Format(DateISO8601)
}

//FormatDateISO8601ShortYear return string with format yyMMdd
// 5 March 2020 09:37:10 -> "200305"
func FormatDateISO8601ShortYear(t time.Time) string {
	return t.Format(DateISO8601ShortYear)
}

//FormatDateISO8601WithDash return string with format yyyy-MM-dd
// 5 March 2020 09:37:10 -> "2020-03-05"
func FormatDateISO8601WithDash(t time.Time) string {
	return t.Format(DateISO8601WithDash)
}

//FormatDateISO8601ShortYearWithDash return string with format yy-MM-dd
// 5 March 2020 09:37:10 -> "20-03-05"
func FormatDateISO8601ShortYearWithDash(t time.Time) string {
	return t.Format(DateISO8601ShortYearWithDash)
}

const (
	Time                  = "15:04:05"
	TimeShort             = "15:04"
	TimeWithoutColon      = "150405"
	TimeShortWithoutColon = "1504"
)

//FormatTime return string with format HH:mm:ss
// 5 March 2020 09:37:10 -> "09:37:10"
func FormatTime(t time.Time) string {
	return t.Format(Time)
}

//FormatTimeShort return string with format HH:mm
// 5 March 2020 09:37:10 -> "09:37"
func FormatTimeShort(t time.Time) string {
	return t.Format(TimeShort)
}

//FormatTimeWithoutColon return string with format HH:mm
// 5 March 2020 09:37:10 -> "093710"
func FormatTimeWithoutColon(t time.Time) string {
	return t.Format(TimeWithoutColon)
}

//FormatTimeShortWithoutColon return string with format HH:mm
// 5 March 2020 09:37:10 -> "0937"
func FormatTimeShortWithoutColon(t time.Time) string {
	return t.Format(TimeShortWithoutColon)
}

const (
	ReadableDateTime                   = "02 Jan 2006, 15:04:05"
	ReadableDateTimeShort              = "02 Jan 2006, 15:04"
	ReadableDateTimeLongMonth          = "02 January 2006, 15:04:05"
	ReadableDateTimeWithSlash          = "02/Jan/2006, 15:04:05"
	ReadableDateTimeShortWithSlash     = "02/Jan/2006, 15:04"
	ReadableDateTimeLongMonthWithSlash = "02/January/2006, 15:04:05"
)

//FormatTimeShortWithoutColon return string with format dd MM yyyy HH:mm:ss
// 5 March 2020 09:37:10 -> "05 Mar 2020, 09:37:10"
func FormatReadableDateTime(t time.Time) string {
	return t.Format(ReadableDateTime)
}

//FormatTimeShortWithoutColon return string with format dd MM yyyy HH:mm
// 5 March 2020 09:37:10 -> "05 Mar 2020, 09:37"
func FormatReadableDateTimeShort(t time.Time) string {
	return t.Format(ReadableDateTimeShort)
}

//FormatTimeShortWithoutColon return string with format dd MM yyyy HH:mm:ss
// 5 March 2020 09:37:10 -> "05 March 2020, 09:37:10"
func FormatReadableDateTimeLongMonth(t time.Time) string {
	return t.Format(ReadableDateTimeLongMonth)
}

//ReadableDateTimeWithSlash return string with format dd/MM/yyyy HH:mm:ss
// 5 March 2020 09:37:10 -> "05/Mar/2020, 09:37:10"
func FormatReadableDateTimeWithSlash(t time.Time) string {
	return t.Format(ReadableDateTimeWithSlash)
}

//ReadableDateTimeWithSlash return string with format dd/MM/yyyy HH:mm
// 5 March 2020 09:37:10 -> "05/Mar/2020, 09:37"
func FormatReadableDateTimeShortWithSlash(t time.Time) string {
	return t.Format(ReadableDateTimeShortWithSlash)
}

//ReadableDateTimeWithSlash return string with format dd/MM/yyyy HH:mm
// 5 March 2020 09:37:10 -> "05/March/2020, 09:37"
func FormatReadableDateTimeLongMonthWithSlash(t time.Time) string {
	return t.Format(ReadableDateTimeLongMonthWithSlash)
}
