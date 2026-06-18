package main

func angleClock(hour int, minutes int) float64 {
    minute_0_to_1 := float64(minutes) / 60;
    minute_hand := minute_0_to_1 * 360;

    hour_0_to_12 := float64(hour % 12) + minute_0_to_1;
    hour_hand := hour_0_to_12 / 12 * 360;

    return min(Abs(hour_hand - minute_hand), Abs(hour_hand + 360 - minute_hand), Abs(hour_hand - (minute_hand + 360)));
}

func Abs(x float64) float64 {
    if x < 0 { x *= -1; }
    return x;
}
