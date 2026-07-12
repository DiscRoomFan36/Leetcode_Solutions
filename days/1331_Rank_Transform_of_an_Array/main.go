package main

import (
	"cmp"
	"slices"
)

func arrayRankTransform(arr []int) []int {
    if len(arr) == 0 { return []int{}; }

    type Num_And_Index struct {
        num, index int;
    };

    num_and_index := make([]Num_And_Index, len(arr));
    for i, n := range arr { num_and_index[i] = Num_And_Index{n, i}; }

    slices.SortFunc(num_and_index,
        func (a, b Num_And_Index) int {
            return cmp.Compare(a.num, b.num);
        },
    );

    result := make([]int, len(arr));

    rank := 1;
    result[num_and_index[0].index] = rank;

    for i := 1; i < len(num_and_index); i++ {
        last := num_and_index[i-1];
        this := num_and_index[i  ];

        if last.num != this.num { rank += 1; }
        result[this.index] = rank;
    }

    return result;
}
