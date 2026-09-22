package main

func gardenNoAdj(n int, paths [][]int) []int {
    node_to_paths := make([][]int, n);
    for _, path := range paths {
        x, y := path[0]-1, path[1]-1;
        node_to_paths[x] = append(node_to_paths[x], y);
        node_to_paths[y] = append(node_to_paths[y], x);
    }

    colors := make([]int, n);

    // greedily give nodes lowest color possible.
    for node := range n {

        // the loop is maximumly 4*3 == 12, so its fine.
        chosen_color := 0;
        for color := 1; color <= 4; color += 1 {
            good_color := true;
            for _, other_node := range node_to_paths[node] {
                if colors[other_node] == color {
                    good_color = false;
                    break;
                }
            }
            if good_color {
                chosen_color = color;
                break;
            }
        }

        if chosen_color == 0 { panic("greedy solution cannot win."); }
        colors[node] = chosen_color;
    }

    return colors;
}
