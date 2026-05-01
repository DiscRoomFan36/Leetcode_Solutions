package main

func findPrimePairs(n int) [][]int {
    primes := primes_upto(n);

    // fmt.Println(primes);

    // for fast prime check.
    prime_set := make(map[int]struct{});
    for _, p := range primes { prime_set[p] = struct{}{}; }

    result := [][]int{};

    for _, p := range primes {
        needed_pair := n - p;
        // p must be bigger than the pair.
        if needed_pair < p { break; }

        if _, ok := prime_set[needed_pair]; ok {
            result = append(result, []int{p, needed_pair});
        }
    }

    return result;
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
