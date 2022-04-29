package main

import (
	"fmt"
)

type Matrix struct {
	Cell [][]Item
}

type Item struct {
	num       int
	subString string
}

func main() {
	word1 := "blue"
	word2 := "clues"

	matrix := createMatrix(len(word1)+1, len(word2)+1)

	maxSubString := matrix.subString(word1, word2)
	fmt.Println(maxSubString)
	matrix.printMatrix()

	maxSubSequence := matrix.subSequence(word1, word2)
	fmt.Println(maxSubSequence)
	matrix.printMatrix()
}

func createMatrix(rows, cols int) Matrix {
	cell := make([][]Item, rows)

	for i := range cell {
		cell[i] = make([]Item, cols)
	}

	return Matrix{
		Cell: cell,
	}
}

func (m *Matrix) updateMatrix(row, col int, item Item) {
	m.Cell[row][col] = item
}

func (m *Matrix) printMatrix() {
	for i, row := range m.Cell {
		if i > 0 {
			fmt.Println(row[1:])
		}
	}

	fmt.Println()
}

func (m *Matrix) subString(word1, word2 string) Item {
	var maxSubString Item

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				m.Cell[i][j].num = m.Cell[i-1][j-1].num + 1
				m.Cell[i][j].subString = m.Cell[i-1][j-1].subString + string(word1[i-1])

				if maxSubString.num < m.Cell[i][j].num {
					maxSubString = m.Cell[i][j]
				}
			} else {
				m.Cell[i][j] = m.Cell[i-1][j]

				if m.Cell[i][j].num < m.Cell[i][j-1].num {
					m.Cell[i][j] = m.Cell[i][j-1]
				}
			}
		}
	}

	return maxSubString
}

func (m *Matrix) subSequence(word1, word2 string) Item {
	var maxSubSequence Item

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				m.Cell[i][j].num = m.Cell[i-1][j-1].num + 1
				m.Cell[i][j].subString = m.Cell[i-1][j-1].subString + string(word1[i-1])

				if maxSubSequence.num < m.Cell[i][j].num {
					maxSubSequence = m.Cell[i][j]
				}
			} else {
				m.Cell[i][j].num = 0
				m.Cell[i][j].subString = ""
			}
		}
	}

	return maxSubSequence
}
