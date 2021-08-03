package _type

import "fmt"

type Bucket struct {
	Head *Node
}

func (bucket *Bucket) Insert(seed string) {
	node := &Node{Value: seed}

	node.Next = bucket.Head
	bucket.Head = node
}

func (bucket *Bucket) Delete(seed string) {
	if bucket.Head.Value == seed {
		bucket.Head = bucket.Head.Next

		return
	}

	previous := bucket.Head

	for previous.Next != nil {
		if previous.Next.Value == seed {
			previous.Next = previous.Next.Next
		}

		previous = previous.Next
	}
}

func (bucket Bucket) Find(seed string) Node {
	node := bucket.Head

	for node != nil {
		if node.Value == seed {
			return *node
		}

		node = node.Next
	}

	return Node{}
}

func (bucket Bucket) Has(seed string) bool {
	node := bucket.Find(seed)

	return node.Value != ""
}

func (bucket Bucket) Print() {
	for e := bucket.Head; e != nil; e = e.Next {
		fmt.Println("Value:", e.Value)
	}
}
