package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sum int = 0
var lineNum int = 0

var line string = ""

var symbolIndexByLine map[int][]int = map[int][]int{}
var numbersMapByLine map[int]map[int]string = map[int]map[int]string{}

var validNumbersByLine map[int]map[int]int = map[int]map[int]int{}

var numbers []string = []string{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
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

	for scanner.Scan() {
		line = scanner.Text()

		parseLines()
		lineNum++
	}

	keys := make([]int, 0)
	for k, _ := range symbolIndexByLine {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, linenum := range keys {
		fmt.Printf("line %d\n", linenum)
		if linenum >= 1 {
			for _, symbol := range symbolIndexByLine[linenum] {
				fmt.Printf("symbol %d\n", symbol)
				if validNumbersByLine[linenum+1] == nil {
					validNumbersByLine[linenum+1] = map[int]int{}
				}
				if validNumbersByLine[linenum-1] == nil {
					validNumbersByLine[linenum-1] = map[int]int{}
				}
				validNumbersByLine[linenum] = map[int]int{}
				// fmt.Printf("line %d -1, length %d\n", linenum, len(numbersMapByLine[linenum-1]))
				for key, num := range numbersMapByLine[linenum-1] {

					if symbol-1 >= key && symbol-1 <= key+len(num) {

						n, _ := strconv.Atoi(num)
						validNumbersByLine[linenum-1][key] = n
						// sum += n
					} else if symbol >= key && symbol <= key+len(num) {

						n, _ := strconv.Atoi(num)
						validNumbersByLine[linenum-1][key] = n
						// sum += n
					} else if symbol+1 >= key && symbol+1 <= key+len(num) {

						n, _ := strconv.Atoi(num)
						// sum += n
						validNumbersByLine[linenum-1][key] = n
					} else {
						// fmt.Printf("no match on symbol %d line %d index %d\n", symbol, linenum-2, key)
					}
				}
				// fmt.Printf("line %d +1, length %d\n", linenum, len(numbersMapByLine[linenum+1]))
				for key, num := range numbersMapByLine[linenum+1] {

					if symbol-1 >= key && symbol-1 <= key+len(num) {

						n, _ := strconv.Atoi(num)
						validNumbersByLine[linenum+1][key] = n
						// sum += n
					} else if symbol >= key && symbol <= key+len(num) {

						n, _ := strconv.Atoi(num)
						validNumbersByLine[linenum+1][key] = n
						// sum += n
					} else if symbol+1 >= key && symbol+1 <= key+len(num) {

						n, _ := strconv.Atoi(num)
						validNumbersByLine[linenum+1][key] = n
						// sum += n
					} else {
						// fmt.Printf("no match on symbol %d  line %d index %d\n", symbol, linenum, key)
					}
				}
				// fmt.Printf("line %d   , length %d\n", linenum, len(numbersMapByLine[linenum]))
				for key, num := range numbersMapByLine[linenum] {
					fmt.Printf("symbol %d key %d keyend %d\n", symbol, key, key+len(num))
					if symbol == key+len(num) || symbol == key-1 {
						fmt.Println("OK")
						n, _ := strconv.Atoi(num)
						validNumbersByLine[linenum][key] = n
						// sum += n
					} else {
						// fmt.Printf("no match on symbol %d  line %d index %d\n", symbol, linenum-1, key)
					}
				}
			}
		}
	}

	validNums := make([]int, 0)
	for k, _ := range validNumbersByLine {
		validNums = append(validNums, k)
	}
	sort.Ints(validNums)

	for _, linenum := range validNums {
		for i, num := range validNumbersByLine[linenum] {
			fmt.Printf("Line %d\t\t\tindex %d\t\t\tnum %d\n", linenum, i, num)

			sum += num
		}
	}

	fmt.Printf("SUM: %d\n", sum)
}

func parseLines() {
	symbolIndexByLine[lineNum] = parseSymbols(line)
	numbersMapByLine[lineNum] = parseNumbers(line)
}

func parseSymbols(line string) []int {
	var symbols []int = []int{}

	for i, c := range line {
		if isSymbol(c) {
			// fmt.Printf("Symbol %c\n", c)
			symbols = append(symbols, i)
		}
	}

	return symbols
}

func parseNumbers(line string) map[int]string {
	var symbols map[int]string = map[int]string{}
	var firstIndex int = 0
	var number string = ""

	for i, c := range line {
		if isNumber(c) {
			if number == "" {
				firstIndex = i
			}

			number += toString(c)
		} else if number != "" {
			// fmt.Printf("appending %s\n", number)
			symbols[firstIndex] = number
			number = ""
		}
	}

	return symbols
}

func isSymbol(c rune) bool {
	return isNumber(c) == false && toString(c) != "."
}

func isNumber(c rune) bool {
	for _, n := range numbers {
		if n == toString(c) {
			return true
		}
	}

	return false
}

func runeAtIndex(s string, idx int) rune {
	for i, c := range s {
		if i == idx {
			return c
		}
	}

	return 0
}

func toString(c rune) string {
	return fmt.Sprintf("%c", c)
}

func addNumberOnLine(first rune, second rune, current rune) bool {
	var val bool

	return val
}
