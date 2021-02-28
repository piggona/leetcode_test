package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1)+len(nums2) == 0 {
		return 0
	}
	length := len(nums1) + len(nums2)
	var odd = func() bool {
		if length%2 == 0 {
			return false
		}
		return true
	}()
	var remain = func() int {
		if odd {
			return length / 2
		}
		return length/2 - 1
	}()

	return recursiveMedian(nums1, nums2, remain, odd)
}

// 随时可以被改为取两个数组的第K个数
func recursiveMedian(nums1 []int, nums2 []int, remain int, odd bool) float64 {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	if len(nums2) == 0 {
		if odd {
			return float64(nums1[remain])
		}
		return (float64(nums1[remain]) + float64(nums1[remain+1])) / 2
	}
	if remain == 0 {
		if odd {
			return float64(min(nums1[0], nums2[0]))
		}
		a, b := getSmallestTwoItems(nums1, nums2)
		return (float64(a) + float64(b)) / 2
	}
	if remain == 1 {
		remain--
		if nums1[0] < nums2[0] {
			nums1 = nums1[1:]
		} else {
			nums2 = nums2[1:]
		}
		return recursiveMedian(nums1, nums2, remain, odd)
	}
	var pos = (remain / 2) - 1
	if len(nums2) < (remain / 2) {
		pos = len(nums2) - 1
	}
	if nums1[pos] == nums2[pos] {
		return recursiveMedian(nums1[pos+1:], nums2[pos+1:], remain-2*(pos+1), odd)
	}
	if nums1[pos] < nums2[pos] {
		return recursiveMedian(nums1[pos+1:], nums2, remain-pos-1, odd)
	}
	return recursiveMedian(nums1, nums2[pos+1:], remain-pos-1, odd)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func getSmallestTwoItems(nums1 []int, nums2 []int) (int, int) {
	if len(nums1)+len(nums2) < 2 {
		panic("don't have enough items...")
	}
	var i, j int
	var result = make([]int, 2)
	var cur int
	for cur < 2 {
		if i >= len(nums1) {
			for cur < 2 {
				result[cur] = nums2[j]
				j++
				cur++
			}
			break
		}
		if j >= len(nums2) {
			for cur < 2 {
				result[cur] = nums1[i]
				i++
				cur++
			}
			break
		}
		if nums1[i] < nums2[j] {
			result[cur] = nums1[i]
			i++
		} else {
			result[cur] = nums2[j]
			j++
		}
		cur++
	}
	return result[0], result[1]
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{1, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

}
