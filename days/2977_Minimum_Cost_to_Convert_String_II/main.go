package main

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
    if source == target { return 0 }
    if len(source) != len(target) { return -1 }

    str_to_id := make(map[string]int)

    m := len(original)
    for i := range m {
        orig, chang := original[i], changed[i]
        if !Contains(str_to_id, orig ) { str_to_id[orig ] = len(str_to_id) }
        if !Contains(str_to_id, chang) { str_to_id[chang] = len(str_to_id) }
    }

    const INF = 1 << 60

    unique_keys := len(str_to_id)

    warshall_grid := make([][]int, unique_keys)
    for i := range unique_keys {
        warshall_grid[i] = make([]int, unique_keys)
        for j := range unique_keys {
            warshall_grid[i][j] = INF
        }
        // diagonals are allready themselves.
        warshall_grid[i][i] = 0
    }

    for i := range m {
        orig, chang, cos := original[i], changed[i], cost[i]

        from_id := str_to_id[orig ]
        to_id   := str_to_id[chang]

        warshall_grid[from_id][to_id] = min(warshall_grid[from_id][to_id], cos)
    }


    // Floyd-Warshall algorithm
    for k := range unique_keys {
        for i := range unique_keys {
            for j := range unique_keys {
                warshall_grid[i][j] = min(warshall_grid[i][j], warshall_grid[i][k] + warshall_grid[k][j])
            }
        }
    }


    // stores all the keys in reverse. for fast lookup.
    reversed_trie := Trie{}
    for k, _ := range str_to_id {
        // store id values here?
        reversed_trie.Add_Word_Reversed(k)
    }


    // hint 3, i couldnt have done it without you.
    n := len(source)

    lowest_cost_for_prefix := make([]int, n+1)
    for i := range lowest_cost_for_prefix { lowest_cost_for_prefix[i] = INF }
    lowest_cost_for_prefix[0] = 0

    for j := 1; j <= n; j++ {
        lowest_cost_for_j := INF
        if source[j-1] == target[j-1] {
            lowest_cost_for_j = min(lowest_cost_for_j, lowest_cost_for_prefix[j-1])
        }

        src__node := &reversed_trie.root
        dest_node := &reversed_trie.root
        for i := j-1; i >= 0; i-- {
            src__node = src__node.Next[source[i] - 'a']
            dest_node = dest_node.Next[target[i] - 'a']

            // end of the line.
            if src__node == nil || dest_node == nil { break }

            if (src__node.num_occurences > 0) && (dest_node.num_occurences > 0) {
                src  := source[i:j]
                dest := target[i:j]

                from, ok1 := str_to_id[src ]
                to,   ok2 := str_to_id[dest]
                // might be better to do both ok checks here,
                // for branch prediction reasons.
                //
                // gives the compiler the leaway to remove the second call to str_to_id,
                // maybe, i dont know how good golang's optimizer is.
                if !ok1 || !ok2 { continue }

                change_cost := warshall_grid[from][to]

                lowest_cost_for_j = min(lowest_cost_for_j, lowest_cost_for_prefix[i] + change_cost)
            }
        }

        lowest_cost_for_prefix[j] = lowest_cost_for_j
    }

    // fmt.Println("lowest_cost_for_prefix:", lowest_cost_for_prefix)

    if lowest_cost_for_prefix[n] == INF { return -1 }
    return int64(lowest_cost_for_prefix[n])
}

func Contains[T comparable, U any](the_map map[T]U, item T) bool {
    _, found := the_map[item]
    return found
}


type Trie_Node struct {
    // character - 'a' to select next.
    Next [26]*Trie_Node

    // gotta check if this node is real.
    num_occurences int
}
type Trie struct {
    root Trie_Node
}

// asumes the word is made of lowercase ascii 
func (trie *Trie) Add_Word_Reversed(word string) {
    node := &trie.root

    for i := len(word)-1; i >= 0; i-- {
        c := word[i]

        if !('a' <= c && c <= 'z') { panic("word contained an invalid character") }
        index := c-'a'

        if node.Next[index] == nil {
            node.Next[index] = new(Trie_Node)
        }
        node = node.Next[index]
    }

    node.num_occurences += 1
}
