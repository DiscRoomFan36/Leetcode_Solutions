package main

func countCollisions(directions string) int {
    // remove the left ones that dissapear
    for len(directions) > 0 && directions[0                ] == 'L' { directions = directions[1:                 ] }
    for len(directions) > 0 && directions[len(directions)-1] == 'R' { directions = directions[ :len(directions)-1] }

    // all other cars will crash, execpt for 'S' (stright) ones.
    stright_roads := 0
    for _, c := range directions {
        // it should be really fast to count the number of stright roads.
        if c == 'S' { stright_roads += 1 }
    }
    return len(directions) - stright_roads
}