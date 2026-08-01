package main

import "math"

func predictTheWinner(nums []int) bool {
    type State struct {
        nums []int;
        player_1_score, player_2_score int;
        maximizing_player_1 bool;
    }

    bool_to_int := func(b bool) int {
        if b   { return 1;
        } else { return 0; }
    }

    var recur func(state State, alpha, beta int) int;
    recur = func(state State, alpha, beta int) int {
        if len(state.nums) == 0 {
            player_1_wins := state.player_1_score >= state.player_2_score;
            return bool_to_int(player_1_wins);
        }

        if state.maximizing_player_1 {
            states := []State{
                {state.nums[1:],                 state.player_1_score + state.nums[0                ], state.player_2_score, false},
                {state.nums[:len(state.nums)-1], state.player_1_score + state.nums[len(state.nums)-1], state.player_2_score, false},
            };

            max_eval := math.MinInt;
            for _, new_state := range states {
                eval := recur(new_state, alpha, beta);
                max_eval = max(max_eval, eval);
                alpha = max(alpha, eval);
                if beta <= alpha { break; }
            }

            return max_eval;
        } else {
            states := []State{
                {state.nums[1:],                 state.player_1_score, state.player_2_score + state.nums[0                ], true},
                {state.nums[:len(state.nums)-1], state.player_1_score, state.player_2_score + state.nums[len(state.nums)-1], true},
            };

            min_eval := math.MaxInt;
            for _, new_state := range states {
                eval := recur(new_state, alpha, beta);
                min_eval = min(min_eval, eval);
                beta = min(beta, eval);
                if beta <= alpha { break; }
            }
            return min_eval;
        }
    }

    result := recur(State{nums, 0, 0, true}, math.MinInt, math.MaxInt);
    return result == 1;
}
