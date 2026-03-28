package main

func numberOfBeams(bank []string) int {
    last_count_of_beams := 0;
    result := 0;

    for _, beams := range bank {
        beam_count := 0;
        for _, c := range beams {
            if c == '1' { beam_count += 1; }
        }
        if beam_count == 0 { continue; }

        result += last_count_of_beams * beam_count;
        last_count_of_beams = beam_count;
    }
    return result;
}
