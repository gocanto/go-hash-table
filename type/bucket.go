package _type

import "fmt"

type Bucket struct {
	Head *Node
}

func (bucket *Bucket) Insert(value string) {
	node := &Node{Value: value}

	node.Next = bucket.Head
	bucket.Head = node
}

func (bucket *Bucket) Delete(value string) {
	if bucket.Head.Value == value {
		bucket.Head = bucket.Head.Next

		return
	}

	previous := bucket.Head

	for previous.Next != nil {
		if previous.Next.Value == value {
			previous.Next = previous.Next.Next
		}

		previous = previous.Next
	}
}

func (bucket Bucket) Find(value string) Node {
	node := bucket.Head

	for node != nil {
		if node.Value == value {
			return *node
		}

		node = node.Next
	}

	return Node{}
}

func (bucket Bucket) Has(value string) bool {
	node := bucket.Find(value)

	return node.Value != ""
}

func (bucket Bucket) Print() {
	for e := bucket.Head; e != nil; e = e.Next {
		fmt.Println("Value:", e.Value)
	}
}
