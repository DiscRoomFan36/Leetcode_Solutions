package main

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
    hit, _, _ := circle_rectangle_collision(xCenter, yCenter, radius, x1, y1, x2 - x1, y2 - y1);
    return hit;
}

// returns if there was a collision, and the closest point on the rectangle.
func circle_rectangle_collision[T Number](cx, cy, cr T, rx, ry, rw, rh T) (bool, T, T) {
    // things with negative width don't exist.
    if rw < 0 || rh < 0 { return false, rx, ry; }

    rectangle := make_rectangle(rx, ry, rw, rh);
    circle    := make_vec2(cx, cy);

    // point on rectangle that is closest to the circle point.
    testing_point := closest_point_on_rectangle_to_point(rectangle, circle);

    return DistSqr(circle, testing_point) <= cr * cr, testing_point.x, testing_point.y;
}

func closest_point_on_rectangle_to_point[T Number](rec Rectangle[T], point Vec2[T]) Vec2[T] {
    result := make_vec2(
        Clamp(point.x, rec.x, rec.x + rec.w),
        Clamp(point.y, rec.y, rec.y + rec.h),
    );
    return result;
}

type Int    interface { ~int | ~uint | ~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64 };
type Float  interface { ~float32 | ~float64 };
type Number interface { Int | Float };

type Vec2[T Number]      struct { x, y T; }
type Rectangle[T Number] struct { x, y, w, h T; }

func make_vec2[T Number](x, y T) Vec2[T] { return Vec2[T]{x, y}; }

func make_rectangle[T Number](x, y, w, h T) Rectangle[T] {
    return Rectangle[T]{x: x, y: y, w: w, h: h};
}



func Clamp[T Number](x, mini, maxi T) T {
    // TODO remove this?
    if mini > maxi { panic("mini was bigger than maxi"); }

    if x < mini { return mini; }
    if x > maxi { return maxi; }
    return x;
}

// Distance Squared
func DistSqr[T Number](a, b Vec2[T]) T {
    x, y := (a.x - b.x), (a.y - b.y);
    return x*x + y*y;
}
