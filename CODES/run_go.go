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

/*func searchDict(key string, testDict map[string]int) int64 {
	_ = testDict[key]
}*/

/*func displayKeyNames(testDict map[string]int) {
	for key := range testDict {
		fmt.Println(key)
	}
}*/

func main() {
	dictSize := 1000
	searchTimes := []int{50, 500, 5000, 50000 }

	testDict := createDict(dictSize)
	//displayKeyNames(testDict)

	for i :=0; i< 5; i++ {
		fmt.Printf("--------\n")
	
		for _, searchTime := range searchTimes {
			//totalTime := int64(0)
			startTime := time.Now()

			for i := 0; i < searchTime; i++ {
				randomKey := getRandomKey(testDict)
				//timeTaken := searchDict(randomKey, testDict)
				_ = testDict[randomKey]
			}
			endTime := time.Now()
			totalTime := endTime.Sub(startTime)
			fmt.Printf("Search time for %d searches: %.2f ms\n", searchTime, totalTime.Seconds()*1000)
		}
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
