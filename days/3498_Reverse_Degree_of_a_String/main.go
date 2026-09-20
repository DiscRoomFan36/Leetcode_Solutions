package main

func reverseDegree(s string) int {
    sum := 0;
    for i, c := range s {
        score := (i + 1) * int(('z' - 'a') - (c - 'a') + 1);
        sum += score;
    }
    return sum;
}
