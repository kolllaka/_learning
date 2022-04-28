package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["alice"] = []string{"peggy", "you"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{"tedduu"}
	graph["tedduu"] = []string{}

	queue := make(chan string, len(graph))

	queue <- "you"

	//fmt.Println("size of graph:", len(graph))

	todo, err := do(isToDo, queue, graph)

	if err != nil {
		fmt.Printf("isTodo: %v\n", err)
	} else {
		fmt.Printf("isTodo: %s\n", todo)
	}
}

func isToDo(name string) bool {
	return strings.HasSuffix(name, "u")
	//return name == "bob"
}

func do(isToDo func(string) bool, queue chan string, graph map[string][]string) (string, error) {
	var find string
	isVerified := make(map[string]bool)

	for i := 0; i < len(graph); i++ {
		nameQ := <-queue

		isVerified[nameQ] = true
		//fmt.Printf("from %s graph: %v\n", nameQ, graph[nameQ])

		for _, v := range graph[nameQ] {
			if isVerified[v] {
				continue
			}
			isVerified[v] = true

			if isToDo(v) {
				find = v

				return find, nil
			}
			//fmt.Println("is 1 v", v)

			queue <- v

			continue
		}
	}

	close(queue)

	return find, errors.New("нету в сети")
}
