package main

func finalValueAfterOperations(operations []string) int {
    X := 0;
    for _, op := range operations {
        switch op {
            case "++X", "X++": { X += 1; }
            case "--X", "X--": { X -= 1; }
        }
    }
    return X;
}