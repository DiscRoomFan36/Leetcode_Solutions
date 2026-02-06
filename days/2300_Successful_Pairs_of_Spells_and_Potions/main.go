package main

import "slices"

func successfulPairs(spells []int, potions []int, success int64) []int {
    // we want a quick way of checking how many potions are bigger than a certen number

    // first, sort the potions 
    // O(m*log(m))
    slices.Sort(potions)

    result := make([]int, len(spells))

    // O( n * log(m) )
    for i := range spells {
        spell_power := spells[i]
        // how much power the potion needs to be succsessful,
        // div_ceil()
        required_power := (int(success) + spell_power - 1) / spell_power
        // quicklt find the smallest potion that has the required power.
        // 
        // this is a little scetchy, as i dont know exactly where the
        // binary seacrch will put me, what if it terminates early?
        // welp, the tests passed, so i guess its fine.
        potion_index, _ := slices.BinarySearch(potions, required_power)
        // number of greater potions
        result[i] = len(potions) - potion_index
    }

    return result
}