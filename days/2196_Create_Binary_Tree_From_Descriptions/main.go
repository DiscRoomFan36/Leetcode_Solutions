package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
    type Node_Info struct {
        the_node *TreeNode;
        has_a_parent bool;
    };

    value_to_node := make(map[int]Node_Info);

    get_node := func(value int, is_child bool) *TreeNode {
        node_info, found := value_to_node[value];

        if !found {
            node_info.the_node = new(TreeNode);
            node_info.the_node.Val = value;
            node_info.has_a_parent = false;
        }

        if is_child { node_info.has_a_parent = true; }

        value_to_node[value] = node_info;

        return node_info.the_node;
    }

    for _, description := range descriptions {
        parent, child, is_left := description[0], description[1], description[2];

        parent_node := get_node(parent, false);
        child_node  := get_node( child,  true);

        if is_left == 1 {
            parent_node.Left = child_node;
        } else {
            parent_node.Right = child_node;
        }
    }

    // find the root.
    //
    // root is the only node that is not a child to some other node.
    for _, node_info := range value_to_node {
        fmt.Println("node_info:", node_info);
        if !node_info.has_a_parent {
            return node_info.the_node;
        }
    }
    panic("no root node found");
}
