package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	totalSize := len1 + len2
	midIndex2 := totalSize / 2
	midIndex1 := midIndex2
	if totalSize % 2 == 0 {
		midIndex1 = midIndex2 - 1
	}
	index1 := 0
	index2 := 0
	target1 := 0
	target2 := 0
	target1Found := false
	target2Found := false
	for index := 0; index < totalSize; index ++ {
		var currentVal int
		val1 := 0
		val1Exist := false
		val2 := 0
		val2Exist := false
		if index1 < len1 {
			val1 = nums1[index1]
			val1Exist = true
		}
		if index2 < len2 {
			val2 = nums2[index2]
			val2Exist = true
		}
		if val1Exist && val2Exist {
			if val1 <= val2 {
				currentVal = val1
				index1 ++
			} else {
				currentVal = val2
				index2 ++
			}
		} else if val1Exist {
			currentVal = val1
			index1 ++
		} else if val2Exist {
			currentVal = val2
			index2 ++
		}

		if !target1Found && index == midIndex1 {
			target1 = currentVal
			target1Found = true
		}
		if !target2Found && index == midIndex2 {
			target2 = currentVal
			target2Found = true
		}
		if target1Found && target2Found {
			break
		}
	}
	return (float64(target1) + float64(target2)) / 2
}

func main() {
	nums1 := []int{1,3}
	nums2 := []int{2}
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("res: %v\n", res)
}
