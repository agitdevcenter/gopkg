package formatting

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	location, _ = time.LoadLocation("Asia/Jakarta")
	baseTime    = time.Date(2020, 3, 5, 9, 37, 10, 123000000, location)
)

func TestConvertFormat(t *testing.T) {
	assert.Equal(t, "060102150405", ConvertFormat("yyMMddHHmmss"))
	assert.Equal(t, "02-01-2006 15:04:05", ConvertFormat("dd-MM-yyyy HH:mm:ss"))
	assert.Equal(t, "02 Jan 2006, 15:04:05", ConvertFormat("dd MMM yyyy, HH:mm:ss"))
	assert.Equal(t, "02/January/2006 15:04:05.000", ConvertFormat("dd/MMMM/yyyy HH:mm:ssSSS"))
}

func TestFormatTimestamp(t *testing.T) {
	assert.Equal(t, "200305093710", FormatTimestamp(baseTime))
}

func TestFormatTimestampMillis(t *testing.T) {
	assert.Equal(t, "200305093710123", FormatTimestampMillis(baseTime))
}

func TestFormatTimestampFullYear(t *testing.T) {
	assert.Equal(t, "20200305093710", FormatTimestampFullYear(baseTime))
}

func TestFormatTimestampFullYearMillis(t *testing.T) {
	assert.Equal(t, "20200305093710123", FormatTimestampFullYearMillis(baseTime))
}

func TestFormatDateWithDash(t *testing.T) {
	assert.Equal(t, "05-03-2020", FormatDateWithDash(baseTime))
}

func TestFormatDateWithSlash(t *testing.T) {
	assert.Equal(t, "05/03/2020", FormatDateWithSlash(baseTime))
}

func TestFormatDateShortMonth(t *testing.T) {
	assert.Equal(t, "05 Mar 2020", FormatDateShortMonth(baseTime))
}

func TestFormatDateShortMonthWithDash(t *testing.T) {
	assert.Equal(t, "05-Mar-2020", FormatDateShortMonthWithDash(baseTime))
}

func TestFormatDateLongMonth(t *testing.T) {
	assert.Equal(t, "05 March 2020", FormatDateLongMonth(baseTime))
}

func TestFormatDateLongMonthWithDash(t *testing.T) {
	assert.Equal(t, "05-March-2020", FormatDateLongMonthWithDash(baseTime))
}

func TestFormatDateISO8601(t *testing.T) {
	assert.Equal(t, "20200305", FormatDateISO8601(baseTime))
}

func TestFormatDateISO8601ShortYear(t *testing.T) {
	assert.Equal(t, "200305", FormatDateISO8601ShortYear(baseTime))
}

func TestFormatDateISO8601WithDash(t *testing.T) {
	assert.Equal(t, "2020-03-05", FormatDateISO8601WithDash(baseTime))
}

func TestFormatDateISO8601ShortYearWithDash(t *testing.T) {
	assert.Equal(t, "20-03-05", FormatDateISO8601ShortYearWithDash(baseTime))
}

func TestFormatTime(t *testing.T) {
	assert.Equal(t, "09:37:10", FormatTime(baseTime))
}

func TestFormatTimeShort(t *testing.T) {
	assert.Equal(t, "09:37", FormatTimeShort(baseTime))
}

func TestFormatTimeWithoutColon(t *testing.T) {
	assert.Equal(t, "093710", FormatTimeWithoutColon(baseTime))
}

func TestFormatTimeShortWithoutColon(t *testing.T) {
	assert.Equal(t, "0937", FormatTimeShortWithoutColon(baseTime))
}

func TestFormatReadableDateTime(t *testing.T) {
	assert.Equal(t, "05 Mar 2020, 09:37:10", FormatReadableDateTime(baseTime))
}

func TestFormatReadableDateTimeShort(t *testing.T) {
	assert.Equal(t, "05 Mar 2020, 09:37", FormatReadableDateTimeShort(baseTime))
}

func TestFormatReadableDateTimeLongMonth(t *testing.T) {
	assert.Equal(t, "05 March 2020, 09:37:10", FormatReadableDateTimeLongMonth(baseTime))
}

func TestFormatReadableDateTimeWithSlash(t *testing.T) {
	assert.Equal(t, "05/Mar/2020, 09:37:10", FormatReadableDateTimeWithSlash(baseTime))
}

func TestFormatReadableDateTimeShortWithSlash(t *testing.T) {
	assert.Equal(t, "05/Mar/2020, 09:37", FormatReadableDateTimeShortWithSlash(baseTime))
}

func TestFormatReadableDateTimeLongMonthWithSlash(t *testing.T) {
	assert.Equal(t, "05/March/2020, 09:37:10", FormatReadableDateTimeLongMonthWithSlash(baseTime))
}
