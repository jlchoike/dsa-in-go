package collections

type Element struct {
	data string
	next *Element
}

type LinkedList struct {
	head *Element
}

func (list *LinkedList) add(data string) {

	newElement := Element{data: data, next: nil}
	list.head = &newElement
}

func (list *LinkedList) contains(data string) bool {

	contains := false
	current := list.head
	for current != nil {

		if current.data == data {
			contains = true
		}
		current = current.next
	}

	return contains

}
