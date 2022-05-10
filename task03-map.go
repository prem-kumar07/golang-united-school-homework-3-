package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	keys := make([]int, 0, len(input))

	v_slice := make([]string, len(keys))

	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		v_slice = append(v_slice, input[k])
	}
	return v_slice
}

