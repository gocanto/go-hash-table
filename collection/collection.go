package collection

import "hash_table/type"

const Size = 10

type HashTable struct {
	items [Size]*_type.Bucket
}

func (current *HashTable) Add(value string) {
	index := hash(value)
	current.items[index].Insert(value)
}

func (current HashTable) Find(value string) _type.Node {
	index := hash(value)
	return current.items[index].Find(value)
}

func (current *HashTable) Delete(value string) {
	index := hash(value)
	current.items[index].Delete(value)
}

func (current HashTable) Print() {
	for i := range current.items {
		current.items[i].Print()
	}
}

func hash(seed string) int {
	result := 0

	for _, letter := range seed {
		result += int(letter)
	}

	return result % Size
}

func Init() *HashTable {
	collection := &HashTable{}

	for i := range collection.items {
		collection.items[i] = &_type.Bucket{}
	}

	return collection
}
