package formatting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaskCustomerName(t *testing.T) {
	name := "Johnny Depp"
	result := MaskingName(name)
	assert.Equal(t, "***nny *epp", result)

	name = "Nur Ady"
	result = MaskingName(name)
	assert.Equal(t, "Nur Ady", result)

	name = ""
	result = MaskingName(name)
	assert.Equal(t, "", result)
}

func TestMaskCustomerNameNewFormat(t *testing.T) {
	name := "Johnny Depp"
	result := MaskingNameNewFormat(name)
	assert.Equal(t, "Jo***y De**", result)

	name = "Nur Ady"
	result = MaskingNameNewFormat(name)
	assert.Equal(t, "N** A**", result)

	name = "Matt Le Tissier"
	result = MaskingNameNewFormat(name)
	assert.Equal(t, "Ma** Le Ti****r", result)

	name = ""
	result = MaskingNameNewFormat(name)
	assert.Equal(t, "", result)
}

func TestMaskPhoneNumber(t *testing.T) {
	phoneNum := "080000000112"
	result := MaskingPhoneNumber(phoneNum)
	assert.Equal(t, "xxxxxxxx0112", result)
}
