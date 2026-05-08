package main

func minJumps(nums []int) int {
    n := len(nums);

    max_number := 0;
    number_to_indexes := make(map[int][]int);
    for i, n := range nums {
        max_number = max(max_number, n);
        number_to_indexes[n] = append(number_to_indexes[n], i);
    }

    // these numbers go upto 10^6, hope thats fine...
    is_not_prime := sieve_of_Eratosthenes(max_number+1);
    // map from prime to positions it can go.
    prime_to_indexes := make(map[int][]int);
    for prime, is_not_prime := range is_not_prime {
        if is_not_prime { continue; }

        indexes := make([]int, 0);
        for num := prime; num <= max_number; num += prime {
            for _, index := range number_to_indexes[num] {
                Append(&indexes, index);
            }
        }

        prime_to_indexes[prime] = indexes;
    }

    curr_positions := make([]int, 0);
    next_positions := make([]int, 0);

    Append(&curr_positions, 0);

    considered := make([]bool, n);
    considered[0] = true;

    prime_considered := make(map[int]bool);

    num_steps := 0;

    for {
        for _, pos := range curr_positions {
            // generate next positions

            if pos == n-1 { return num_steps; } // end condition.

            handle_next_pos := func(next_pos int) {
                if considered[next_pos] { return; }
                considered[next_pos] = true;

                // could check for n-1 here, probably not worth it.

                Append(&next_positions, next_pos);
            }

            // jump forwards or backwords.
            if pos > 0   { handle_next_pos(pos-1); }
            if pos < n-1 { handle_next_pos(pos+1); }

            // do prime teleportation.
            prime := nums[pos]
            // dont check the same prime more than once.
            if prime_considered[prime] { continue; }
            prime_considered[prime] = true;

            if is_not_prime[prime] { continue; } // check if its a prime first.
            for _, j := range prime_to_indexes[prime] {
                handle_next_pos(j);
            }
        }

        // we just did one step.
        num_steps += 1;

        curr_positions, next_positions = next_positions, curr_positions;
        // clear next_positions in time for next loop.
        next_positions = next_positions[:0];
    }
}

func primes_upto(n int) []int {
    bools := sieve_of_Eratosthenes(n);
    result := []int{};
    for i, b := range bools {
        if !b { result = append(result, i); }
    }
    return result;
}

// returns an array of bools with index i true if it is not a prime.
func sieve_of_Eratosthenes(n int) []bool {
	is_not_prime := make([]bool, n);

	if n <= 0 { return is_not_prime; }
	is_not_prime[0] = true;

	if n <= 1 { return is_not_prime; }
	is_not_prime[1] = true;

	// do all evens for speed
	for i := 2 * 2; i < n; i += 2 { is_not_prime[i] = true; }

	curr_prime := 3;
	for curr_prime < n {
		// find the next prime.
		for (curr_prime < n) && is_not_prime[curr_prime] {
			curr_prime += 1;
		}

		// we know all evens are already done. so skip faster.
		for i := curr_prime * curr_prime; i < n; i += curr_prime * 2 {
			is_not_prime[i] = true;
		}
		curr_prime += 1;
	}

	return is_not_prime;
}

func Append[T any](slice *[]T, items ...T) {
	*slice = append(*slice, items...);
}
