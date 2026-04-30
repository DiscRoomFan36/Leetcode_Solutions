package main

func findPeaks(mountain []int) []int {
    result := []int{};
    for i := 1; i < len(mountain)-1; i++ {
        if mountain[i-1] < mountain[i] && mountain[i+1] < mountain[i] {
            result = append(result, i);
        }
    }
    return result;
}
