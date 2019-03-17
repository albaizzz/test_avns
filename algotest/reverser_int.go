package test_avns

func Reverse_int(params int) int {
	newVal := 0
	for params > 0 {
		lastVal := params % 10
		newVal *= 10
		newVal += lastVal
		params /= 10
	}
	return newVal
}
