// Author: Muhammad Farhan Mardadi / fmardadi@gmail.com / 1st November 2023

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type SpeSkillTest struct{}

func main() {
	s := SpeSkillTest{}

	// 1. Narcissistic Number
	testCase1 := []int{1634, 153, 11, 370}
	for _, val := range testCase1 {
		fmt.Printf("Narcissistic Number %d: %t \n", val, s.narcissisticNumber(val))
	}

	fmt.Println()

	// 2. Parity Outlier
	testCase2 := [][]int{
		{2, 4, 0, 100, 4, 11, 2602, 36},
		{160, 3, 1719, 19, 11, 13, -21},
		{11, 13, 15, 19, 9, 13, -21},
	}
	for _, val := range testCase2 {
		result := s.parityOutlier(val)
		if result != -1 {
			fmt.Printf("Parity Outlier %d: %d \n", val, result)
		} else {
			fmt.Printf("Parity Outlier %d: %t \n", val, false)
		}
	}

	fmt.Println()

	// 3. Needle in the Haystack
	testCase3 := []string{"blue", "black", "red"}
	haystack := []string{"red", "blue", "yellow", "black", "grey"}
	fmt.Printf("Haystack: %s \n", haystack)
	for _, val := range testCase3 {
		fmt.Printf("Needle in the Haystack %s: %d \n", val, s.needleInTheHaystack(haystack, val))
	}

	fmt.Println()

	// 4. The Blue Ocean Reverse
	testCase4 := []struct {
		arr1 []int
		arr2 []int
	}{
		{[]int{1, 2, 3, 4, 6, 10}, []int{1}},
		{[]int{1, 5, 5, 5, 5, 3}, []int{5}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3}},
	}
	for _, val := range testCase4 {
		fmt.Printf("The Blue Ocean Reverse %d - %d : %d \n", val.arr1, val.arr2, s.blueOceanReverse(val.arr1, val.arr2))
	}
}

func (s *SpeSkillTest) narcissisticNumber(n int) bool {
	var total float64

	numbers := strings.Split(fmt.Sprintf("%d", n), "")
	len := float64(len(numbers))

	for _, number := range numbers {
		val, _ := strconv.Atoi(number)
		total = total + math.Pow(float64(val), len)
	}

	if int(total) == n {
		return true
	}

	return false
}

func (s SpeSkillTest) parityOutlier(n []int) int {
	var oddCounter []int
	var evenCounter []int

	for _, number := range n {
		if number%2 == 0 {
			evenCounter = append(evenCounter, number)
		} else {
			oddCounter = append(oddCounter, number)
		}
	}

	if len(evenCounter) == 1 {
		return evenCounter[0]
	} else if len(oddCounter) == 1 {
		return oddCounter[0]
	} else {
		return -1
	}

}

func (s SpeSkillTest) needleInTheHaystack(haystack []string, needle string) int {
	for i, val := range haystack {
		if val == needle {
			return i
		}
	}

	return -1
}

func (s SpeSkillTest) blueOceanReverse(arr1, arr2 []int) []int {
	var differ []int
	mapArr2 := make(map[int]bool)

	for _, val := range arr2 {
		mapArr2[val] = true
	}

	for _, val := range arr1 {
		if _, ok := mapArr2[val]; !ok {
			differ = append(differ, val)
		}
	}

	return differ
}
