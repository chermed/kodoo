package kotils

func IntInSlice(i int, arr []int) bool {
	for _, value := range arr {
		if value == i {
			return true
		}
	}
	return false
}
