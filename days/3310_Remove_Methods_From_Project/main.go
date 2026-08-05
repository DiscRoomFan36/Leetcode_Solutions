package main

func remainingMethods(n int, k int, invocations [][]int) []int {
    type Node_Info struct {
        invokes []int;
        gets_invoked_by []int;

        is_suspicious bool;
        seen_in_second_check bool;
    }
    nodes := make([]Node_Info, n);

    get_node_info := func(node int) *Node_Info {
        return &nodes[node];
    }

    // populate node info's
    for _, invocation := range invocations {
        a, b := invocation[0], invocation[1];

        a_info := get_node_info(a);
        b_info := get_node_info(b);

        Append(&a_info.invokes, b);
        Append(&b_info.gets_invoked_by, a);
    }

    // mark suspicious
    nodes_to_check := []int{k};
    for len(nodes_to_check) > 0 {
        node := Pop(&nodes_to_check);

        node_info := get_node_info(node);
        if node_info.is_suspicious { continue; }
        node_info.is_suspicious = true;

        Append(&nodes_to_check, node_info.invokes...);
    }

    // check if any in k's group are non sus,
    any_non_sus := false;

    to_check := []int{k};
    for len(to_check) > 0 {
        checking := Pop(&to_check);
        checking_info := get_node_info(checking);

        if checking_info.seen_in_second_check { continue; }
        checking_info.seen_in_second_check = true;

        if !checking_info.is_suspicious {
            // we found one in the group
            any_non_sus = true;
            break;
        } else {
            // keep checking
            Append(&to_check, checking_info.invokes...);
            Append(&to_check, checking_info.gets_invoked_by...);
        }
    }

    result := make([]int, 0);
    if any_non_sus {
        // return 0..n-1
        for i := range n {
            Append(&result, i);
        }
    } else {
        // return all non sus
        for i := range n {
            node_info := get_node_info(i);
            if node_info.is_suspicious { continue; }
            Append(&result, i);
        }
    }
    return result;
}

func Append[T any](slice *[]T, items ...T) *T {
    if len(items) == 0 { return nil; }
	*slice = append(*slice, items...);
    return &(*slice)[len(*slice)-1];
}

func Pop[T any](slice *[]T) T {
	item  := (*slice)[ len(*slice)-1];
	*slice = (*slice)[:len(*slice)-1];
	return item;
}
