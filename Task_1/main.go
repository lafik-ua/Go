package main

import "fmt"

func main() {
	arr_int := []int{2, 1, 2, 2, 3, 3, 4}
	inResult := make(map[int]bool)
	var result []int

	// iterate over the slice using range in for loop without index
	for _, val := range arr_int {
		if _, ok := inResult[val]; !ok {
			inResult[val] = true
			result = append(result, val)
		}
	}
	fmt.Println(result)
}

/*
1 - Iterates over the slice.
2 - Checks if a given value of the slice is in the set of the result values. If not, it adds the value to the resulting slice. If so, it skips the value (we already have it in the output slice).
3 - Returns new output slice with duplicates removed.
*/

// func unique[T comparable](s []T) []T {
// 	inResult := make(map[T]bool)
// 	var result []T
// 	for _, str := range s {
// 		if _, ok := inResult[str]; !ok {
// 			inResult[str] = true
// 			result = append(result, str)
// 		}
// 	}
// 	return result
// }

// func main() {
// 	arr_int := []int{1, 1, 2, 2, 3, 3, 4}
// 	arr_str := []string{"abc", "abc", "cde", "efg", "efg", "abc", "cde"}

// 	fmt.Println(unique(arr_str))
// 	fmt.Println(unique(arr_int))

// }
