package main

func maxPalindromes(s string, k int) int {
    answer := 0;
    start := 0;
    for r := k - 1; r < len(s); r++ {
        l := r - k + 1;
        if l >= start && is_palindrom(s[l:r+1]) {
            answer++;
            start = r + 1;
            continue;
        }

        l = r - k;
        if l >= start && is_palindrom(s[l:r+1]) {
            answer++;
            start = r + 1;
        }
    }
	return answer;
}

func is_palindrom(s string) bool {
    for i := range len(s)/2 {
        j := len(s)-1-i;
        if s[i] != s[j] { return false; }
    }
    return true;
}
