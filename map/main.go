package main

import "fmt"

func main() {
	// words := []string{"hello", "world", "from", "the",
	// 	"best", "language", "in", "the", "world"}
	// //mapWord := make(map[string]int)

	// mapWord := wordCount(words)
	// printMapInt(mapWord)

	voted := map[string]bool{}
	voter := []string{"Den", "Ann", "Bill", "Den", "Peggy", "Thaf"}

	for _, name := range voter {
		check(name, voted)
	}

	printMapBool(voted)
}

func wordCount(words []string) map[string]int {
	m := make(map[string]int)

	for _, word := range words {
		m[word]++
	}

	return m
}

func printMapInt(wordMap map[string]int) {
	for k, v := range wordMap {
		fmt.Println(k, v)
	}
}

func printMapBool(wordMap map[string]bool) {
	for k, v := range wordMap {
		fmt.Println(k, v)
	}
}

func check(name string, voted map[string]bool) {
	_, ok := voted[name]
	if ok {
		fmt.Printf("%s уже голосовал\n", name)

		return
	}

	voted[name] = true
	fmt.Printf("%s проголосовал\n", name)
}
