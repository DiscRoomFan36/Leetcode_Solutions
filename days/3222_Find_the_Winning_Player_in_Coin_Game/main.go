package main

func winningPlayer(x int, y int) string {
    alice_wins := false;

    for x >= 1 && y >= 4 {
        alice_wins = !alice_wins;
        x -= 1;
        y -= 4;
    }

    if alice_wins {
        return "Alice";
    } else {
        return "Bob";
    }
}
