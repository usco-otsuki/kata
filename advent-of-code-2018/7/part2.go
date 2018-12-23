package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	NUM_WORKERS = 5
	BASE_SECOND = 60
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	graph := make(map[string][]string)
	count := make(map[string]int)
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
	queue := []string{}
	for c := 'A'; c <= 'Z'; c++ {
		count[string(c)] = BASE_SECOND + int(c-'A') + 1
		if _, ok := isDest[string(c)]; !ok {
			if _, ok := graph[string(c)]; ok {
				queue = append(queue, string(c))
			}
		}
	}

	t := -1
	workers := []string{}
	for len(workers) > 0 || len(queue) > 0 {
		sort.Strings(queue)
		newWorkers := []string{}
		for i := 0; i < len(workers); i++ {
			count[workers[i]]--
			if count[workers[i]] > 0 {
				newWorkers = append(newWorkers, workers[i])
			} else {
				list, ok := graph[workers[i]]
				if !ok {
					continue
				}
				for _, item := range list {
					toVisit[item]--
					if toVisit[item] == 0 {
						queue = append(queue, item)
					}
				}
			}
		}
		workers = newWorkers
		numWorkers := len(workers)
		for i := 0; i < NUM_WORKERS-numWorkers && len(queue) > 0; i++ {
			top := queue[0]
			queue = queue[1:]
			for len(queue) > 0 && queue[0] == top {
				queue = queue[1:]
			}
			workers = append(workers, top)
		}

		t++

	}
	fmt.Println(t)
}
