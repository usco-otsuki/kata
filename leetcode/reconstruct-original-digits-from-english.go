package main

func originalDigits(s string) string {
	result := ""
	cnt := map[rune]int{}
	for _, c := range s {
		cnt[c]++
	}
	freqs := make([]int, 10)
	freqs[0] = cnt['z']
	freqs[2] = cnt['w']
	freqs[4] = cnt['u']
	freqs[6] = cnt['x']
	freqs[8] = cnt['g']
	freqs[3] = cnt['h'] - freqs[8]
	freqs[1] = cnt['o'] - freqs[0] - freqs[2] - freqs[4]
	freqs[5] = cnt['f'] - freqs[4]
	freqs[9] = cnt['i'] - freqs[5] - freqs[6] - freqs[8]
	freqs[7] = (cnt['e'] - freqs[0] - freqs[1] - 2*freqs[3] - freqs[5] - freqs[8] - freqs[9]) / 2

	for i, freq := range freqs {
		num := strconv.Itoa(i)
		result += strings.Repeat(num, freq)
	}
	return result
	/*
	 z, e, r, o, n, t, w, h, f, u, i, v, s, x, g
	 zero -> z
	 one ->
	 two -> w
	 three ->
	 four ->
	 five ->
	 six -> x
	 seven ->
	 eight -> g
	 nine ->

	 c_0 + c_1 + 2 * c_3 + c_5 + 2 * c_7 + c_8 + c_9 = c_e
	 c_0 + c_3 + c_4 = c_r
	 c_0 + c_1 + c_2 + c_4 = c_o
	 c_1 + c_7 + 2 * c_9 = c_n
	 c_2 + c_3 + c_8 = c_t
	 c_2 = c_w
	 c_3 + c_8 = c_h
	 c_4 + c_5 = c_f
	 c_4 = c_u
	 c_5 + c_6 + c_8 + c_9 = c_i
	*/
}

func main() {}
