package main

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

func countMentions(numberOfUsers int, events_but_in_a_bad_format [][]string) []int {
    type Event struct {
        is_message      bool;
        timestamp       int;
        mentions_string string;
    }

    events := make([]Event, len(events_but_in_a_bad_format));
    for i, bad_event := range events_but_in_a_bad_format {

        time, err := strconv.Atoi(bad_event[1]);
        if err != nil { panic("timestamp was not a number"); }

        good_event := Event{
            is_message      : bad_event[0] == "MESSAGE",
            timestamp       : time,
            mentions_string : bad_event[2],
        }
        events[i] = good_event;
    }

    // the timestamps are not in order, lets make them in order.
    slices.SortFunc(events,
        func(a, b Event) int {
            if a.timestamp != b.timestamp {
                return cmp.Compare(a.timestamp, b.timestamp);
            } else {
                // timestamps are the same, offline messages go first
                if !a.is_message {
                    return -1;
                } else {
                    return +1;
                }
            }
        },
    )

    // fmt.Println(events)

    user_mentions := make([]int, numberOfUsers);

    const OFFLINE_TIME_DURATION = 60;

    // keep track of when everybody last went offline
    user_last_offline_start := make([]int, numberOfUsers);
    for i := range user_last_offline_start {
        user_last_offline_start[i] = -OFFLINE_TIME_DURATION;
    }

    for _, event := range events {
        if event.is_message {
            switch event.mentions_string {
                case "ALL": {
                    for i := range user_mentions {
                        user_mentions[i] += 1;
                    }
                }

                case "HERE": {
                    for i := range user_mentions {
                        if user_last_offline_start[i] + OFFLINE_TIME_DURATION <= event.timestamp {
                            user_mentions[i] += 1;
                        }
                    }
                }

                default: {
                    // should be something like ["id0", "id1"]
                    //
                    // lol the go compiler that they use for this dose not have 'SplitSeq()'
                    for _, id := range strings.Split(event.mentions_string, " ") {
                        id_num, err := strconv.Atoi(id[2:]);
                        if err != nil { panic("bad id parse"); }

                        // real world would check this bounds.
                        user_mentions[id_num] += 1;
                    }
                }
            }

        } else {
            id, err := strconv.Atoi(event.mentions_string);
            if err != nil { panic("mentions_string was not an id"); }

            // real world would check this bounds.
            user_last_offline_start[id] = event.timestamp;
        } 
    }

    return user_mentions;
}