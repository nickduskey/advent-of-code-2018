package day1

import (
	"github.com/nickduskey/aoc/utils"
	"log"
	"strings"
	"fmt"
	"strconv"
)

func Run() {
	input, err := utils.ReadInputFile("day1/input")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	inputString := string(input[:])
	pieces := strings.Split(inputString, "\n")
	frequency := 0
	pastFrequencies := []int{frequency}
	i := 0
	for {
		if i >= len(pieces) - 1 {
			// reset i
			i = 0
		}
		if pieces[i] == "" {
			continue
		}
		change, err := strconv.Atoi(pieces[i])
		if err != nil {
			log.Fatal(err)
		}
		frequency += change
		if isDup := isDuplicateFrequencies(pastFrequencies, frequency); isDup {
			fmt.Println("is duplicate")
			fmt.Println(frequency)
			break
		}
		pastFrequencies = append(pastFrequencies, frequency)
		i++
	}
	fmt.Println("done")
}

func isDuplicateFrequencies(past []int, current int) bool {
	for _, f := range past {
		if f == current {
			return true
		}
	}
	return false
}
