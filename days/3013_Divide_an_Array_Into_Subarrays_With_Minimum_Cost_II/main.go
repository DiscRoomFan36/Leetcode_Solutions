package main

import "github.com/alexandrestein/gods/trees/redblacktree"

func minimumCost(nums []int, k int, dist int) int64 {
    running_minimum_set := make_multiset()
    running_rest_set    := make_multiset()
    running_sum  := 0

    Adjust_This := func() {
        for running_minimum_set.Size() < k-2 && !running_rest_set.IsEmpty() {
            x, ok := running_rest_set.First()
            if ok {
                running_rest_set.Remove(x)
                running_minimum_set.Add(x)
                running_sum += x
            }
        }

        for running_minimum_set.Size() > k-2 {
            x, ok := running_minimum_set.Last()
            if ok {
                running_minimum_set.Remove(x)
                running_rest_set.Add(x)
                running_sum -= x
            }
        }
    }

    Add_This := func(x int) {
        if !running_rest_set.IsEmpty() {
            first, ok := running_rest_set.First()
            if ok && x >= first {
                running_rest_set.Add(x)
            } else {
                running_minimum_set.Add(x)
                running_sum += x
            }
        } else {
            running_minimum_set.Add(x)
            running_sum += x
        }
        Adjust_This()
    }



    for i := 1; i < k-1; i++ { Add_This(nums[i]) }

    best_cost := running_sum + nums[k-1]
    for i := k; i < len(nums); i++ {
        j := i - dist - 1
        if j > 0 {
            x := nums[j]
            // remove 1 copy of x
            if running_minimum_set.Contains(x) {
                running_minimum_set.Remove(x)
                running_sum -= x
            } else if running_rest_set.Contains(x) {
                running_rest_set.Remove(x)
            }
            Adjust_This()
        }

        Add_This(nums[i-1])
        current_cost := running_sum + nums[i]
        if current_cost < best_cost { best_cost = current_cost }
    }

    result := best_cost + nums[0]
    return int64(result)
}


// allows for quick accsess of both, the smallest element, and the largest element.
type MultiSet struct {
    tree   *redblacktree.Tree
    // how many times something appears
    counter map[int]int
    size    int
}

func make_multiset() MultiSet {
    return MultiSet{
        tree:    redblacktree.NewWithIntComparator(),
        counter: make(map[int]int),
        size:    0,
    }
}

func (ms *MultiSet) Add(x int) {
    if count, exists := ms.counter[x]; exists {
        ms.counter[x] = count + 1
    } else {
        ms.counter[x] = 1
        ms.tree.Put(x, struct{}{})
    }
    ms.size += 1
}

func (ms *MultiSet) Remove(x int) bool {
    count, exists := ms.counter[x];
    if !exists { return false }

    if count == 1 {
        delete(ms.counter, x)
        ms.tree.Remove(x)
    } else {
        ms.counter[x] = count - 1
    }
    ms.size -= 1
    return true
}

func (ms *MultiSet) Size() int     { return ms.size }
func (ms *MultiSet) IsEmpty() bool { return ms.size == 0 }

func (ms *MultiSet) First() (int, bool) {
	if ms.tree.Empty() {
		return 0, false
	}
	return ms.tree.Left().Key.(int), true
}

func (ms *MultiSet) Last() (int, bool) {
	if ms.tree.Empty() {
		return 0, false
	}
	return ms.tree.Right().Key.(int), true
}

func (ms *MultiSet) Contains(x int) bool {
	_, exists := ms.counter[x]
	return exists
}
