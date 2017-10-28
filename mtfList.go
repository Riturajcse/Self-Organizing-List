package selfOrganizingList

import "fmt"

type mtfNode struct {
	value int
	next  *mtfNode
}

type mtfList struct {
	head *mtfNode
}

func newMtfNode(value int) *mtfNode {
	return &mtfNode{value, nil}
}

func (sl *mtfList) Prepend(value int) {
	n := newMtfNode(value)
	n.next = sl.head
	sl.head = n
}

func (sl *mtfList) Append(value int) {
	n := newMtfNode(value)

	if sl.head == nil {
		sl.head = n
		return
	}

	curr := sl.head
	for ; curr.next != nil; curr = curr.next {
	}

	curr.next = n
}

func (sl *mtfList) MoveAfter(value int, after int) {
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

func (sl *mtfList) MoveBefore(value int, before int) {
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

func (sl *mtfList) MoveToBack(value int) {
	if find := sl.getNode(value); find != nil {
		sl.delElement(value)
		sl.Append(value)
	}
}

func (sl *mtfList) MoveToFront(value int) {
	if find := sl.getNode(value); find != nil {
		sl.delElement(value)
		sl.Prepend(value)
	}
}

func (sl *mtfList) Find(find int) int {
	if sl.head == nil {
		return -1
	}
	curr := sl.head
	if curr.value == find {
		return 0
	}
	for i := 0; curr != nil; curr, i = curr.next, i+1 {
		if curr.value == find {
			sl.delElement(find)
			sl.Prepend(find)
			return i
		}
	}
	return -1
}

func (sl *mtfList) delElement(del int) {
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

func (sl *mtfList) getNode(value int) *mtfNode {
	for curr := sl.head; curr != nil; curr = curr.next {
		if curr.value == value {
			return curr
		}
	}

	return nil
}

func (sl *mtfList) getNodeWithPrevious(value int) (*mtfNode, *mtfNode) {
	var prev *mtfNode
	for curr := sl.head; curr != nil; curr = curr.next {
		if curr.value == value {
			return curr, prev
		}
		prev = curr
	}

	return nil, nil
}

func (sl *mtfList) PrintList() {
	for curr := sl.head; curr != nil; curr = curr.next {
		fmt.Print(curr.value, " ")
	}
}
