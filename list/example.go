package main

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insert(data string) {

	newNode := Node{data: data, next: nil}
	list.head = &newNode
}

func (list *LinkedList) contains(data string) bool {

	contains := false
	current := list.head
	for current != nil {

		if current.data == data {
			contains = true
		}
	}

	return contains

}

func main() {

	list := LinkedList{}
	list.insert("value 1")
	list.contains("value 1")

}
