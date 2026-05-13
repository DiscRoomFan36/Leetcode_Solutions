package main

func sumOfPrimesInRange(n int) int {
    reverse := Reversed(n);

    // make it so n is always smaller.
    if reverse < n { n, reverse = reverse, n; }

    is_not_prime := sieve_of_Eratosthenes(reverse+1);

    sum_of_primes := 0;
    for i := n; i <= reverse; i++ {
        if is_not_prime[i] { continue; }
        sum_of_primes += i;
    }
    return sum_of_primes;
}

func Reversed(n int) int {
    reverse := 0;
    for n > 0 {
        reverse = reverse * 10 + n % 10;
        n /= 10;
    }
    return reverse;
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
