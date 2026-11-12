package topkfrequentelements

import (
	"cmp"
	"maps"
	"slices"
)

// Time: O(nlogn), Space: O(n)
func TopKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)

	for _, n := range nums {
		counts[n]++
	}

	sorted := slices.SortedFunc(maps.Keys(counts), func(a, b int) int {
		return cmp.Compare(counts[b], counts[a])
	})

	return sorted[:k]
}

// Time: O(n), Space: O(n)
func TopKFrequentBucketSort(nums []int, k int) []int {
	counts := make(map[int]int)

	for _, n := range nums {
		counts[n]++
	}

	buckets := make([][]int, len(nums)+1)

	for num, freq := range counts {
		buckets[freq] = append(buckets[freq], num)
	}

	result := make([]int, 0, k)

	for freq := len(buckets) - 1; freq >= 0 && len(result) < k; freq-- {
		for _, num := range buckets[freq] {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
	}

	return result

}
