package main

import "slices"

func nextGreatestLetter(letters []byte, target byte) byte {
    index, found := slices.BinarySearch(letters, target)

    if found {
        // go forwards until a different letter
        for i := index+1; i < len(letters); i++ {
            if letters[i] != target { return letters[i] }
        }
        return letters[0] // no larger letter, return base case
    }

    // there are no larger letters
    if index == len(letters) { return letters[0] }

    return letters[index]
}