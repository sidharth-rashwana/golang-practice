package main

import (
	"fmt"
	"sort"
)

/*
Maximum Number of Non-Overlapping Intervals
Given an array of intervals where each interval has a start and end time, return the maximum number of non-overlapping intervals.
*/

func maximizeNonOverlappingMeetings(meetings [][]int32) int32 {
	if len(meetings) == 0 {
		return 0
	}

	// Sort intervals by end time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][1] < meetings[j][1]
	})

	count := 0
	lastEnd := -1 << 31 // minimum int value

	for _, interval := range meetings {
		start, end := interval[0], interval[1]
		if start >= int32(lastEnd) {
			count++
			lastEnd = int(end)
		}
	}

	return int32(count)
}

func main() {
	meetings := [][]int32{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println(maximizeNonOverlappingMeetings(meetings))
}
