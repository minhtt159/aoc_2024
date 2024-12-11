package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func read_input(file_name string) []string {
	result := []string{}
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

type BlinkResult struct {
	first  int
	second int
}

var lastResult = []int{}

func blinkStone(stoneInt int) BlinkResult {
	if stoneInt == 0 {
		return BlinkResult{1, -1}
	}
	stoneLength := int(math.Log10(float64(stoneInt))) + 1
	if stoneLength%2 == 0 {
		firstHalf := stoneInt / int(math.Pow10(stoneLength/2))
		secondHalf := stoneInt % int(math.Pow10(stoneLength/2))
		// fmt.Println("Split", stoneInt, firstHalf, secondHalf)
		return BlinkResult{firstHalf, secondHalf}
	}
	return BlinkResult{stoneInt * 2024, -1}
}

func blinkStoneWithCache() func(stoneInt int) BlinkResult {
	cache := make(map[int]BlinkResult)
	return func(stoneInt int) BlinkResult {
		if ret, ok := cache[stoneInt]; ok {
			// fmt.Println("Blink cache", stoneInt, ret)
			return ret
		}
		blinkResult := blinkStone(stoneInt)
		cache[stoneInt] = blinkResult
		return blinkResult
	}
}

// Attempt 1
func countStone(blinker func(p int) BlinkResult, stoneInt int, blinks int) int {
	if stoneInt == -1 {
		return 0
	}
	blinkResult := blinker(stoneInt)
	if blinks == 1 {
		if blinkResult.second == -1 {
			return 1
		}
		return 2
	}
	return countStone(blinker, blinkResult.first, blinks-1) + countStone(blinker, blinkResult.second, blinks-1)
}

func countStoneWithCache() func(blinker func(p int) BlinkResult, stoneInt int, blinks int) int {
	cache := make(map[int]map[int]int)
	return func(blinker func(p int) BlinkResult, stoneInt int, blinks int) int {
		if _, ok := cache[stoneInt]; !ok {
			cache[stoneInt] = make(map[int]int)
		}
		if ret, ok := cache[stoneInt][blinks]; ok {
			// fmt.Println("Count cache hit", stoneInt, blinks, ret)
			return ret
		}
		cache[stoneInt][blinks] = countStone(blinker, stoneInt, blinks)
		return cache[stoneInt][blinks]
	}
}

// Attempt 2
var globalCache = make(map[int]map[int]int)

func countStoneWithCache2(cache map[int]map[int]int, blinker func(stoneInt int) BlinkResult, stoneInt int, blinks int) int {
	if stoneInt == -1 {
		return 0
	}
	if _, ok := cache[stoneInt]; !ok {
		cache[stoneInt] = make(map[int]int)
	}
	if ret, ok := cache[stoneInt][blinks]; ok {
		// fmt.Println("Count cache hit", stoneInt, blinks, ret)
		return ret
	}
	blinkResult := blinker(stoneInt)
	if blinks == 1 {
		// lastResult = append(lastResult, blinkResult.first)
		if blinkResult.second == -1 {
			cache[stoneInt][blinks] = 1
		} else {
			// lastResult = append(lastResult, blinkResult.second)
			cache[stoneInt][blinks] = 2
		}
	} else {
		if _, ok := cache[blinkResult.first]; !ok {
			cache[blinkResult.first] = make(map[int]int)
		}
		leftResult := countStoneWithCache2(cache, blinker, blinkResult.first, blinks-1)
		cache[blinkResult.first][blinks-1] = leftResult
		if _, ok := cache[blinkResult.second]; !ok {
			cache[blinkResult.second] = make(map[int]int)
		}
		rightResult := countStoneWithCache2(cache, blinker, blinkResult.second, blinks-1)
		cache[blinkResult.second][blinks-1] = rightResult

		cache[stoneInt][blinks] = leftResult + rightResult
	}
	return cache[stoneInt][blinks]
}

func main() {
	// file_name := "input.txt"
	file_name := "real_input.txt"
	input := read_input(file_name)
	array := input[0]

	var allStones []int
	for _, stone_id := range strings.Split(array, " ") {
		stoneInt, _ := strconv.Atoi(stone_id)
		allStones = append(allStones, stoneInt)
	}
	fmt.Println(len(allStones), allStones)

	// maxRound := 5
	// maxRound := 25
	maxRound := 75
	result := 0
	blinker := blinkStoneWithCache()

	// Attempt 1
	// counter := countStoneWithCache()

	for _, thisStone := range allStones {
		// result += counter(blinker, thisStone, maxRound)
		// Attempt 2
		result += countStoneWithCache2(globalCache, blinker, thisStone, maxRound)
		//
		// break
	}

	fmt.Println(result)
	fmt.Println(lastResult)
}
