package main

func modifyString(_s string) string {
	if len(_s) == 0 {
		return _s
	}
	if len(_s) == 1 && _s[0] == '?' {
		return "a"
	}

	s := make([]byte, len(_s))
	for i := range _s {
		s[i] = _s[i]
	}

	if s[0] == '?' {
		// length must be at least 2
		if s[1] == 'a' {
			s[0] = 'b'
		} else {
			s[0] = 'a'
		}
	}

	for i := 1; i < len(s); i++ {
		if s[i] != '?' {
			continue
		}

		is_next_letter := (i != len(s)-1)

		letter := byte('a')
		if s[i-1] == 'a' {
			if is_next_letter && s[i+1] == 'b' {
				letter = 'c'
			} else {
				letter = 'b'
			}
		}

		if is_next_letter && s[i+1] == 'a' {
			if s[i-1] == 'b' {
				letter = 'c'
			} else {
				letter = 'b'
			}
		}

		s[i] = letter
	}

	return string(s)
}
