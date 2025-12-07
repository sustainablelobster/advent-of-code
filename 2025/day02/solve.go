package main

import (
	"fmt"
	"os"
	"strconv"
)

func isValidId1(id int) bool {
	idStr := strconv.Itoa(id)
	idLen := len(idStr)

	if idLen%2 != 0 {
		return true
	}

	midpoint := idLen / 2
	return idStr[:midpoint] != idStr[midpoint:]
}

func isValidId2(id int) bool {
	idStr := strconv.Itoa(id)
	idLen := len(idStr)
	blockMax := idLen / 2

	for i := 1; i <= blockMax; i++ {
		if idLen%i != 0 {
			continue
		}

		block := idStr[:i]
		allBlocksMatch := true

		for j := i; j+i <= idLen; j += i {
			nextBlock := idStr[j : j+i]
			if block != nextBlock {
				allBlocksMatch = false
				break
			}
		}

		if allBlocksMatch {
			return false
		}
	}

	return true
}

func getInvalidIdSum(start int, end int, validator func(int) bool) int {
	sum := 0

	for i := start; i <= end; i++ {
		if !validator(i) {
			sum += i
		}
	}

	return sum
}

func main() {
	infile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	answer1 := 0
	answer2 := 0

	for {
		var n1 int
		var n2 int

		_, err = fmt.Fscanf(infile, "%d-%d", &n1, &n2)
		if err != nil {
			break
		}

		answer1 += getInvalidIdSum(n1, n2, isValidId1)
		answer2 += getInvalidIdSum(n1, n2, isValidId2)
	}

	fmt.Println("Answer 1:", answer1)
	fmt.Println("Answer 2:", answer2)
}
