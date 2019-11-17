package main

func reverse(s string) string {
	runify := []byte(s)
	for i, j := 0, len(runify)-1; i < j; i, j = i+1, j-1 {
		runify[i], runify[j] = runify[j], runify[i]
	}
	return string(runify)
}

// func reverse(num []int) {
// 	first := 0
// 	last := len(num) - 1
// 	for first < last {
// 		num[first], num[last] = num[last], num[first]
// 		first++
// 		last--
// 	}
// }
func main() {}
