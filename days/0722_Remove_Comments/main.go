package main

import "strings"

func removeComments(source []string) []string {
	// in place, lets go, 2nd try, EZ
	for i := 0; i < len(source); {
		s := source[i]
		line_index := strings.Index(s, "//")
		multi_index := strings.Index(s, "/*")

		do_line := (line_index != -1) && (multi_index == -1 || line_index < multi_index)
		do_multi := (multi_index != -1) && (line_index == -1 || multi_index < line_index)

		if do_line {
			_, before, _ := split_once(s, "//")
			if len(before) == 0 {
				source = append(source[:i], source[i+1:]...)
				continue
			}
			source[i] = before
			i += 1
			continue
		}

		if do_multi {
			_, before, after := split_once(s, "/*")

			ok, _, after_2 := split_once(after, "*/")
			if ok {
				if len(before) == 0 && len(after_2) == 0 {
					source = append(source[:i], source[i+1:]...)
					continue
				}
				source[i] = before + after_2
				continue // we still gotta check for '//' comments
			} else {
				// find '*/' on another line
				end := -1
				found := false
				after_3 := ""
				for j := i + 1; j < len(source); j++ {
					found, _, after_3 = split_once(source[j], "*/")
					if found {
						end = j
						break
					}
				}
				if !found {
					panic("Block comments must be closed.")
				}

				if len(before) == 0 && len(after_3) == 0 {
					source = append(source[:i], source[end+1:]...)
					continue
				} else {
					source[i] = before + after_3
					source = append(source[:i+1], source[end+1:]...)
					continue // gotta check for '//'
				}
			}
		}

		i += 1
	}

	return source
}

// returns the string split into two pieces,
//
// if the separator was not in the string,
// before contains the whole string, and after is empty.
// (and found == false)
func split_once(str, sep string) (found bool, before, after string) {
	index := strings.Index(str, sep)
	if index == -1 {
		return false, str, ""
	}

	_before := str[:index]
	_after := str[index+len(sep):]

	return true, _before, _after
}
