package decode_the_morse_code

import (
	"strings"
)

// You may get original char by morse code like this: MORSE_CODE[char]

func EncodeMorse(s string) string {
	s = strings.TrimSpace(s)
	words := []string{}
	chars := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			words = append(words, strings.Join(chars, " "))
			chars = []string{}
			continue
		}
		enc, ok := MORSE_ENCODE_TABLE[s[i:i+1]]
		if !ok {
			panic("missing entry")
		}
		chars = append(chars, enc)
	}
	words = append(words, strings.Join(chars, " "))
	return strings.Join(words, "   ")
}

func FindMaxOneLen(bits string) int {
	maxLen := 0
	curLen := 0
	for i := 0; i < len(bits); i++ {
		if bits[i] == '0' {
			if maxLen < curLen {
				maxLen = curLen
			}
			curLen = 0
		} else {
			curLen++
		}
	}
	if maxLen < curLen {
		maxLen = curLen
	}
	return maxLen
}

func DecodeBits(bits string) string {
	bits = strings.Trim(bits, "0")
	maxOneLen := FindMaxOneLen(bits)
	var rate int
	for rate = maxOneLen; rate >= 1; rate-- {
		if len(bits)%rate > 0 {
			continue
		}
		skip := false
		for j := 0; j < len(bits); j += rate {
			if strings.Contains(bits[j:j+rate], "0") && strings.Contains(bits[j:j+rate], "1") {
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		break
	}

	units := ""
	for i := 0; i < len(bits); i += rate {
		units += bits[i : i+1]
	}

	decoded := ""
	for i := 0; i < len(units); {
		if i+7 <= len(units) && !strings.Contains(units[i:i+7], "1") {
			// pause between words
			decoded += "   "
			i += 7
		} else if i+3 <= len(units) && !strings.Contains(units[i:i+3], "1") {
			// pause between characters
			decoded += " "
			i += 3
		} else if i+3 <= len(units) && !strings.Contains(units[i:i+3], "0") {
			// dash
			decoded += "-"
			i += 3
		} else if units[i:i+1] == "1" {
			// dot
			decoded += "."
			i++
		} else {
			// pause between dots and dashes in a character
			i++
		}
	}
	return decoded
}
