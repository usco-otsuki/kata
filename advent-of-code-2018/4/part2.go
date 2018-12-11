package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	sort.Strings(lines)
	lines = append(lines, "[0000-00-00 00:00] Guard #0 begins shift")

	re := regexp.MustCompile(`\[.+:(\d+)\] (.+)`)
	guard := 0
	sleptAt := 0
	counter := make(map[int]map[int]int)
	sum := make(map[int]int)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches == nil {
			log.Fatalf("not match")
		}
		var min int
		fmt.Sscanf(matches[1], "%d", &min)
		msg := matches[2]
		switch msg {
		case "falls asleep":
			sleptAt = min
		case "wakes up":
			if _, ok := counter[guard]; !ok {
				counter[guard] = make(map[int]int)
			}
			for i := sleptAt; i < min; i++ {
				counter[guard][i]++
			}
			sum[guard] += min - sleptAt
		default:
			if n, err := fmt.Sscanf(msg, "Guard #%d begins shift", &guard); n == 0 || err != nil {
				log.Fatalf("Could not parse 'guard begins shift'")
			}
		}
	}
	maxGuard := 0
	maxCount := 0
	maxMinute := 0
	for guard, minutes := range counter {
		for minute, c := range minutes {
			if maxCount < c {
				maxCount = c
				maxGuard = guard
				maxMinute = minute
			}
		}
	}
	fmt.Println(maxGuard, maxMinute)
	fmt.Println(maxGuard * maxMinute)
}
