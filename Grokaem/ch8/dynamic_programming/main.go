package main

import "fmt"

type Matrix struct {
	cell   [][]Item
	weight float64
	items  []Item
}

type Item struct {
	name   string
	cost   float64
	weight float64
}

func main() {
	shop := []Item{
		{"Magniphone", 3000, 4},
		{"Nout", 2000, 3},
		{"Guitar", 1500, 1},
		{"IPhone", 2000, 1},
	}
	var minWeight float64 = 1
	var maxWeight float64 = 4

	countCol := int(maxWeight / minWeight)
	matrix := createMatrix(len(shop), countCol)
	matrix.items = shop
	matrix.weight = minWeight

	result := matrix.DynamicProgramm()
	fmt.Printf("купить: %s на цену: %v\n", result.name, result.cost)
	//matrix.printMatrix()
}

func createMatrix(rows, cols int) *Matrix {
	cell := make([][]Item, rows)

	for i := range cell {
		cell[i] = make([]Item, cols)
	}

	return &Matrix{
		cell: cell,
	}
}

func (m *Matrix) updateMatrix(row, col int, value Item) {
	m.cell[row][col] = value
}

func (m *Matrix) printMatrix() {
	for i, v := range m.cell {
		fmt.Printf("%-10s%v\n", m.items[i].name, v)
	}

	fmt.Println()
}

func (m *Matrix) maxValue(row, col int) Item {
	var max2 Item
	var max1 Item

	if row > 1 {
		max1 = m.cell[row-1][col]
	}

	if col-int(m.items[row].weight/m.weight) > -1 && row > 1 {
		max2.cost = m.items[row].cost + m.cell[row-1][col-int(m.items[row].weight/m.weight)].cost
		max2.name = m.items[row].name + " " + m.cell[row-1][col-int(m.items[row].weight/m.weight)].name
	} else {
		max2.cost = m.items[row].cost
		max2.name = m.items[row].name
	}

	if max1.cost < max2.cost {
		return max2
	}

	return max1
}

func (m *Matrix) DynamicProgramm() Item {
	for i, v := range m.cell {
		for j := range v {
			if m.weight*float64(j+1) < m.items[i].weight {
				if i == 0 {
					m.updateMatrix(i, j, Item{
						name: "",
						cost: 0,
					})
					continue
				}

				m.updateMatrix(i, j, m.cell[i-1][j])
				continue
			}

			max := m.maxValue(i, j)
			m.updateMatrix(i, j, max)
		}
	}

	return m.cell[len(m.items)-1][len(m.cell[0])-1]
}
