package main

func gcd(values ...int) int {
	if len(values) == 0 {
		panic("")
	}
	if len(values) == 1 {
		return values[0]
	}
	a := values[0]
	b := gcd(values[1:]...)
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func hasGroupsSizeX(deck []int) bool {
	cnt := map[int]int{}
	for _, val := range deck {
		cnt[val]++
	}
	uniqValues := []int{}
	for _, val := range cnt {
		uniqValues = append(uniqValues, val)
	}
	x := gcd(uniqValues...)
	if x >= 2 {
		return true
	}
	return false
}

func main() {}
