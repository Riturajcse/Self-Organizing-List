package selfOrganizingList

import "testing"

func TestMtfPrepend(t *testing.T) {
	mList := &mtfList{}
	mList.Prepend(0)
	mList.Prepend(1)
	mList.Prepend(2)
	curr := mList.head
	for i := 2; curr != nil; curr, i = curr.next, i-1 {
		if curr.value != i {
			t.Error("Failing to Prepend in MTF List. Expecting: ", i, " got: ", curr.value)
		}
	}
}

func TestMtfAppend(t *testing.T) {
	mList := &mtfList{}
	mList.Append(0)
	mList.Append(1)
	mList.Append(2)
	curr := mList.head
	for i := 0; curr != nil; curr, i = curr.next, i+1 {
		if curr.value != i {
			t.Error("Failing to Append in MTF List. Expecting: ", i, " got: ", curr.value)
		}
	}
}

func TestMtfMoveAfter(t *testing.T) {
	mList := &mtfList{}
	mList.Append(0)
	mList.Append(1)
	mList.Append(2)
	mList.MoveAfter(0, 2)
	mList.MoveAfter(1, 2)
	mList.MoveAfter(1, 1) // Same value shouldn't do anything
	mList.MoveAfter(3, 1) // Invalid value to move shouldn't do anything
	mList.MoveAfter(1, 3) // Invalid after value shouldn't do anything
	curr := mList.head
	for i := 2; curr != nil; curr, i = curr.next, i-1 {
		if curr.value != i {
			t.Error("Failing to MoveAfter in MTF List. Expecting: ", i, " got: ", curr.value)
		}
	}
}

func TestMtfMoveBefore(t *testing.T) {
	mList := &mtfList{}
	mList.Append(0)
	mList.Append(1)
	mList.Append(2)
	mList.MoveBefore(2, 0)
	mList.MoveBefore(1, 0)
	mList.MoveBefore(1, 1) // Same value shouldn't do anything
	mList.MoveBefore(3, 1) // Invalid value to move shouldn't do anything
	mList.MoveBefore(1, 3) // Invalid before value shouldn't do anything
	curr := mList.head
	for i := 2; curr != nil; curr, i = curr.next, i-1 {
		if curr.value != i {
			t.Error("Failing to MoveBefore in MTF List. Expecting: ", i, " got: ", curr.value)
		}
	}
}

func TestMtfMoveToBack(t *testing.T) {
	mList := &mtfList{}
	mList.Append(0)
	mList.Append(1)
	mList.Append(2)
	mList.MoveToBack(1)
	mList.MoveToBack(0)
	mList.MoveToBack(3) // Invalid value to move shouldn't do anything
	curr := mList.head
	for i := 2; curr != nil; curr, i = curr.next, i-1 {
		if curr.value != i {
			t.Error("Failing to MoveToBack in MTF List. Expecting: ", i, " got: ", curr.value)
		}
	}
}

func TestMtfFind(t *testing.T) {
	mList := &mtfList{}
	mList.Append(0)
	mList.Append(1)
	mList.Append(4)
	mList.Append(9)
	curr := mList.head
	for i := 0; curr != nil; curr, i = curr.next, i+1 {
		_, temp := mList.Find(i*i), mList.Find(i*i)
		// Here when we try to find the same element second time this should be present at the HEAD
		if temp != 0 {
			t.Error("Failed to follow MTF List criteria. Expecting: ", 0, " got: ", temp)
		}
	}
}
