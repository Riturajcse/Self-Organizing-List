package selfOrganizingList

import "fmt"

type transposeNode struct {
	value int
	next  *transposeNode
}

type transposeList struct {
	head *transposeNode
}

func newTransposeNode(val int) *transposeNode {
	return &transposeNode{val, nil}
}

func (sl *transposeList) Prepend(value int) {
	n := newTransposeNode(value)
	n.next = sl.head
	sl.head = n
}

func (sl *transposeList) Append(value int) {
	n := newTransposeNode(value)

	if sl.head == nil {
		sl.head = n
		return
	}

	curr := sl.head
	for ; curr.next != nil; curr = curr.next {
	}

	curr.next = n
}

func (sl *transposeList) MoveAfter(value int, after int) {
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

func (sl *transposeList) MoveBefore(value int, before int) {
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

func (sl *transposeList) MoveToBack(value int) {
	if find := sl.getNode(value); find != nil {
		sl.delElement(value)
		sl.Append(value)
	}
}

func (sl *transposeList) Find(find int) int {
	if sl.head == nil {
		return -1
	}
	curr := sl.head
	if curr.value == find {
		return 0
	}
	for i := 0; curr != nil; curr, i = curr.next, i+1 {
		if curr.value == find {
			sl.transpose(curr)
			return i
		}
	}
	return -1
}

func (sl *transposeList) transpose(transNode *transposeNode) {
	curr := sl.head
	var prev *transposeNode
	for curr != nil && curr != transNode {
		prev = curr
		curr = curr.next
	}
	tempData := prev.value
	prev.value = curr.value
	curr.value = tempData
}

func (sl *transposeList) delElement(del int) {
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

func (sl *transposeList) getNode(value int) *transposeNode {
	for curr := sl.head; curr != nil; curr = curr.next {
		if curr.value == value {
			return curr
		}
	}

	return nil
}

func (sl *transposeList) getNodeWithPrevious(value int) (*transposeNode, *transposeNode) {
	var prev *transposeNode
	for curr := sl.head; curr != nil; curr = curr.next {
		if curr.value == value {
			return curr, prev
		}
		prev = curr
	}

	return nil, nil
}

func (sl *transposeList) PrintList() {
	for curr := sl.head; curr != nil; curr = curr.next {
		fmt.Print(curr.value, " ")
	}
}
