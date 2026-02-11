package main

func pacificAtlantic(heights [][]int) [][]int {
    if len(heights) == 0 { return [][]int{} }
    m, n := len(heights), len(heights[0])

    type Item struct { x, y int }

    hill_climb_with_starting_state := func(stack []Item) [][]bool {
        result := make_grid[bool](n, m)

        for len(stack) > 0 {
            item := Pop(&stack)
            x, y := item.x, item.y

            if result[y][x] { continue }
            result[y][x] = true

            this_height := heights[y][x]

            if 0 < x   && this_height <= heights[y][x-1] {
                Append(&stack, Item{x-1, y})
            }
            if x < n-1 && this_height <= heights[y][x+1] {
                Append(&stack, Item{x+1, y})
            }

            if 0 < y   && this_height <= heights[y-1][x] {
                Append(&stack, Item{x, y-1})
            }
            if y < m-1 && this_height <= heights[y+1][x] {
                Append(&stack, Item{x, y+1})
            }
        }

        return result
    }


    stack := []Item{}

    // cells in contact with pacific
    for i := range n { Append(&stack, Item{  i,   0}) }
    for j := range m { Append(&stack, Item{  0,   j}) }
    can_flow_pacific := hill_climb_with_starting_state(stack)

    // clear the stack
    stack = stack[:0]
    // cells in contact with atlantic
    for i := range n { Append(&stack, Item{  i, m-1}) }
    for j := range m { Append(&stack, Item{n-1,   j}) }
    can_flow_atlantic := hill_climb_with_starting_state(stack)


    result := [][]int{}
    for j := range m {
        for i := range n {
            if can_flow_pacific[j][i] && can_flow_atlantic[j][i] {
                Append(&result, []int{j, i})
            }
        }
    }
    return result
}


func Append[T any](slice *[]T, items ...T) {
	*slice = append(*slice, items...)
}

func Pop[T any](slice *[]T) T {
    item  := (*slice)[ len(*slice)-1]
    *slice = (*slice)[:len(*slice)-1]
    return item
}

// makes a grid, with default values. i hope go can turn this into something cool.
//
// might be better to make([]T, m*n) then grab some slices.
func make_grid[T any](n, m int) [][]T {
    grid := make([][]T, m)
    for j := range m { grid[j] = make([]T, n) }
    return grid
}