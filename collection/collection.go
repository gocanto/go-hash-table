package collection

import "hash_table/type"

const Size = 10

type HashTable struct {
	items [Size]*_type.Bucket
}

func (current *HashTable) Add(seed string) {
	index := hash(seed)
	current.items[index].Insert(seed)
}

func (current HashTable) Find(seed string) _type.Node {
	index := hash(seed)
	return current.items[index].Find(seed)
}

func (current *HashTable) Delete(seed string) {
	index := hash(seed)
	current.items[index].Delete(seed)
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
