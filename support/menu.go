package support

import (
	"bufio"
	"fmt"
	"hash_table/collection"
	"os"
	"strings"
)

func Menu(items *collection.HashTable) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("                                   ")
	fmt.Println("+ ------------------------------- +")
	fmt.Println("|        Chose option             |")
	fmt.Println("+ ------------------------------- +")
	fmt.Println("                                   ")

	opt, _ := getInput("a - Add Name, d - Delete Name, s - Search, p - Print: ", reader)

	switch opt {
	case "a":
		add(items, reader)
		Menu(items)
	case "d":
		remove(items, reader)
		Menu(items)
	case "s":
		search(items, reader)
		Menu(items)
	case "p":
		items.Print()
		Menu(items)
	default:
		fmt.Println("Invalid option ...")
		Menu(items)
	}
}

func add(items *collection.HashTable, reader *bufio.Reader) {
	name, _ := getInput("Name: ", reader)

	items.Add(name)
}

func remove(items *collection.HashTable, reader *bufio.Reader) {
	name, _ := getInput("Name: ", reader)

	items.Delete(name)
}

func search(items *collection.HashTable, reader *bufio.Reader) {
	name, _ := getInput("Name: ", reader)

	node := items.Find(name)

	if node.Value != "" {
		fmt.Println("The found name is: ", node.Value)

		return
	}

	fmt.Println("The given name was not found...")
}

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}
