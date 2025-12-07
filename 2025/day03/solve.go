package main

import (
	"fmt"
	"os"
)

func charToInt(r byte) int {
	return int(r - '0')
}

func maxJoltage(bank string, numBatts int) int {
	batts := make([]int, numBatts)
	bankLen := len(bank)

	for i, iMax := 0, bankLen-numBatts; i <= iMax; i++ {
		for j := 0; j < numBatts; j++ {
			batt := charToInt(bank[i+j])

			if batt > batts[j] {
				batts[j] = batt
				for k := j + 1; k < numBatts; k++ {
					batts[k] = charToInt(bank[i+k])
				}
				break
			}
		}
	}

	joltage := 0

	for _, batt := range batts {
		joltage = (joltage * 10) + batt
	}

	return joltage
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
		var bank string

		_, err = fmt.Fscanln(infile, &bank)
		if err != nil {
			break
		}

		answer1 += maxJoltage(bank, 2)
		answer2 += maxJoltage(bank, 12)
	}

	fmt.Println("Answer 1:", answer1)
	fmt.Println("Answer 2:", answer2)
}
