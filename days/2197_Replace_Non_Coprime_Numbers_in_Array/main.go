package main

import "unsafe"

func replaceNonCoprimes(nums []int) []int {
    result := []uint{};

    for _, n := range nums {
        // add a new number onto the end.
        result = append(result, uint(n));

        // fold until you can fold no more
        for len(result) != 1 {
            a, b := result[len(result)-2], result[len(result)-1];
            this_gcd := gcd(a, b);
            // check if we need to fold this down.
            if this_gcd <= 1 { break; }

            // combine the last numbers
            lcm := a * (b / this_gcd);
            result[len(result)-2] = lcm;
            result = result[:len(result)-1];
        }
    }

    // play stupid games, get stupid solutions.
    return Unsafe_Slice_Transmute[uint, int](result);
}

// https://en.wikipedia.org/wiki/Greatest_common_divisor#Binary_GCD_algorithm
func gcd(a, b uint) uint {
    if a == 0 { return b }
    if b == 0 { return a }

    // find shared powers of 2
    d := 0;
    for (a % 2 == 0) && (b % 2 == 0) {
        a /= 2;
        b /= 2;
        d += 1;
    }

    // make both odd
    for a % 2 == 0 { a /= 2; }
    for b % 2 == 0 { b /= 2; }

    // until a == b
    for a != b {
        if a < b { a, b = b, a; }

        a = (a - b) / 2;
        for a % 2 == 0 { a /= 2; }
    }

    return (1 << d) * a;
}


func Unsafe_Slice_Transmute[T any, U any](slice []T) []U {
	var x T; var y U
	T_size := unsafe.Sizeof(x)
	U_size := unsafe.Sizeof(y)

	flag := false
	if U_size >= T_size { flag = (U_size % T_size != 0)
	} else {              flag = (T_size % U_size != 0) }
	if flag { panic("T and U must have sizes that are multiples of each other") }

	data := unsafe.Pointer(unsafe.SliceData(slice))
	size := uintptr(len(slice)) * T_size / U_size
	return unsafe.Slice((*U)(data), size)
}
