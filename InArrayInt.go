package c_code

import "math"

// in_array
func InArrayInt(array_val int, array []int) bool {
	for _, v := range array {
		if v == array_val {
			return true
		}
	}
	return false
}

func RemoveArrayInt(array_val int, array []int) []int {
	a := []int{}
	for _, v := range array {
		if v != array_val {
			a = append(a, v)
		}
	}
	return a
}

//in_array
func InArrayString(array_val string, array []string) bool {
	for _, v := range array {
		if v == array_val {
			return true
		}
	}
	return false
}
func RemoveArrayString(array_val string, array []string) []string {
	a := []string{}
	for _, v := range array {
		if v != array_val {
			a = append(a, v)
		}
	}
	return a
}

func ArrayChunkString(s []string, size int) [][]string {
	if size < 1 {
		return [][]string{}
	}
	length := len(s)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]string
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, s[i*size:end])
		i++
	}
	return n
}

//array_count_values
func ArrayCountValues(array_list []string) (new_list map[string]int) {
	new_list = make(map[string]int)
	for _, v := range array_list {
		if _, ok := new_list[v]; ok {
			new_list[v]++
		} else {
			new_list[v] = 1
		}
	}
	return
}

//func ArrayCountValues(array_list []string) (new_list map[string]int) {
//	new_list = make(map[string]int)
//	for _, v := range array_list {
//		if _, ok := new_list[v]; ok {
//			new_list[v]++
//		} else {
//			new_list[v] = 1
//		}
//	}
//	return
//}
