package test_avns

func countTrailingBits(N uint) uint {
	var _lookup = [...]uint{
		32, 0, 1, 26, 2, 23, 27, 0, 3, 16, 24, 30, 28, 11, 0, 13, 4,
		7, 17, 0, 25, 22, 31, 15, 29, 10, 12, 6, 0, 21, 14, 9, 5,
		20, 8, 19, 18}

	return _lookup[(N&-N)%37]
}

func Binary_gap(n int) int {

	if n < 5 {
		return 0
	}

	// enum way in golang:
	type state int
	const (
		first_0 state = iota
		first_1
		gap_started_0
	)

	var num int
	num = n
	max_gap, cur_gap := 0, 0
	var (
		cur_state  state
		next_state state
	)

	if num&0x01 == 1 {
		cur_state = first_1
	} else {
		cur_state = first_0
	}

	num = num >> 1
	for ; num != 0; num >>= 1 {
		switch cur_state {
		case first_0:
			if num&0x01 == 1 {
				next_state = first_1
			} else {
				next_state = first_0
			}
			break

		case first_1:
			if num&0x01 == 1 {
				next_state = first_1
			} else {
				cur_gap = 1
				next_state = gap_started_0
			}
			break

		case gap_started_0:
			if num&0x01 == 1 {
				next_state = first_1
				if max_gap < cur_gap {
					max_gap = cur_gap
				}
			} else {
				cur_gap++
				next_state = gap_started_0
			}
			break
		}
		cur_state = next_state
	}

	return max_gap
}
