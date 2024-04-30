package main

import "fmt"

/*
Count Elements Greater Than Previous Average

Given an array of positive integers, return the number of elements that are strictly greater than the average of all previous elements. Skip the first element.
*/
func countResponseTimeRegressions(responseTimes []int32) int32 {
	if len(responseTimes) <= 1 {
		return 0
	}

	sum := int64(responseTimes[0])
	count := int32(0)

	for idx := 1; idx < len(responseTimes); idx++ {

		avg := float64(sum) / float64(idx)

		if float64(responseTimes[idx]) > avg {
			count++
		}

		sum += int64(responseTimes[idx])
	}

	return int32(count)
}

func main() {
	fmt.Println("Elements greater than previous average : ", countResponseTimeRegressions([]int32{100, 200, 150, 210, 300}))
	fmt.Println("Elements greater than previous average : ", countResponseTimeRegressions([]int32{100, 200, 150, 300}))
	/*
		- Day 0: 100 (no previous days, skip)
		- Day 1: 200 > average(100) = 100 → count = 1
		- Day 2: 150 vs average(100, 200) = 150 → not greater → count = 1
		- Day 3: 300 > average(100, 200, 150) = 150 → count = 2 Return 2.
	*/
}
