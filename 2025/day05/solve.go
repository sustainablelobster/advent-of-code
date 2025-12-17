package main

import (
	"fmt"
	"os"
	"sort"
)

type interval struct {
	start int
	end   int
}

func main() {
	infile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	var rawIntervals []interval

	for {
		var currInterval interval

		_, err = fmt.Fscanf(infile, "%d-%d", &currInterval.start, &currInterval.end)
		if err != nil {
			break
		}

		rawIntervals = append(rawIntervals, currInterval)
	}

	sort.Slice(rawIntervals, func(i, j int) bool {
		return rawIntervals[i].start < rawIntervals[j].start
	})

	var intervals []interval

	current := rawIntervals[0]
	for _, next := range rawIntervals[1:] {
		if next.start > current.end {
			intervals = append(intervals, current)
			current = next
		} else if next.end > current.end {
			current.end = next.end
		}
	}
	intervals = append(intervals, current)

	answer1 := 0

	for {
		var n int

		_, err = fmt.Fscanf(infile, "%d", &n)
		if err != nil {
			break
		}

		for _, currInterval := range intervals {
			if n >= currInterval.start && n <= currInterval.end {
				answer1++
				break
			}
		}
	}

	answer2 := 0

	for _, currInterval := range intervals {
		answer2 += currInterval.end - currInterval.start + 1
	}

	fmt.Println("Answer 1:", answer1)
	fmt.Println("Answer 2:", answer2)
}
