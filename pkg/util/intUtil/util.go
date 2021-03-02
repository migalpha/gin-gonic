package intUtil

func UniqueIntSlice(intSlice []int) []int {
	/*
		int slice with repeted values to unique int slice  values
	*/
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
