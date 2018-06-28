package esolang_interpreters2

func replaceString(index int, c byte, str string) string {
	if index < len(str) {
		str = str[:index] + string(c) + str[index+1:]
	} else {
		str = str[:index] + string(c)
	}
	return str
}

func findMatch(str string, open, close byte, index, sign int) int {
	index += sign
	for count := 1; count > 0; index += sign {
		if str[index] == open {
			count++
		} else if str[index] == close {
			count--
		}
	}
	return index
}

func Interpreter(code, tape string) string {
	cp := 0 // code pointer
	sp := 0 // storage pointer
	for cp >= 0 && cp < len(code) && sp >= 0 && sp < len(tape) {
		switch code[cp] {
		case '>':
			sp++
		case '<':
			sp--
		case '*':
			if tape[sp] == '0' {
				tape = replaceString(sp, '1', tape)
			} else {
				tape = replaceString(sp, '0', tape)
			}
		case '[':
			if tape[sp] == '0' {
				cp = findMatch(code, '[', ']', cp, 1)
				continue
			}
		case ']':
			if tape[sp] != '0' {
				cp = findMatch(code, ']', '[', cp, -1) + 1
				continue
			}
		}
		cp++
	}
	return tape
}
