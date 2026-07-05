package main

func pathsWithMaxScore(board []string) []int {
    m, n := len(board), len(board[0]);

    type Score_And_Num struct {
        score, num_ways int;
    }

    line_1 := make([]Score_And_Num, n);
    line_2 := make([]Score_And_Num, n);

    line_1[n-1].num_ways = 1;
    for i := n-1-1; i >= 0; i-- {
        c := board[m-1][i];
        if c == 'X' { break; }
        if c == 'E' { panic("what!"); }

        line_1[i].score = int(c - '0') + line_1[i+1].score;
        line_1[i].num_ways = line_1[i+1].num_ways;
    }

    const MOD = 1000000007;

    for j := m-1-1; j >= 0; j-- {
        // swap buffers
        line_1, line_2 = line_2, line_1;

        for i := n-1; i >= 0; i-- {
            c := board[j][i];
            // block, cant continue.
            if c == 'X' { continue; }

            // clear the area.
            line_1[i] = Score_And_Num{};

            get_next_best := func(a, b, c Score_And_Num) Score_And_Num {
                best_score := max(a.score, b.score, c.score);



                num_ways := 0;
                add_to_num_ways := func(x Score_And_Num) {
                    if x.score == best_score {
                        if num_ways == 0 { num_ways  = x.num_ways;
                        } else           {
                            num_ways = (num_ways + x.num_ways) % MOD;
                        }
                    }
                }

                add_to_num_ways(a);
                add_to_num_ways(b);
                add_to_num_ways(c);

                if num_ways == 0 { best_score = 0; }

                return Score_And_Num{best_score, num_ways};
            }

            up       := line_2[i];
            right    := Score_And_Num{};
            up_right := Score_And_Num{};

            if i != n-1 {
                right    = line_1[i+1];
                up_right = line_2[i+1];
            }

            best_score_and_path := get_next_best(up, right, up_right);

            if c == 'E' {
                // the end!
                return []int{best_score_and_path.score, best_score_and_path.num_ways}
            }

            if best_score_and_path.num_ways != 0 {
                best_score_and_path.score += int(c - '0');
            }
            line_1[i] = best_score_and_path;
        }
    }

    panic("not possible");
}
