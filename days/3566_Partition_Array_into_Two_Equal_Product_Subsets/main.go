package main

func checkEqualPartitions(nums []int, target int64) bool {
    product := 1;
    for _, n := range nums { product *= n; }
    // target and product will not overflow this,
    if int64(product) != target*target { return false; }

    var recur func(target int, i int) bool;
    recur = func(target int, i int) bool {
        if target == 1 { return true; }
        if i >= len(nums) { return false; }

        n := nums[i];
        if target % n == 0 {
            if recur(target / n, i + 1) {
                return true;
            }
        }

        return recur(target, i + 1);
    }

    return recur(int(target), 0);
}
