package main

func getCommon(nums1 []int, nums2 []int) int {
    i, j := 0, 0;
    for i < len(nums1) && j < len(nums2) {
        for i < len(nums1) && nums1[i] < nums2[j] {
            i += 1;
        }
        if i >= len(nums1) { break; }

        for j < len(nums2) && nums2[j] < nums1[i] {
            j += 1;
        }
        if j >= len(nums2) { break; }

        // if the same, it is the result.
        if nums1[i] == nums2[j] { return nums1[i]; }

        // now nums2[j] is bigger than nums1[i], just do the loop again.
    }

    return -1;
}
