package main

import (
	"fmt"
	"math"
)

type Graph struct {
	Cost int
	Name string
}

type toDoGraph struct {
	cost   int
	parent string
	done   bool
}

func main() {
	graph := make(map[string][]Graph)

	// graph = map[string][]Graph{
	// 	"book": {
	// 		{5, "record"},
	// 		{0, "poster"},
	// 	},
	// 	"poster": {
	// 		{35, "drum"},
	// 		{30, "guitar"},
	// 	},
	// 	"record": {
	// 		{20, "drum"},
	// 		{15, "guitar"},
	// 	},
	// 	"drum": {
	// 		{10, "piano"},
	// 	},
	// 	"guitar": {
	// 		{20, "piano"},
	// 	},
	// }

	// graph = map[string][]Graph{
	// 	"start": {
	// 		{6, "a"},
	// 		{2, "b"},
	// 	},
	// 	"a": {
	// 		{1, "fin"},
	// 	},
	// 	"b": {
	// 		{3, "a"},
	// 		{5, "fin"},
	// 	},
	// }

	// graph = map[string][]Graph{
	// 	"start": {
	// 		{2, "a"},
	// 		{5, "b"},
	// 	},
	// 	"a": {
	// 		{7, "c"},
	// 		{8, "b"},
	// 	},
	// 	"b": {
	// 		{2, "c"},
	// 		{4, "d"},
	// 	},
	// 	"c": {
	// 		{1, "fin"},
	// 	},
	// 	"d": {
	// 		{3, "fin"},
	// 	},
	// }

	graph = map[string][]Graph{
		"start": {
			{10, "a"},
		},
		"a": {
			{20, "b"},
		},
		"b": {
			{30, "fin"},
			{1, "c"},
		},
		"c": {
			{1, "a"},
		},
	}

	//fmt.Println(graph)

	Run("start", "fin", graph)
}

func Run(start, fin string, graph map[string][]Graph) {
	toDoTable := make(map[string]toDoGraph)

	for _, v := range graph[start] {
		toDoTable[v.Name] = toDoGraph{
			cost:   v.Cost,
			parent: start,
		}
	}

	toDoTable = StepAlgoritm(start, graph, toDoTable)
	//fmt.Println(toDoTable)

	printPath(fin, toDoTable)
	fmt.Printf("за - %d\n", toDoTable[fin].cost)
}

func printPath(fin string, todoTable map[string]toDoGraph) {

	fmt.Print(fin)
	item, ok := todoTable[fin]
	if !ok {
		fmt.Println()
		return
	}

	fmt.Print("->")

	printPath(item.parent, todoTable)
}

func StepAlgoritm(start string, graph map[string][]Graph, toDoTable map[string]toDoGraph) map[string]toDoGraph {
	// шаг 1 Найти самый дешёвый необработанный узел
	minCost := math.MaxInt
	var minName string

	for k, v := range toDoTable {
		if v.done {
			continue
		}

		if minCost > v.cost {
			minCost = v.cost
			minName = k
		}
	}
	//fmt.Println(minName)

	// Если minName не найденно, то мы всё обошли
	if _, ok := toDoTable[minName]; !ok {
		return toDoTable
	} else {
		update := toDoTable[minName]
		toDoTable[minName] = toDoGraph{
			update.cost,
			update.parent,
			true,
		}
	}

	// шаг 2 Обновить стоимость всех соседей
	for _, v := range graph[minName] {
		if _, ok := toDoTable[v.Name]; !ok {
			toDoTable[v.Name] = toDoGraph{
				cost:   v.Cost + toDoTable[minName].cost,
				parent: minName,
			}
		}

		cost := v.Cost + toDoTable[minName].cost
		if toDoTable[v.Name].cost > cost {
			toDoTable[v.Name] = toDoGraph{
				cost:   cost,
				parent: minName,
			}
		}

	}
	//fmt.Println(toDoTable)

	StepAlgoritm(minName, graph, toDoTable)

	return toDoTable
}
