package main

import (
	"cmp"
	"slices"
)

func maximumWeight(_intervals [][]int) []int {
    type Interval struct {
        // index in original array
        index int;
        left, right, weight int;
    }
    intervals := make([]Interval, len(_intervals));

    for i := range _intervals {
        intervals[i] = Interval{
            index: i,
            left  : _intervals[i][0],
            right : _intervals[i][1],
            weight: _intervals[i][2],
        };
    }

    slices.SortFunc(intervals,
        func(a, b Interval) int {
            return cmp.Compare(a.right, b.right);
        },
    );

    n := len(intervals);

    type Score_And_Answer struct {
        score int;
        answer []int;
    }

    dp := make([][4+1]Score_And_Answer, n+1);

    for i := range n {
        interval := intervals[i];
        k, _ := slices.BinarySearchFunc(intervals, interval.left,
            func(interval Interval, left int) int {
                return cmp.Compare(interval.right, left);
            },
        );

        for j := 1; j < 4+1; j += 1 {
            s1 := dp[i][j  ].score;
            s2 := dp[k][j-1].score + interval.weight;
            if s1 > s2 {
                dp[i+1][j] = dp[i][j];
                continue;
            }

            // Doing this append is nesseary, because we are sorting them right after, and go isnt that smart.
            new_answer := append([]int{}, dp[k][j-1].answer...);
            new_answer = append(new_answer, interval.index);
            slices.Sort(new_answer);
            if s1 == s2 && slices.Compare(dp[i][j].answer, new_answer) < 0 {
                new_answer = append([]int{}, dp[i][j].answer...);
            }
            dp[i+1][j] = Score_And_Answer{
                s2,
                new_answer,
            }
        }
    }

    return dp[n][4].answer;
}
