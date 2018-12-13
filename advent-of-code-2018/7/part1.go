package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	graph := make(map[string][]string)
	toVisit := make(map[string]int)
	var from, to string
	isDest := make(map[string]bool)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "Step %s must be finished before step %s can begin.", &from, &to)
		if _, ok := graph[from]; !ok {
			graph[from] = []string{}
		}
		graph[from] = append(graph[from], to)
		toVisit[to]++
		isDest[to] = true
	}
	fmt.Println(toVisit)
	queue := []string{}
	for c := 'A'; c <= 'Z'; c++ {
		if _, ok := isDest[string(c)]; !ok {
			queue = append(queue, string(c))
		}
	}

	result := ""
	for len(queue) > 0 {
		sort.Strings(queue)
		top := queue[0]
		queue = queue[1:]
		for len(queue) > 0 && queue[0] == top {
			queue = queue[1:]
		}
		result += top
		list, ok := graph[top]
		if !ok {
			continue
		}
		for _, item := range list {
			toVisit[item]--
			if toVisit[item] == 0 {
				queue = append(queue, item)
			}
		}
		// fmt.Println(queue)
	}
	fmt.Println(result)
}
