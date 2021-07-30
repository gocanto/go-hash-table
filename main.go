package main

import "fmt"

const Size = 7

type Node struct {
	key  string
	next *Node
}

type Bucket struct {
	head *Node
}

type HashTable struct {
	array [Size]*Bucket
}

func Init() *HashTable {
	table := &HashTable{}
	//fmt.Println(table.array)

	for i := range table.array {
		table.array[i] = &Bucket{}
	}

	return table
}

func (bucket *Bucket) Insert(seed string) {
	//target := bucket.Find(seed)
	node := &Node{key: seed}

	//if target.key != "" {
	//	node := &target
	//}

	node.next = bucket.head
	bucket.head = node
}

func (bucket *Bucket) Delete(seed string) {
	if bucket.head.key == seed {
		bucket.head = bucket.head.next

		return
	}

	previous := bucket.head

	for previous.next != nil {
		if previous.next.key == seed {
			previous.next = previous.next.next
		}

		previous = previous.next
	}
}

func (bucket Bucket) Find(seed string) Node {
	node := bucket.head

	for node != nil {
		if node.key == seed {
			return *node
		}

		node = node.next
	}

	return Node{}
}

func (bucket Bucket) Has(seed string) bool {
	node := bucket.Find(seed)

	return node.key != ""
}

func (bucket Bucket) Print() {
	//head := bucket.head

	for e := bucket.head; e != nil; e = e.next {
		fmt.Println(e.key)
	}
}

func (current *HashTable) Add(seed string) {
	index := hash(seed)
	current.array[index].Insert(seed)
}

func main() {
	//table := Init()
	bucket := Bucket{}
	bucket.Insert("Li")
	bucket.Insert("Gus")

	bucket.Delete("Gus")

	bucket.Print()
}
func hash(seed string) int {
	result := 0

	for _, letter := range seed {
		result += int(letter)
	}

	return result % Size
}
