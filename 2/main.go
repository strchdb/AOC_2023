package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var colors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	arg := "input.txt"
	file, err := os.Open(arg)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ";")
		temp := strings.Split(game[0], ":")
		game[0] = temp[1]
		isPossible := true
		id, err := strconv.Atoi(strings.Split(temp[0], " ")[1])
		if err != nil {
			log.Fatal("Could not parse gameid")
		}
		for _, g := range game {
			if !isPossible {
				break
			}
			items := strings.Split(g, ",")
			gameSet := copyMap(colors)
			for _, i := range items {
				t := strings.Split(strings.Trim(i, " "), " ")
				amount, err := strconv.Atoi(t[0])
				if err != nil {
					log.Fatal("could not parse ", t[0])
				}
				color := t[1]
				gameSet[color] = gameSet[color] - amount
				if gameSet[color] < 0 {
					isPossible = false
					break
				}
			}

		}
		if isPossible {
			sum += id
		}
	}
	fmt.Println("Sum of possible games: ", sum)
}
func copyMap(originalMap map[string]int) map[string]int {
	copy := make(map[string]int)

	for key, value := range originalMap {
		copy[key] = value
	}

	return copy
}
