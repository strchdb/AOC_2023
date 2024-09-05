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
	"red":   0,
	"green": 0,
	"blue":  0,
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
		if err != nil {
			log.Fatal("Could not parse gameid")
		}
		gameSet := copyMap(colors)
		for _, g := range game {
			items := strings.Split(g, ",")
			for _, i := range items {
				t := strings.Split(strings.Trim(i, " "), " ")
				amount, err := strconv.Atoi(t[0])
				if err != nil {
					log.Fatal("could not parse ", t[0])
				}
				color := t[1]
				if gameSet[color] < amount {
					gameSet[color] = amount
				}
			}

		}
		multiplied := 1
		for _, v := range gameSet {
			multiplied *= v
		}
		sum += multiplied
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
