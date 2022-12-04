package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func topElfCals() {
	var elfCals, maxCals int

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	elves := strings.Split(string(data), "\n\n")
	for _, elve := range elves {
		calStrs := strings.Split(elve, "\n")
		elfCals = 0

		// Sum up cals for one elf
		for _, calStr := range calStrs {
			i, _ := strconv.Atoi(calStr)
			elfCals += i
		}

		if elfCals > maxCals {
			maxCals = elfCals
		}
	}

	fmt.Println("maxcals:", maxCals)
}

func main() {
	var cals int

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	elfCals := make([]int, 0)

	elves := strings.Split(string(data), "\n\n")
	for _, elve := range elves {
		calStrs := strings.Split(elve, "\n")
		cals = 0

		// Sum up cals for one elf
		for _, calStr := range calStrs {
			i, _ := strconv.Atoi(calStr)
			cals += i
		}

		elfCals = append(elfCals, cals)
	}

	sort.Slice(elfCals, func(i, j int) bool {
		return elfCals[i] > elfCals[j]
	})

	fmt.Println(elfCals[0] + elfCals[1] + elfCals[2])
}
