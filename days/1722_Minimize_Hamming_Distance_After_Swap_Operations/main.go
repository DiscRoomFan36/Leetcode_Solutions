package main

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
    n := len(source);

    // organize the swaps into groups
    group_to_indexes := make(map[int][]int);
    index_to_group := make([]int, n);
    // set all to NO_GROUP
    const NO_GROUP = -1;
    for i := range n { index_to_group[i] = NO_GROUP; }

    // used when creating new groups, to get a unique group id.
    num_groups := 0;

    for _, swap := range allowedSwaps {
        a, b := swap[0], swap[1];
        a_group, b_group := index_to_group[a], index_to_group[b];

        if a_group == NO_GROUP && b_group == NO_GROUP {
            // both have no group.
            new_group := num_groups;
            num_groups += 1;

            index_to_group[a] = new_group;
            index_to_group[b] = new_group;

            group_to_indexes[new_group] = []int{a, b};

        } else if a_group == NO_GROUP || b_group == NO_GROUP {
            // one has a group, the other dose not.
            if a_group == NO_GROUP {
                // swap them, so a has a group.
                a, b = b, a;
                a_group, b_group = b_group, a_group;
            }

            // give b the same group as a,
            index_to_group[b] = a_group;
            group_to_indexes[a_group] = append(group_to_indexes[a_group], b);


        } else {
            // both are in a group.

            // if they are in the same group, we dont need to do anything.
            if a_group == b_group { continue; }

            if len(group_to_indexes[a_group]) < len(group_to_indexes[b_group]) {
                // swap them, so a is bigger
                a, b = b, a;
                a_group, b_group = b_group, a_group;
            }

            // merge groups.
            //
            // this is currently done in a bad n^2 fashion.
            // but is still pretty fast.

            // turn all b's to a's
            for _, i := range group_to_indexes[b_group] {
                if index_to_group[i] != b_group { panic("what"); }

                index_to_group[i] = a_group;
                group_to_indexes[a_group] = append(group_to_indexes[a_group], i);
            }

            // delete the useless group.
            delete(group_to_indexes, b_group);
        }
    }


    hamming_distance := 0;
    // get all the index's that couldn't be swapped, count by themselves.
    for i := range n {
        if index_to_group[i] != NO_GROUP { continue; }

        // there's nothing we can do.
        if source[i] != target[i] {
            hamming_distance += 1;
        }
    }

    // now to do the swapping networks we found.
    for _, indexes := range group_to_indexes {
        // find the difference between how often something occurs in the sorce, vs the taget.
        occurence_counts := make(map[int]int);
        for _, index := range indexes {
            occurence_counts[source[index]] -= 1;
            occurence_counts[target[index]] += 1;
        }

        for _, count := range occurence_counts {
            // if the number is greater than 0, the target has some unfilled goals.
            // if the number is less, thats just the source having the wrong digits.
            if count > 0 {
                hamming_distance += count;
            }
        }
    }

    return hamming_distance;
}
