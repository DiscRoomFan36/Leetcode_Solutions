package main

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
    type Rectangle struct { x1, y1, x2, y2 int; };
    a := Rectangle{ rec1[0], rec1[1], rec1[2], rec1[3], };
    b := Rectangle{ rec2[0], rec2[1], rec2[2], rec2[3], };

    return (a.x2 > b.x1) && (a.x1 < b.x2) && (a.y2 > b.y1) && (a.y1 < b.y2);
}
