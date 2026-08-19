package main

import (
	"cmp"
	"slices"
)

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
    
    slices.SortFunc(reservedSeats,
        func(a, b []int) int {
            a_row, _ := a[0], a[1];
            b_row, _ := b[0], b[1];
            return cmp.Compare(a_row, b_row);
        },
    );
    
    // fmt.Println(reservedSeats);
    total_groups_to_seat := 0;

    calc_max_seats_rows_between := func(low, high int) int {
        missed_rows := high - low;
        return missed_rows * 2;
    }

    last_row := 1;
    for i := 0; i < len(reservedSeats); {
        this_row := reservedSeats[i][0];

        // calc how many rows are empty.
        total_groups_to_seat += calc_max_seats_rows_between(last_row, this_row);
        last_row = this_row + 1;

        // get all the seats in the same row.
        end := i+1;
        for ; end < len(reservedSeats); end += 1 {
            end_row := reservedSeats[end][0];
            if end_row != this_row {
                break;
            }
        }

        reserved_seats_in_row := reservedSeats[i:end];
        reserved := [10+1]bool{};
        for _, reservation := range reserved_seats_in_row {
            row, seat := reservation[0], reservation[1];
            if row != this_row { panic("what"); }
            reserved[seat] = true;
        }

        range_can_fit := func(low, high int) bool {
            for i := low; i <= high; i += 1 {
                if reserved[i] { return false; }
            }
            return true;
        }

        to_seat := 0;
        if range_can_fit(2, 5) {
            // check if second can fit
            if range_can_fit(6, 9) {
                to_seat = 2;
            } else {
                to_seat = 1;
            }
        } else {
            if range_can_fit(4, 7) {
                to_seat = 1;
            } else if range_can_fit(6, 9) {
                to_seat = 1;
            }
        }

        // fmt.Println(this_row, to_seat);
        total_groups_to_seat += to_seat;
        i = end;
    }
    // final rows that were empty
    total_groups_to_seat += calc_max_seats_rows_between(last_row, n+1);

    return total_groups_to_seat;
}
