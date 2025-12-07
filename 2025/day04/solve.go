package main

import (
	"fmt"
	"os"
)

func markAccessible(input [][]byte) int {
	height := len(input)
	width := len(input[0])
	rollsMarked := 0

	for i := range input {
		for j := range input[i] {
			if input[i][j] != '@' {
				continue
			}

			adjacentRolls := 0
			neighbors := [8][2]int{
				{i - 1, j},
				{i - 1, j + 1},
				{i, j + 1},
				{i + 1, j + 1},
				{i + 1, j},
				{i + 1, j - 1},
				{i, j - 1},
				{i - 1, j - 1},
			}

			for _, neighbor := range neighbors {
				n1 := neighbor[0]
				n2 := neighbor[1]
				n1IsValid := n1 >= 0 && n1 < height
				n2IsValid := n2 >= 0 && n2 < width

				if n1IsValid && n2IsValid && input[n1][n2] != '.' {
					adjacentRolls++
				}
			}

			if adjacentRolls < 4 {
				rollsMarked++
				input[i][j] = 'x'
			}
		}
	}

	return rollsMarked
}

func removeAccessible(input [][]byte) {
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'x' {
				input[i][j] = '.'
			}
		}
	}
}

func main() {
	infile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	var input [][]byte

	for {
		var line []byte

		_, err = fmt.Fscanln(infile, &line)
		if err != nil {
			break
		}

		input = append(input, line)
	}

	accessibleRolls := markAccessible(input)
	answer1 := accessibleRolls
	answer2 := accessibleRolls

	for accessibleRolls > 0 {
		removeAccessible(input)
		accessibleRolls = markAccessible(input)
		answer2 += accessibleRolls
	}

	fmt.Println("Answer 1:", answer1)
	fmt.Println("Answer 2:", answer2)
}
