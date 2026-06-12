package main

func finalPrices(prices []int) []int {
    answer := make([]int, len(prices));
    for i := range prices {
        discount := 0;
        for j := i+1; j < len(prices); j++ {
            if prices[j] <= prices[i] {
                discount = prices[j];
                break;
            }
        }
        answer[i] = prices[i] - discount;
    }
    return answer;
}
