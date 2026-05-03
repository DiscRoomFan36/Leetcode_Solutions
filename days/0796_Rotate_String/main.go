package main

func rotateString(s string, goal string) bool {
    n := len(s);
    if n != len(goal) { return false; }
    for i := range n {
        can_be_this_rotation := true;
        for j := range n {
            if s[(i + j) % n] != goal[j] {
                can_be_this_rotation = false;
                break;
            }
        }
        if can_be_this_rotation {
            return true;
        }
    }
    return false;
}
