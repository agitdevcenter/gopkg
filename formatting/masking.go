package formatting

import "strings"

const CharacterX = byte('X')
const CharacterAsterisk = byte('*')

//MaskingNameExceptLastThree mask name except last 3 character
//when name only have 3 character or less than that, no masking required
//example
// Johnny Depp: ***nny *epp
// Nur Ady: Nur Ady
func MaskingName(customerName string) string {
	if customerName == "" {
		return ""
	}
	splits := strings.Split(customerName, " ")
	for i, v := range splits {
		splits[i] = MaskExceptLastNthCharacter(v, 3, CharacterAsterisk)
	}
	return strings.Join(splits, " ")
}

//MaskingNameNewFormat mask name with new format
//different masking name format between 3,4,5... character
//example
// Johnny Depp: Jo***y De**
// Nur Ady: N** A**
// Matt Le Tissier: Ma** Le Ti****r
func MaskingNameNewFormat(customerName string) string {
	if customerName == "" {
		return ""
	}
	splits := strings.Split(customerName, " ")
	for i, v := range splits {
		switch {
		case len(v) > 4:
			splits[i] = v[:2] + MaskExceptLastNthCharacter(v[2:], 1, CharacterAsterisk)
		case len(v) > 3:
			splits[i] = v[:2] + MaskExceptLastNthCharacter(v[2:], 0, CharacterAsterisk)
		case len(v) > 2:
			splits[i] = v[:1] + MaskExceptLastNthCharacter(v[1:], 0, CharacterAsterisk)
		default:
			splits[i] = v
		}
	}
	return strings.Join(splits, " ")
}

//MaskingPhoneNumber mask phone number except last 4 character
//example 08000000111: xxxxxxx0111
func MaskingPhoneNumber(phoneNumber string) string {
	return MaskExceptLastNthCharacter(phoneNumber, 4, CharacterX)
}

func MaskExceptLastNthCharacter(plain string, n int, mask byte) string {
	if len(plain) <= n {
		return plain
	}
	rs := []byte(plain)
	for i := range rs[:len(rs)-n] {
		rs[i] = mask
	}
	return string(rs)
}
