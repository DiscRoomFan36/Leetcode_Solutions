package main

type Foo struct {
	// smallest it can be, don't know if uint helps
	//
	// if smallest == 0, no element is there yet,
	smallest, largest uint32
}

func countCoveredBuildings(n int, buildings [][]int) int {
	column_array := make([]Foo, n)
	row_array := make([]Foo, n)

	for i := range buildings {
		x, y := uint32(buildings[i][0]), uint32(buildings[i][1])

		// zero index the array, but not the Foo's
		col := &column_array[y-1]
		if col.smallest == 0 {
			col.smallest = x
			col.largest = x
		} else {
			col.smallest = min(col.smallest, x)
			col.largest = max(col.largest, x)
		}

		row := &row_array[x-1]
		if row.smallest == 0 {
			row.smallest = y
			row.largest = y
		} else {
			row.smallest = min(row.smallest, y)
			row.largest = max(row.largest, y)
		}
	}

	result := 0
	for i := range buildings {
		x, y := uint32(buildings[i][0]), uint32(buildings[i][1])

		// quickly check if there is something either side of us.
		col := column_array[y-1]
		surrounded_lr := col.smallest != x && col.largest != x

		row := row_array[x-1]
		surrounded_ud := row.smallest != y && row.largest != y

		if surrounded_lr && surrounded_ud {
			result += 1
		}
	}
	return result
}
