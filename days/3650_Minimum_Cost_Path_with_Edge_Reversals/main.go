package main

import "container/heap"

func minCost(n int, _edges [][]int) int {
    type Where_And_Weight struct { where, weight int }


    // a sparse map from an edges to all its connections.
    from_to_where_with_weight := make(map[int][]Where_And_Weight)
    for i := range _edges {
        u, v, w := _edges[i][0], _edges[i][1], _edges[i][2]

        add_to_where_and_weights := func(from_to_where_with_weight map[int][]Where_And_Weight, u, v, w int) {
            arr := from_to_where_with_weight[u]
            for i := range arr {
                if arr[i].where == v {
                    arr[i].weight = min(arr[i].weight, w)
                    return
                }
            }

            from_to_where_with_weight[u] = append(from_to_where_with_weight[u], Where_And_Weight{where: v, weight: w})
        }

        add_to_where_and_weights(from_to_where_with_weight, u, v, w)
        add_to_where_and_weights(from_to_where_with_weight, v, u, w*2)
    }


    // dikstras algo

    // also this is some wack golang thing.
    pq := PriorityQueue{
        arr: []int{ 0 },
        city_weight: map[int]int{ 0 : 0 },
    }
    heap.Init(&pq)

    processed_before := make(map[int]bool)

    for pq.Len() > 0 {
        current_city := heap.Pop(&pq).(int)

        // if we are checking some edge, that means weve allready found the shortest path to it.
        if current_city == n-1 { return pq.city_weight[current_city] }

        // some method of not checking the same city twice.
        if processed_before[current_city] { continue }
        processed_before[current_city] = true


        // check all neighboring routes
        for _, w_and_w := range from_to_where_with_weight[current_city] {
            to_where, edge_weight := w_and_w.where, w_and_w.weight

            new_city_weight := edge_weight + pq.city_weight[current_city]

            // check if weve reached that city before.
            old_city_weight, seen_before := pq.city_weight[to_where]

            if (!seen_before) || (new_city_weight < old_city_weight) {
                pq.city_weight[to_where] = new_city_weight

                // add the new city onto the heap
                heap.Push(&pq, to_where )
            }
        }
    }

    return -1
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue struct {
    arr []int
    // city to least distance
    city_weight map[int]int
}

func (pq PriorityQueue) Len() int { return len(pq.arr) }
func (pq PriorityQueue) Less(i, j int) bool {
    i_w := pq.city_weight[pq.arr[i]]
    j_w := pq.city_weight[pq.arr[j]]
    // we want the lowest weight
    return i_w < j_w
}
func (pq PriorityQueue) Swap(i, j int) {
    pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
}


func (pq *PriorityQueue) Push(x any) {
	item := x.(int)
	pq.arr = append(pq.arr, item)
}

func (pq *PriorityQueue) Pop() any {
	old := pq.arr
	n := len(old)
	item := old[n-1]
	pq.arr = old[0 : n-1]
	return item
}
