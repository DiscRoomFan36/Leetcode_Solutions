package main

func maxArea(height []int) int {
    if len(height) <= 1 { return 0 }

    // 2 pointers, going inwards.
    i, j := 0, len(height)-1

    get_area := func(i, j int) int {
        // this is a good bounds check, this function would work
        // without it but this gets my intent across.
        if j <= i { return 0 }
        wide := j - i
        high := min(height[i], height[j])
        return high * wide
    }

    best_area := get_area(i, j)

    while_loop: for {

        // move left forward until its bigger than j,
        for height[i] < height[j] {
            i += 1
            if i >= j { break while_loop; }
            best_area = max(best_area, get_area(i, j))
        }

        // move right backward until its bigger than i,
        for height[j] < height[i] {
            j -= 1
            if i >= j { break while_loop; }
            best_area = max(best_area, get_area(i, j))
        }

        // both are the same, move both inward.
        if height[i] == height[j] {
            i += 1
            j -= 1
            if i >= j { break while_loop; }
            best_area = max(best_area, get_area(i, j))
        }
    }
    return best_area
}