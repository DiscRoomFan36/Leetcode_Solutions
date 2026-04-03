package main

func numberOfWays(corridor string) int {
    // all chairs must pair,
    type Pair_Of_Chairs struct {
        a, b int;
    };

    // find all pairs,
    //
    // we dont actually need this array, could do it inline, i do not care.
    pairs_of_chairs := make([]Pair_Of_Chairs, 0);
    for i := 0; i < len(corridor); i++ {
        if corridor[i] == 'S' {
            first_chair := i;

            // search for pair
            for i += 1; i < len(corridor); i++ {
                if corridor[i] == 'S' { break; }
            }

            if i >= len(corridor) {
                // could not find a pair for this chair. no corridor can be made.
                return 0;
            }
            second_chair := i;

            Append(&pairs_of_chairs, Pair_Of_Chairs{first_chair, second_chair});
        }
    }

    // fmt.Println(pairs_of_chairs);

    // no chairs, no ways to divide.
    if len(pairs_of_chairs) == 0 { return 0; }

    // find their gaps,
    // product of all gaps is answer
    const MOD = 1000000007;
    result := 1;
    for i := range len(pairs_of_chairs)-1 {
        pair_1 := pairs_of_chairs[i  ];
        pair_2 := pairs_of_chairs[i+1];
        number_of_plants_between := pair_2.a - pair_1.b - 1;

        result *= number_of_plants_between+1;
        result %= MOD;
    }
    return result;
}

func Append[T any](slice *[]T, items ...T) {
	*slice = append(*slice, items...);
}
