package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

// one, two, ,, , six, four, five, nine seven,eight,three

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		digits := findDigits(line)

		if len(digits) == 1 {
			first := digits[0]
			number, err := strconv.Atoi(fmt.Sprintf("%d%d", first, first))
			if err != nil {
				log.Fatal(err)
			}
			sum += number
		} else if len(digits) >= 2 {
			first := digits[0]
			last := digits[len(digits)-1]
			number, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
			if err != nil {
				log.Fatal(err)
			}
			sum += number
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func findDigits(s string) []int {
	digits := []int{}
	words := make(map[string]int)
	words["one"] = 1
	words["two"] = 2
	words["three"] = 3
	words["four"] = 4
	words["five"] = 5
	words["six"] = 6
	words["seven"] = 7
	words["eight"] = 8
	words["nine"] = 9
	for i, c := range s {
		if unicode.IsDigit(c) {
			digits = append(digits, int(c-'0'))
		} else {
			n := len(s) - i
			if n < 3 {
				continue
			}
			if n > 4 {
				for k, v := range words {
					if s[i:len(k)+i] == k {
						digits = append(digits, v)
						break
					}
				}
			} else if n < 5 && n > 3 {
				keys := make([]string, 0, len(words))
				for k := range words {
					if len(k) < 5 {
						keys = append(keys, k)
					}
				}
				for _, v := range keys {
					if s[i:len(v)+i] == v {
						digits = append(digits, words[v])
						break
					}
				}

			} else if n < 4 {
				keys := make([]string, 0, len(words))
				for k := range words {
					if len(k) < 4 {
						keys = append(keys, k)
					}
				}
				for _, v := range keys {
					if s[i:len(v)+i] == v {
						digits = append(digits, words[v])
						break
					}
				}

			}

		}
	}
	return digits
}
