package helpers

func Int32ArrToIntArr(in []int32) []int {
	out := []int{}
	for _, i := range in {
		out = append(out, int(i))
	}
	return out
}

func IntArrToInt32Arr(in []int) []int32 {
	out := []int32{}
	for _, i := range in {
		out = append(out, int32(i))
	}
	return out
}
