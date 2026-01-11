package main

type Keypad_Number int

// their position in the keys array
const (
	KEY_ONE   Keypad_Number = 0
	KEY_TWO                 = 1
	KEY_THREE               = 2
	KEY_FOUR                = 3
	KEY_FIVE                = 4
	KEY_SIX                 = 5
	KEY_SEVEN               = 6
	KEY_EIGHT               = 7
	KEY_NINE                = 8
	KEY_ZERO                = 9
)
const NUM_KEYS = 10

type Keypad struct {
	keys [NUM_KEYS]uint32 // this cannot be an int32, but it just makes it into a uint32 :)
}

func knightDialer(n int) int {
	if n <= 0 {
		return 0
	}

	const MOD = 1000000007

	keypads_base := [2]Keypad{}
	prev_keypad := &keypads_base[0]
	next_keypad := &keypads_base[1]

	// knight can start in any position
	for i := range prev_keypad.keys {
		prev_keypad.keys[i] = 1
	}

	// this will stay static, forever.
	//
	// but you could imagine a harder problem, where this could change
	// (or not change, and have a weird generation rule.)
	//
	// also forgive my C-ish
	//
	// also make arrays const you cowards
	var KEY_CONNECTIONS = [NUM_KEYS][]Keypad_Number{
		/* [KEY_ONE  ] = */ {KEY_SIX, KEY_EIGHT},
		/* [KEY_TWO  ] = */ {KEY_SEVEN, KEY_NINE},
		/* [KEY_THREE] = */ {KEY_FOUR, KEY_EIGHT},
		/* [KEY_FOUR ] = */ {KEY_THREE, KEY_NINE, KEY_ZERO},
		/* [KEY_FIVE ] = */ {}, // haha
		/* [KEY_SIX  ] = */ {KEY_ONE, KEY_SEVEN, KEY_ZERO},
		/* [KEY_SEVEN] = */ {KEY_TWO, KEY_SIX},
		/* [KEY_EIGHT] = */ {KEY_ONE, KEY_THREE},
		/* [KEY_NINE ] = */ {KEY_TWO, KEY_FOUR},
		/* [KEY_ZERO ] = */ {KEY_FOUR, KEY_SIX},
	}

	for range n - 1 {
		for key, connections := range KEY_CONNECTIONS {
			next_keypad.keys[key] = 0
			for _, connection := range connections {
				next_keypad.keys[key] += prev_keypad.keys[connection]
				// would be smarter to mod every step of the way,
				// but the max len(connections) == 3, and this *just* cannot overflow in uint32
				// next_keypad.keys[key] = next_keypad.keys[key] % MOD
			}
		}

		// mod the keys, all at once for speed. probably, i mean, this is
		// a constant length array, surly the compiler is doing something smart here,
		for j := range next_keypad.keys {
			next_keypad.keys[j] %= MOD
		}

		{ // swap buffers
			tmp := prev_keypad
			prev_keypad = next_keypad
			next_keypad = tmp
		}
	}

	// count total
	result := uint32(0)
	for _, key := range prev_keypad.keys {
		result = (result + key) % MOD
	}
	return int(result)
}
