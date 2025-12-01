package main

import (
	"fmt"
	"os"
)

const dialSize = 100

func posMod(n int, m int) int {
	return (n % m + m) % m
}

func intAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	infile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	dial := 50
	answer1 := 0
	answer2 := 0

	for {
		var direction byte
		var distance int
		_, err := fmt.Fscanf(infile, "%c%d\n", &direction, &distance)

		if err != nil {
			break
		}

		if direction == 'L' {
			distance = -distance
		}

		sum := dial + distance
		lastDial := dial
		dial = posMod(sum, dialSize)

		if dial == 0 {
			answer1++
		}

		answer2 += intAbs(sum / dialSize)
		if sum <= 0 && lastDial != 0 {
			answer2++
		}
	}

	fmt.Println("Answer 1:", answer1)
	fmt.Println("Answer 2:", answer2)
}
