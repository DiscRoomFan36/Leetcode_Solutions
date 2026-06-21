package main

func maxIceCream(costs []int, coins int) int {
    max_cost := 0;
    for _, cost := range costs { max_cost = max(max_cost, cost); }

    counts := make([]int, max_cost);
    for _, cost := range costs { counts[cost-1] += 1; }

    total_can_buy := 0;
    for cost_minus_1, count := range counts {
        cost := cost_minus_1 + 1;

        // if the count is zero, we dont care
        if count == 0 { continue; }

        if cost * count < coins {
            coins -= cost * count;
            total_can_buy += count;
        } else {
            // get how many we can buy, then exit.
            total_can_buy += coins / cost;
            break;
        }
    }

    return total_can_buy;
}
