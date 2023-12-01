package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numbers map[string]int = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eigth": 8,
	"nine":  9,
	"orez":  0,
	"eno":   1,
	"owt":   2,
	"eerht": 3,
	"ruof":  4,
	"evif":  5,
	"xis":   6,
	"neves": 7,
	"thgie": 8,
	"enin":  9,
}

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
	var number string
	for _, c := range inp {
		val, err := parseNum(c)
		if err == nil {
			return val
		}
		number += string(c)

		for k, v := range numbers {
			if strings.Contains(number, k) || strings.Contains(reverse(number), k) {
				println(number)
				return v
			}
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
