package main

func loudAndRich(richer [][]int, quiet []int) []int {
    person_to_richer_people := make(map[int][]int);
    for _, rich := range richer {
        more, less := rich[0], rich[1];
        person_to_richer_people[less] = append(person_to_richer_people[less], more);
    }

    answer := make([]int, len(quiet));
    // initalize to invalid number
    for i := range answer { answer[i] = -1; }

    for i := range answer {
        // person to their answer, recursively and cache'd.
        var recur func(int) int;
        recur = func(person int) int {
            if answer[person] == -1 {
                quietest_person := person;
                for _, richer_person := range person_to_richer_people[person] {
                    richest_person := recur(richer_person);
                    if quiet[richest_person] < quiet[quietest_person] {
                        quietest_person = richest_person;
                    }
                }
                answer[person] = quietest_person;
            }

            return answer[person];
        }

        recur(i);
    }
    return answer;
}
