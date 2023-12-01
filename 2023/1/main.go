package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := input.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(input)

	var text string
	var number int
	var sum int

	for scanner.Scan() {
		text = scanner.Text()

		fmt.Println(text)

		number = parse(text)

		fmt.Println(number)

		sum += number
	}

	fmt.Printf("SUM: %d\n", sum)
}

func parse(inp string) int {
	var num int

	numstr := fmt.Sprintf("%d%d", parseFromStart(inp), parseFromEnd(inp))

	num, _ = strconv.Atoi(numstr)

	return num
}

func parseFromStart(inp string) int {
	for _, c := range inp {
		val, err := parseNum(c)
		if err == nil {
			return val
		}
	}

	return 0
}

func parseFromEnd(inp string) int {
	return parseFromStart(reverse(inp))
}

func parseNum(i rune) (int, error) {
	return strconv.Atoi(fmt.Sprintf("%c", i))
}

func reverse(s string) string {
	var str string
	for _, v := range s {
		str = string(v) + str
	}
	return str
}
