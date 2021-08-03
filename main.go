package main

import (
	"fmt"
	"hash_table/collection"
)

func main() {
	table := collection.Init()
	table.Add("Gustavo")
	table.Print()

	fmt.Println("======")
	fmt.Println("value:", table.Find("Gustavo").Value)
	fmt.Println("======")
	table.Delete("Gustavo")
	fmt.Println("======")
	fmt.Println("value:", table.Find("Gustavo").Value)
}
