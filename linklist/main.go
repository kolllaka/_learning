package main

import (
	"fmt"
)

type Element interface {
}
type LinkList struct {
	Data Element   // Поле данных
	Nest *LinkList // Поле указателя
}

func (head *LinkList) GetLength() int {
	point := head
	var length int
	for point != nil {
		length++
		point = point.Nest

	}
	return length

} // Получить длину связанного списка

func (head *LinkList) Traverse() {
	point := head
	for point != nil {
		fmt.Println(point.Data) // Выходные данные
		point = point.Nest

	}
} // Обход связанного списка
func (head *LinkList) Delete(index int) Element {
	if index < 0 || index > head.GetLength() {
		fmt.Println("Индекс недопустим")
		return false

	} else {
		point := head
		for i := 0; i < index; i++ {
			point = point.Nest // Переместить позицию

		}
		point.Nest = point.Nest.Nest
		data := point.Nest.Data
		return data
	}

} // Удалить узел в указанной позиции
func (head *LinkList) Insert(index int, data Element) {
	if index < 0 || index > head.GetLength() {
		fmt.Println("Незаконные данные")
	} else {
		point := head
		for i := 0; i < index; i++ {
			point = point.Nest // Переместить позицию
		}
		var NewLinkList LinkList // Новый узел
		NewLinkList.Data = data
		NewLinkList.Nest = point.Nest
		point.Nest = &NewLinkList
	}

} // Вставить данные по указанному индексу
func (head *LinkList) Search(data Element) {
	point := head
	index := 0
	for point.Nest != nil {
		if point.Data == data {
			fmt.Println(data, "show in at index: ", index)
			break
		} else {
			index++
			point = point.Nest
			if index > head.GetLength() {
				fmt.Println("no this data in the LinkList!!!")
				break
			}
			continue
		}
	}
} // Поиск данных
func (head *LinkList) GetData(index int) Element {
	point := head
	if index < 0 || index > head.GetLength() {
		return false

	} else {

		for i := 0; i < index; i++ {
			point = point.Nest // Переместить позицию
		}
		return point.Data
	}

} // Получить элемент в позиции индекса в LinkList
func (head *LinkList) Add(data Element) {
	point := head
	for point.Nest != nil {
		point = point.Nest // Переместить позицию

	}
	var NewLinklist LinkList
	point.Nest = &NewLinklist // newlinklist подключается после point.nest
	NewLinklist.Data = data

} // Добавить данные в список ссылок

func main() {
	var head LinkList = LinkList{Data: 0, Nest: nil}
	var array []Element
	for i := 0; i < 9; i++ {
		array = append(array, Element(i+1))
		head.Add(array[i]) // Добавить данные

	}
	head.Traverse() // Список ссылок
	head.Search(8)  // Запрос данных
	fmt.Println("length:", head.GetLength())
	fmt.Println("data:", head.GetData(9))
	head.Delete(2)  // Удалить данные с индексом 2
	head.Traverse() // Обход данных

}
