package decode_the_morse_code

import (
	"strings"
)

var MORSE_CODE map[string]string

func DecodeMorse(morseCode string) string {
	words := strings.Split(strings.TrimSpace(morseCode), "   ")
	words_decoded := []string{}
	for _, word := range words {
		chars := strings.Split(word, " ")
		chars_decoded := []string{}
		for _, char := range chars {
			chars_decoded = append(chars_decoded, MORSE_CODE[char])
		}
		words_decoded = append(words_decoded, strings.Join(chars_decoded, ""))
	}
	return strings.Join(words_decoded, " ")
}
