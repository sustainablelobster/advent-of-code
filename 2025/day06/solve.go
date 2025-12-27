package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type problem struct {
	numbers  []int
	operator byte
}

func sum(numbers []int) int {
	s := 0
	for _, n := range numbers {
		s += n
	}
	return s
}

func product(numbers []int) int {
	p := 1
	for _, n := range numbers {
		p *= n
	}
	return p
}

func (p problem) Solve() int {
	if p.operator == '+' {
		return sum(p.numbers)
	}
	return product(p.numbers)
}

func parsePart1Problems(rawInput []string) []problem {
	indexMax := len(rawInput) - 1
	rawOperators := strings.Fields(rawInput[indexMax])
	numOfProbs := len(rawOperators)
	problems := make([]problem, numOfProbs)

	for _, row := range rawInput[0:indexMax] {
		rawNums := strings.Fields(row)
		for j, rawNum := range rawNums {
			n, _ := strconv.Atoi(rawNum)
			problems[j].numbers = append(problems[j].numbers, n)
		}
	}

	for i, rawOperator := range rawOperators {
		problems[i].operator = rawOperator[0]
	}

	return problems
}

func parsePart2Problems(rawInput []string) []problem {
	rowMax := len(rawInput) - 1
	columnMax := len(rawInput[0]) - 1
	rawOperators := strings.Fields(rawInput[rowMax])
	numOfProbs := len(rawOperators)
	problems := make([]problem, numOfProbs)

	for i, p := columnMax, numOfProbs-1; i >= 0; i-- {
		var rawNumBytes []byte
		isEmptyColumn := true

		for j := 0; j < rowMax; j++ {
			if rawInput[j][i] == ' ' {
				continue
			}

			isEmptyColumn = false
			rawNumBytes = append(rawNumBytes, rawInput[j][i])
		}

		if isEmptyColumn {
			p--
		} else {
			rawNum := string(rawNumBytes)
			num, _ := strconv.Atoi(rawNum)
			problems[p].numbers = append(problems[p].numbers, num)
		}
	}

	for i, rawOperator := range rawOperators {
		problems[i].operator = rawOperator[0]
	}

	return problems
}

func sumSolutions(problems []problem) int {
	total := 0

	for _, p := range problems {
		total += p.Solve()
	}

	return total
}

func main() {
	infile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	scanner := bufio.NewScanner(infile)
	var rawInput []string

	for scanner.Scan() {
		rawInput = append(rawInput, scanner.Text())
	}

	part1Problems := parsePart1Problems(rawInput)
	answer1 := sumSolutions(part1Problems)
	fmt.Println("Answer 1:", answer1)

	part2Problems := parsePart2Problems(rawInput)
	answer2 := sumSolutions(part2Problems)
	fmt.Println("Answer 2:", answer2)
}
