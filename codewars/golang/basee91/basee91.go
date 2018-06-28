package basee91

import "fmt"

const TTABLE = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~\""

func Min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

func getDecimal(index, shift, nbits uint, bytes []byte) (uint, uint, uint) {
	var result uint = 0
	var bitCnt uint = 0
	for int(index) < len(bytes) && nbits > 0 {
		n := Min(nbits, 8-shift)
		result += uint(((bytes[index] >> shift) & ((1 << n) - 1)) * (1 << bitCnt))
		bitCnt += n
		nbits -= n
		shift += n
		if shift >= 8 {
			shift = 0
			index++
		}
	}
	return result, index, shift
}

func Encode(d []byte) []byte {
	var index, shift uint
	encoding := []byte{}
	for int(index) < len(d) {
		val, newIndex, newShift := getDecimal(index, shift, 13, d)
		fmt.Println(val)
		fmt.Println(newIndex)
		fmt.Println(newShift)
		if val >= 89 {
			encoding = append(encoding, TTABLE[val/91])
			encoding = append(encoding, TTABLE[val%91])
		} else {
			val, newIndex, newShift = getDecimal(index, shift, 14, d)
			encoding = append(encoding, TTABLE[val/91])
			encoding = append(encoding, TTABLE[val%91])
		}
		index, shift = newIndex, newShift
	}
	return encoding
}

func Decode(d []byte) []byte {
	return nil // do it!
}
