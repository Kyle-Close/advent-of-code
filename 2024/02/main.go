package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	scanner := bufio.NewScanner(file)
	var sCount int

	for scanner.Scan() {
		// Split line on space to get each level
		fields := strings.Fields(scanner.Text())
		// Check if original string is valid
		safe := isSafe(fields)

		if safe {
			sCount++
			continue
		}

		// If not valid, we need to remove elements until we get valid or mark as not valid
	inner:
		for i := 0; i < len(fields)-1; i++ {
			sub := removeEl(fields, i)
			if isSafe(sub) {
				sCount++
				break inner
			}
		}

	}

	log.Println(sCount)
}

func removeEl(slice []string, index int) []string {
	res := make([]string, len(slice)-1)
	copy(res, slice[:index])
	copy(res[index:], slice[index+1:])
	return res
}

func isSafe(fields []string) bool {
	if len(fields) == 0 {
		return true
	}

	isInc := false

	for i, v := range fields {
		current, curErr := strconv.Atoi(v)
		if curErr != nil {
			log.Fatal(curErr.Error())
		}

		// Check if there is a next element to check. If not, safe to return true
		if len(fields) == i+1 {
			return true
		}

		next, nErr := strconv.Atoi(fields[i+1])
		if nErr != nil {
			log.Fatal(nErr.Error())
		}

		if i == 0 {
			if current-next < 0 {
				isInc = true
			}
		}

		if !isInc {
			if current-next <= 0 || current-next > 3 {
				return false
			}
		} else {
			if current-next >= 0 || current-next < -3 {
				return false
			}
		}
	}
	return true
}
