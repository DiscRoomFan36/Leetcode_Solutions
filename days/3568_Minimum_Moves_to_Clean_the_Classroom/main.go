package main

func minMoves(classroom []string, energy int) int {
    m, n := len(classroom), len(classroom[0]);

    type Vec2 struct { x, y int; };

    starting_position := Vec2{};
    // litter_location to bit in bitmask.
    litter_locations  := make(map[Vec2]int);
    full_litter_bitmask := 0;

    for row, row_string := range classroom {
        for col, col_char := range row_string {
            if col_char == 'S' {
                starting_position = Vec2{col, row};
            }

            // enter litter into map.
            if col_char == 'L' {
                shift := len(litter_locations);
                bit   := 1 << shift;

                litter_locations[Vec2{col, row}] = bit;
                full_litter_bitmask |= bit;
            }
        }
    }

    // handle edge case where there is no litter.
    if full_litter_bitmask == 0 { return 0; }

    type State struct {
        pos Vec2;
        // what litter we have collected
        mask int;
        energy int;
    }

    // for the cashe. 
    type State_No_Energy struct {
        pos Vec2;
        mask int;
    }

    // returns best energy had at that position.
    state_cashe := make(map[State_No_Energy]int);

    curr_stack := make([]State, 0);
    next_stack := make([]State, 0);

    Append(&curr_stack, State{pos: starting_position, mask: 0, energy: energy});
    steps := 0;

    directions := []Vec2{
        { 1,  0},
        {-1,  0},
        { 0,  1},
        { 0, -1},
    };

    for len(curr_stack) > 0 {
        steps += 1;

        for len(curr_stack) > 0 {
            state := Pop(&curr_stack);

            // check if we ran out of energy and can go no further.
            //
            // duplicate check for energy == 0,
            // but im thinking of the case where max_energy == 0.
            if state.energy <= 0 { continue; }

            for _, direction := range directions {
                pos := Vec2{state.pos.x + direction.x, state.pos.y + direction.y};
                new_state := State{
                    pos: pos,
                    mask: state.mask,
                    energy: state.energy-1,
                };

                // bounds check new position
                if !Is_Between(pos.x, 0, n-1) { continue; }
                if !Is_Between(pos.y, 0, m-1) { continue; }

                tile := classroom[pos.y][pos.x]
                // cannot go into walls
                if tile == 'X' { continue; }
                // pickup new litter, maybe
                if tile == 'L' {
                    new_state.mask |= litter_locations[Vec2{pos.x, pos.y}];
                    // win condition.
                    if new_state.mask == full_litter_bitmask { return steps; }
                }
                // recharge
                if tile == 'R' { new_state.energy = energy; }

                // we can go no further, and we didnt win already.
                if new_state.energy <= 0 { continue; }

                // check if weve been here before.
                new_state_no_energy := State_No_Energy{pos: new_state.pos, mask: new_state.mask};
                // if we were once in a better position, continue.
                if state_cashe[new_state_no_energy] >= new_state.energy { continue; }
                state_cashe[new_state_no_energy] = new_state.energy;

                // add new state to next_stack
                Append(&next_stack, new_state);
            }
        }

        curr_stack, next_stack = next_stack, curr_stack;
    }

    // found no way to clean all litter.
    return -1;
}


func Append[T any](slice *[]T, items ...T) *T {
	*slice = append(*slice, items...);
    return &(*slice)[len(*slice)-1];
}

func Pop[T any](slice *[]T) T {
	item  := (*slice)[ len(*slice)-1];
	*slice = (*slice)[:len(*slice)-1];
	return item;
}

func Is_Between(x, low, high int) bool {
    return low <= x && x <= high;
}
