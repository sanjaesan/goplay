package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "newFunctionCall"
	fmt.Println(camelToUnder(str))
}

func camelToUnder(str string) string {
	var sb strings.Builder
	for _, val := range str {
		if val >= 'A' && val <= 'Z' {
			sb.WriteRune('_')
		}
		sb.WriteRune(val)
	}
	return strings.ToLower(sb.String())
}
