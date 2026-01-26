package main

func maximumEnergy(energy []int, k int) int {
    if len(energy) == 0 { return 0 }
    // they would jump to their own position, or before, and we dont deal with that
    if k <= 0 { return 0 }

    max_energy := energy[len(energy)-1]

    energy_at := make([]int, len(energy))
    for i := len(energy)-1; i >= 0; i-- {
        energy_at[i] = energy[i]
        if i + k < len(energy) { energy_at[i] += energy_at[i+k] }

        max_energy = max(max_energy, energy_at[i])
    }

    return max_energy
}