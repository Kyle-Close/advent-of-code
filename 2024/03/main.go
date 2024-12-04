package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	regex, regErr := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	if regErr != nil {
		log.Fatal(regErr.Error())
	}

	matches := regex.FindAll(bytes, -1)
	enable := true
	var ans int

	for _, v := range matches {
		numRegex, numRegErr := regexp.Compile(`\d+`)
		if numRegErr != nil {
			log.Fatal(numRegErr.Error())
		}

		if string(v) == "do()" {
			enable = true
			continue
		} else if string(v) == "don't()" {
			enable = false
			continue
		}

		if !enable {
			continue
		}

		nums := numRegex.FindAll(v, -1)
		if len(nums) != 2 {
			log.Fatal(string(v))
		}

		num1, num1Err := strconv.Atoi(string(nums[0]))
		if num1Err != nil {
			log.Fatal(num1Err.Error())
		}
		num2, num2Err := strconv.Atoi(string(nums[1]))
		if num2Err != nil {
			log.Fatal(num2Err.Error())
		}

		ans += num1 * num2
	}

	log.Println(ans)
}
