package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "dani Maulana"
	fmt.Println("input =", input)

	outputArr := []string{}
	mapInput := make(map[string]int)

	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")

	// for _, v := range input {
	// 	if _, ok := mapInput[string(v)]; !ok {
	// 		countCharr := strings.Count(input, string(v))
	// 		mapInput[string(v)] = countCharr
	// 		if strings.Count(input, string(v)) == 1 {
	// 			outputArr = append(outputArr, string(v))
	// 		} else {
	// 			countCharrStr := strconv.Itoa(countCharr)
	// 			outputArr = append(outputArr, countCharrStr+string(v))
	// 		}
	// 	}
	// }

	for _, v := range input {
		if _, ok := mapInput[string(v)]; !ok {
			countCharr := countString(input, string(v))
			mapInput[string(v)] = countCharr
			if countString(input, string(v)) == 1 {
				outputArr = append(outputArr, string(v))
			} else {
				countCharrStr := strconv.Itoa(countCharr)
				outputArr = append(outputArr, countCharrStr+string(v))
			}
		}
	}

	output := strings.Join(outputArr, "")

	fmt.Println("output =", output)
}

// dani
func countString(str, subStr string) int {
	count := 0
	for _, v := range str {
		if string(v) == subStr {
			count++
		}
	}
	return count
}
