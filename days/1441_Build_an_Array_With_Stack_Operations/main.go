package main

func buildArray(target []int, n int) []string {
    if len(target) == 0 { return []string{} }
    const PUSH = "Push"
    const POP  = "Pop"

    // every number needs 2, except the ones all-ready in the array, they need 1.
    total_amount_of_elements := target[len(target)-1]*2 - len(target)
    result := make([]string, 0, total_amount_of_elements)

    stream := 1

    for _, next_target := range target {
        for stream < next_target {
            result = append(result, PUSH, POP)
            stream += 1
        }
        result = append(result, PUSH)
        stream += 1
    }

    return result
}
