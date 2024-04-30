package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	total := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[total] = nums1[i]
			i--
		} else {
			nums1[total] = nums2[j]
			j--
		}
		total--
	}

	for j >= 0 {
		nums1[total] = nums2[j]
		j--
		total--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0, 0} // the number of 0's in nums1 = total length of nums2
	nums2 := []int{4, 5, 6, 7}

	merge(nums1, 3, nums2, 4)
	fmt.Println(nums1)
}
