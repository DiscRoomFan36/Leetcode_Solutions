package main

import (
	"cmp"
	"slices"
)

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
    const X    = 0;
    const Y    = 1;
    const TIME = 2;

    // sort meetings by time.
    slices.SortFunc(meetings,
        func(a, b []int) int {
            return cmp.Compare(a[TIME], b[TIME]);
        },
    )


    // fmt.Println(meetings);

    know_secret := make([]bool, n);
    know_secret[0] = true;
    know_secret[firstPerson] = true;

    // group meetings by time.
    start_of_group := 0;
    for start_of_group < len(meetings) {
        group_meeting_time := meetings[start_of_group][TIME];

        end_of_group := start_of_group;
        for i := start_of_group+1; i < len(meetings); i++ {
            if meetings[i][TIME] == group_meeting_time {
                end_of_group = i;
            } else {
                break;
            }
        }

        meetings_group := meetings[start_of_group:end_of_group+1];

        type Set_And_Know_Secret struct {
            set map[int]struct{};
            know_the_secret bool;
        }

        array_of_sets := make([]Set_And_Know_Secret, 0);
        for _, meeting := range meetings_group {
            x, y := meeting[X], meeting[Y];

            know_the_secret := know_secret[x] || know_secret[y];

            x_location := -1;
            y_location := -1;
            for i := range array_of_sets {
                if Contains(array_of_sets[i].set, x) {
                    x_location = i;
                }
                if Contains(array_of_sets[i].set, y) {
                    y_location = i;
                }
            }

            switch true {
                case x_location == -1 && y_location == -1: {
                    // create new set.
                    new_set_and_know_secret := Set_And_Know_Secret{
                        set: map[int]struct{}{
                            x: {}, y: {},
                        },
                        know_the_secret: know_the_secret,
                    };

                    array_of_sets = append(array_of_sets, new_set_and_know_secret);
                }

                case x_location == -1 || y_location == -1: {
                    // one of them is in a set. add to other set.

                    // make it so x is the known set.
                    if x_location == -1 {
                        x, y = y, x;
                        x_location, y_location = y_location, x_location;
                    }

                    array_of_sets[x_location].set[y] = struct{}{};
                    array_of_sets[x_location].know_the_secret = array_of_sets[x_location].know_the_secret || know_the_secret;
                }

                case x_location != -1 && y_location != -1: {
                    // both are in a group,
                    if x_location != y_location {
                        // merge the groups.
                        set_1,  set_2  := array_of_sets[x_location].set,             array_of_sets[y_location].set;
                        know_1, know_2 := array_of_sets[x_location].know_the_secret, array_of_sets[y_location].know_the_secret;

                        if len(set_1) < len(set_2) {
                            // make set_1 the bigger set.
                            set_1, set_2 = set_2, set_1;
                        }

                        // merge
                        for k := range set_2 {
                            set_1[k] = struct{}{};
                        }

                        array_of_sets[x_location] = Set_And_Know_Secret{
                            set: set_1,
                            know_the_secret: know_1 || know_2,
                        }
                        // remove y, not needed.
                        Remove_Unordered(&array_of_sets, y_location);
                    }
                }

                default: { panic("unhandled case") }
            }
        }

        for _, set_and_known := range array_of_sets {
            // if the secret is know, set all people to true.
            if set_and_known.know_the_secret {
                for k := range set_and_known.set {
                    know_secret[k] = true;
                }
            }
        }

        start_of_group = end_of_group + 1;
    }

    result := make([]int, 0);
    for i, know := range know_secret {
        if know { result = append(result, i); }
    }
    return result;
}


func Contains[T comparable, U any](the_map map[T]U, item T) bool {
    _, found := the_map[item];
    return found;
}

// good old swap and remove
func Remove_Unordered[T any](slice *[]T, index int) {
	if index != len(*slice)-1 {
		(*slice)[index] = (*slice)[len(*slice)-1];
	}
	*slice = (*slice)[:len(*slice)-1];
}
