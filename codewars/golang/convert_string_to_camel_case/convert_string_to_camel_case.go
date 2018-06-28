package convert_string_to_camel_case

import "strings"

func toCamelCase(words []string) []string {
	result := []string{}
	for i, word := range words {
		if i > 0 {
			result = append(result, strings.ToUpper(word[0:1])+word[1:])
		} else {
			result = append(result, word)
		}
	}
	return result
}

func ToCamelCase(s string) string {
	word := ""
	words := []string{}
	for _, c := range s {
		if c == '_' || c == '-' {
			if len(word) > 0 {
				words = append(words, word)
			}
			word = ""
			continue
		}
		word += string(c)
	}
	if len(word) > 0 {
		words = append(words, word)
	}
	return strings.Join(toCamelCase(words), "")
}
