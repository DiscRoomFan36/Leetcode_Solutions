package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestNodes(root *TreeNode, queries []int) [][]int {
	if root == nil {
		panic("invalid root node, need 1 element")
	}

	// you motherfuckers, the tree isn't balenced is it? I hate you all.

	// there is a universe where it would be faster to count the number
	// of elements in the root so we can quickly append to it.
	//
	// but that number would take some time, time spent waiting on
	// memory to load random pointers.
	//
	// remember kids, always keep track of how many nodes are in you tree
	// structure, it could be useful for optimizations like these
	sorted_array := make([]int, 0)
	to_sorted_array(root, &sorted_array)

	result := make([][]int, len(queries))
	for i := range queries {
		result[i] = make([]int, 2)
	}

	for i, n := range queries {
		// just a standard binary search
		//
		// not even going to implement it myself
		index, found := slices.BinarySearch(sorted_array, n)

		if found {
			result[i][0] = n
			result[i][1] = n
		} else {
			if index == 0 {
				// off the left edge,
				result[i][0] = -1
				result[i][1] = sorted_array[0]
			} else if index == len(sorted_array) {
				// off the right edge
				result[i][0] = sorted_array[len(sorted_array)-1]
				result[i][1] = -1
			} else {
				result[i][0] = sorted_array[index-1]
				result[i][1] = sorted_array[index]
			}
		}
	}

	return result
}

// assumes the tree is a proper binary tree.
func to_sorted_array(node *TreeNode, result *[]int) {
	if node.Left != nil {
		to_sorted_array(node.Left, result)
	}
	*result = append(*result, node.Val)
	if node.Right != nil {
		to_sorted_array(node.Right, result)
	}
}
