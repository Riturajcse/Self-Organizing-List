package selfOrganizingList

import "fmt"

type counterNode struct {
	value   int
	counter int
	next    *counterNode
}

type counterList struct {
	head *counterNode
}

func newCounterNode(val int) *counterNode {
	return &counterNode{val, 0, nil}
}

func (sl *counterList) Prepend(value int) {
	n := newCounterNode(value)
	n.next = sl.head
	sl.head = n
}

func (sl *counterList) Append(value int) {
	n := newCounterNode(value)

	if sl.head == nil {
		sl.head = n
		return
	}

	curr := sl.head
	for ; curr.next != nil; curr = curr.next {
	}

	curr.next = n
}

func (sl *counterList) MoveAfter(value int, after int) {
	if value == after {
		return
	}

	valueNode := sl.getNode(value)
	afterNode := sl.getNode(after)
	if valueNode == nil || afterNode == nil {
		return
	}

	sl.delElement(valueNode.value)
	valueNode.next = afterNode.next
	afterNode.next = valueNode
}

func (sl *counterList) MoveBefore(value int, before int) {
	if value == before {
		return
	}

	valueNode := sl.getNode(value)
	beforeNode, prevBeforeNode := sl.getNodeWithPrevious(before)
	if valueNode == nil || beforeNode == nil {
		return
	}

	sl.delElement(valueNode.value)
	if prevBeforeNode == nil {
		sl.Prepend(valueNode.value)
		return
	}

	valueNode.next = beforeNode
	prevBeforeNode.next = valueNode
}

func (sl *counterList) MoveToBack(value int) {
	if find := sl.getNode(value); find != nil {
		sl.delElement(value)
		sl.Append(value)
	}
}

func (sl *counterList) MoveToFront(value int) {
	if find := sl.getNode(value); find != nil {
		sl.delElement(value)
		sl.Prepend(value)
	}
}

func (sl *counterList) Find(find int) int {
	if sl.head == nil {
		return -1
	}
	curr := sl.head
	if curr.value == find {
		curr.counter = curr.counter + 1
		return 0
	}
	for i := 0; curr != nil; curr, i = curr.next, i+1 {
		if curr.value == find {
			curr.counter = curr.counter + 1
			sl.reArrange(curr, find)
			return i
		}
	}
	return -1
}

func (sl *counterList) reArrange(curr *counterNode, val int) {
	currentNode := sl.head
	n := newCounterNode(val)
	n.counter = curr.counter
	sl.delElement(curr.value)
	if currentNode.counter <= curr.counter {
		n.next = sl.head
		sl.head = n
		return
	}
	for ; currentNode.next != nil; currentNode = currentNode.next {
		if currentNode.next.counter <= curr.counter {
			temp := currentNode.next
			currentNode.next = n
			n.next = temp
			break
		}
	}
}

func (sl *counterList) delElement(del int) {
	curr := sl.head
	if sl.head == nil {
		return
	}
	if curr.value == del {
		sl.head = curr.next
		return
	}

	for ; curr.next != nil; curr = curr.next {
		if curr.next.value == del {
			curr.next = curr.next.next
			break
		}
	}
}

func (sl *counterList) getNode(value int) *counterNode {
	for curr := sl.head; curr != nil; curr = curr.next {
		if curr.value == value {
			return curr
		}
	}

	return nil
}

func (sl *counterList) getNodeWithPrevious(value int) (*counterNode, *counterNode) {
	var prev *counterNode
	for curr := sl.head; curr != nil; curr = curr.next {
		if curr.value == value {
			return curr, prev
		}
		prev = curr
	}

	return nil, nil
}

func (sl *counterList) PrintList() {
	for curr := sl.head; curr != nil; curr = curr.next {
		fmt.Print(curr.value, " ")
	}
}
