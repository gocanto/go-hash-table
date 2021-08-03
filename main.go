package main

import (
	"fmt"
	"github.com/golang-room/hash-table/collection"
)

func main() {
	table := collection.Init()
	table.Add("Gustavo")
	table.Print()

	fmt.Println("======")
	fmt.Println("value:", table.Find("Gustavo").Key)
	fmt.Println("======")
	table.Delete("Gustavo")
	fmt.Println("======")
	fmt.Println("value:", table.Find("Gustavo").Key)
	//bucket := Bucket{}
	//bucket.Insert("Li")
	//bucket.Insert("Gus")
	//
	//bucket.Delete("Gus")
	//
	//bucket.Print()
}
