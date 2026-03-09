package main

func findFinalValue(nums []int, original int) int {
    nums_set := make(map[int]struct{});
    for _, n := range nums {
        // dont care about smaller numbers
        if n < original { continue;
        } else if n == original {
            // who, ray.
            original *= 2;
            continue;
        } else {
            // we could do some smarter check to see if this
            // could possibly be the original later,
            // since we always mult by 2.
            if n % 2 == 1 { continue; }
            nums_set[n] = struct{}{};
        }
    }

    for Contains(nums_set, original) { original *= 2; }

    return original;
}

func Contains[T comparable, U any](the_map map[T]U, item T) bool {
    _, found := the_map[item];
    return found;
}
