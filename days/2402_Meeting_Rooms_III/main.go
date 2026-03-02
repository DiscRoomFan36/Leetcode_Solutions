package main

import (
	"cmp"
	"slices"
)

func mostBooked(n int, meetings [][]int) int {
    if n == 0 { panic("bad input."); }

    // i can see it coming a mile away.
    slices.SortFunc(meetings,
        func(a, b []int) int {
            return cmp.Compare(a[0], b[0]);
        },
    )

    meeting_room_time_when_free := make([]int, n);
    meeting_counts              := make([]int, n);

    for i := range meetings {
        start_time := meetings[i][0];
        end_time   := meetings[i][1];

        // find a room,
        next_free_room_number := 0;
        // the first room to become free if we cannot find a room.
        first_time_free := meeting_room_time_when_free[0];

        for room_number, time_free := range meeting_room_time_when_free {
            if time_free <= start_time {
                // found a real free room.
                next_free_room_number = room_number;
                break;
            }

            if time_free < first_time_free {
                first_time_free = time_free;
                next_free_room_number = room_number;
            }
        }

        // meetings always last the same amount of time.
        duration := end_time - start_time;
        real_end_time := max(meeting_room_time_when_free[next_free_room_number], start_time) + duration;

        meeting_room_time_when_free[next_free_room_number] = real_end_time;
        meeting_counts             [next_free_room_number] += 1;
    }

    // fmt.Println(meeting_room_time_when_free, meeting_counts)

    best_count       := 0;
    best_room_number := 0;
    for room_number, count := range meeting_counts {
        if count > best_count {
            best_room_number = room_number;
            best_count       = count;
        }
    }
    return best_room_number;
}