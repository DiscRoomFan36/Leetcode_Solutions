
package main

func advantageCount(nums1 []int, nums2 []int) []int {
    n := len(nums1)

    type NumAndBool struct {
        num int32 // for the cache. "0 <= nums1[i], nums2[i] <= 10^9"
        flag bool
    }

    nums_with_used := make([]NumAndBool, n)
    for i := range n {
        nums_with_used[i] = NumAndBool{
            num: int32(nums1[i]),
            flag: false,
        }
    }

    slices.SortFunc(nums_with_used,
        func(a, b NumAndBool) int {
            return cmp.Compare(a.num, b.num)
        },
    )

    // not always correct, but its a good starting point.
    smallest_unused_index := 0

    for i := range n {
        x := int32(nums2[i])
        index, found := slices.BinarySearchFunc(nums_with_used, x,
            func(a NumAndBool, b int32) int {
                return cmp.Compare(a.num, b)
            },
        )

        // we need numbers bigger than x
        if found { index += 1 }

        // find the smallest number bigger than x, (not in use already)
        for (index < n) && (nums_with_used[index].num == x || nums_with_used[index].flag) {
            index += 1
        }

        if index >= n {
            // get the smallest number
            smallest_index := -1
            for j := smallest_unused_index; j < index; j++ {
                if !nums_with_used[j].flag {
                    smallest_index = j
                    break
                }
            }
            if smallest_index == -1 { panic("this cannot happen") }

            smallest_unused_index = smallest_index + 1
            index = smallest_index
        }

        nums_with_used[index].flag = true
        nums1[i] = int(nums_with_used[index].num)
    }

    return nums1
}
