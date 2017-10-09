package selfOrganizingList
import "fmt"

type transposeNode struct {
  value int
  next *transposeNode
}

type transposeList struct {
  head *transposeNode
}

func newTransposeNode (val int) *transposeNode {
  return &transposeNode{val, nil}
}

func (sl *transposeList) Prepend (value int) {
  n := newTransposeNode(value)
  n.next = sl.head
  sl.head = n
}

func (sl *transposeList) Append (value int) {
  n := newTransposeNode(value)

  if sl.head == nil {
    sl.head = n
    return 
  }

  curr := sl.head
  for ; curr.next != nil; curr = curr.next {
  }

  curr.next = n;
}

func (sl *transposeList) Find (find int) int {
  if sl.head == nil {
    return -1
  }
  curr := sl.head
  if curr.value == find {
    return 0
  }
  for i := 0;curr != nil; curr, i = curr.next, i + 1 {
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
  for ;curr != nil && curr != transNode;{
    prev = curr
    curr = curr.next 
  }
  tempData := prev.value
  prev.value = curr.value
  curr.value = tempData
}

func (sl *transposeList) PrintList () {
  for curr := sl.head; curr != nil; curr = curr.next {
    fmt.Print(curr.value, " ")
  }
}
