package collections

import "testing"

func TestAddElement(t *testing.T) {

	list := LinkedList{}
	list.add("test value")

	if list.head == nil {
		t.Fatalf(`LinkedList head reference not nil after invoking add!`)
	}
}
