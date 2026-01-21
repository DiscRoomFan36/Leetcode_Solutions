package main

func splitArray(nums []int) int64 {
	n := len(nums)

	is_not_prime := sieve_of_Eratosthenes(n)

	sum_a := 0
	sum_b := 0

	for i := range n {
		if is_not_prime[i] {
			sum_b += nums[i]
		} else {
			sum_a += nums[i]
		}
	}

	result := sum_a - sum_b
	if result < 0 {
		result = -result
	}
	return int64(result)
}

// returns an array of bools with index i true if it is not a prime.
func sieve_of_Eratosthenes(n int) []bool {
	is_not_prime := make([]bool, n)

	if n <= 0 {
		return is_not_prime
	}
	is_not_prime[0] = true
	if n <= 1 {
		return is_not_prime
	}
	is_not_prime[1] = true

	// do all evens for speed
	for i := 2 * 2; i < n; i += 2 {
		is_not_prime[i] = true
	}

	curr_prime := 3
	for curr_prime < n {
		// find the next prime.
		for (curr_prime < n) && is_not_prime[curr_prime] {
			curr_prime += 1
		}

		// we know all evens are already done. so skip faster.
		for i := curr_prime * 3; i < n; i += curr_prime * 2 {
			is_not_prime[i] = true
		}
		curr_prime += 1
	}

	return is_not_prime
}
