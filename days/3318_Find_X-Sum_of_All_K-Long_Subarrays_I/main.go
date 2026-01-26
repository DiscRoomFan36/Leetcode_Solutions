package main

import (
	"cmp"
	"slices"
)

type Num_And_Occurence struct {
    num int
    occur int
}

func add_to(num_and_occur *[]Num_And_Occurence, num int) {
    for i := range *num_and_occur {
        if (*num_and_occur)[i].num == num {
            (*num_and_occur)[i].occur += 1
            return
        }
    }

    Append(num_and_occur, Num_And_Occurence{num: num, occur: 1})
}

func remove_from(num_and_occur *[]Num_And_Occurence, num int) {
    for i := range *num_and_occur {
        if num == (*num_and_occur)[i].num {
            (*num_and_occur)[i].occur -= 1
            if (*num_and_occur)[i].occur == 0 {
                // remove it from the array, there are none left.
                Remove_Unordered(num_and_occur, i)
            }
            return
        }
    }
    panic("Tried to remove a number that was not there!")
}

func findXSum(nums []int, k int, x int) []int {
    if len(nums) == 0 { return []int{} }
    if k < x { panic("dont do this.") }

    // this can at most have k capacity.
    num_and_occur := make([]Num_And_Occurence, 0, k)
    // add inital elements
    for _, n := range nums[:k-1] { add_to(&num_and_occur, n) }

    result := make([]int, len(nums) - (k-1))
    for i := range len(nums) - (k-1) {
        // add the one thats new in the window.
        add_to(&num_and_occur, nums[i+k-1])

        slices.SortFunc(num_and_occur,
            func(a, b Num_And_Occurence) int {
                if a.occur != b.occur { return -cmp.Compare(a.occur, b.occur) }
                return -cmp.Compare(a.num, b.num)
            },
        )

        ith_result := 0
        for j := range min(x, len(num_and_occur)) {
            ith_result += (num_and_occur[j].num * num_and_occur[j].occur)
        }
        result[i] = ith_result

        // remove the one that's leaving the window.
        remove_from(&num_and_occur, nums[i])
    }

    return result
}


func Append[T any](slice *[]T, items ...T) {
	*slice = append(*slice, items...)
}
func Remove_Unordered[T any](slice *[]T, index int) {
	if index != len(*slice)-1 {
		(*slice)[index] = (*slice)[len(*slice)-1]
	}
	*slice = (*slice)[:len(*slice)-1]
}
