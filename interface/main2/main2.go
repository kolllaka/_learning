package main

import "fmt"

func main() {
	person := make(map[string]interface{}, 0)

	person["name"] = "Alice"
	person["age"] = 21
	person["height"] = 167.64

	for k, v := range person {
		switch v.(type) {
		case int:
			person[k] = v.(int) + 1
		case float64:
			person[k] = v.(float64) + 1
		}
	}
	//person["age"] = person["age"].(int) + 1

	fmt.Printf("%+v", person)
}
