package formatting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatAmount(t *testing.T) {
	amount := int64(1230000)
	result := FormatAmount(amount)
	assert.Equal(t, "Rp1.230.000", result)
}

func TestSanitizeName(t *testing.T) {
	name := "Nama mama"
	result := SanitizeName(name)
	assert.Equal(t, name, result)

	name = "null papa"
	result = SanitizeName(name)
	assert.Equal(t, "papa", result)

	name = "Nama null"
	result = SanitizeName(name)
	assert.Equal(t, "Nama", result)

	name = ""
	result = SanitizeName(name)
	assert.Equal(t, "-", result)

	name = "null"
	result = SanitizeName(name)
	assert.Equal(t, "-", result)
}

func TestSanitizePhoneNumber(t *testing.T) {
	phoneNumber := "628000000111"
	result := SanitizePhoneNumber(phoneNumber)
	assert.Equal(t, phoneNumber, result)

	phoneNumber = "08123000111"
	result = SanitizePhoneNumber(phoneNumber)
	assert.Equal(t, "628123000111", result)

	phoneNumber = "8123000111"
	result = SanitizePhoneNumber(phoneNumber)
	assert.Equal(t, "628123000111", result)

	phoneNumber = "02100020202"
	result = SanitizePhoneNumber(phoneNumber)
	assert.Equal(t, "02100020202", result)
}
