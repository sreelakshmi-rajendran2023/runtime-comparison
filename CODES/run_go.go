package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func createDict(dictSize int) map[string]int {
	testDict := make(map[string]int)
	prefix := strings.ToLower("abcdefghijklmnopqrstuvwxyz_")
	for i := 0; i < dictSize; i++ {
		key := fmt.Sprintf("%s%03d", prefix, i)
		testDict[key] = i
	}
	return testDict
}

func searchDict(key string, testDict map[string]int) int64 {
	startTime := time.Now().UnixNano()
	_ = testDict[key] // Ignoring the result for now
	endTime := time.Now().UnixNano()
	searchTime := endTime - startTime
	return searchTime
}

/*func displayKeyNames(testDict map[string]int) {
	for key := range testDict {
		fmt.Println(key)
	}
}*/

func main() {
	dictSize := 1000
	searchTimes := []int{50, 500, 5000}

	testDict := createDict(dictSize)
	//displayKeyNames(testDict)

	for _, searchTime := range searchTimes {
		totalTime := int64(0)
		for i := 0; i < searchTime; i++ {
			randomKey := getRandomKey(testDict)
			timeTaken := searchDict(randomKey, testDict)
			totalTime += timeTaken
		}
		averageTime := float64(totalTime) / float64(searchTime)
		fmt.Printf("\nAverage search time for %d searches: %.2f nanoseconds\n", searchTime, averageTime)
	}
}

func getRandomKey(testDict map[string]int) string {
	keys := make([]string, 0, len(testDict))
	for key := range testDict {
		keys = append(keys, key)
	}
	randomIndex := rand.Intn(len(keys))
	return keys[randomIndex]
}
