package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left, right []int

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			log.Fatal("Error parsing input line")
		}

		leftNum, leftErr := strconv.Atoi(fields[0])
		if leftErr != nil {
			log.Fatal(leftErr.Error())
		}
		rightNum, rightErr := strconv.Atoi(fields[1])
		if rightErr != nil {
			log.Fatal(rightErr.Error())
		}

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	// Sort the lists
	slices.Sort(left)
	slices.Sort(right)

	var distance int
	oMap := make(map[int]int)

	for i := 0; i < len(left); i++ {
		distance += int(math.Abs(float64(left[i]) - float64(right[i])))
		if oMap[right[i]] != 0 {
			oMap[right[i]] += 1
		} else {
			oMap[right[i]] = 1
		}
	}

	var similarity int

	for i := 0; i < len(left); i++ {
		occ := oMap[left[i]]
		inc := occ * left[i]
		similarity += inc
	}

	log.Println(similarity)
}
