package selfOrganizingList

import (
	"testing"
)

func TestTransposePrepend(t *testing.T) {
	tList := &transposeList{}
	tList.Prepend(0)
	tList.Prepend(1)
	tList.Prepend(2)
	curr := tList.head
	for i := 2; curr != nil; curr, i = curr.next, i-1 {
		if curr.value != i {
			t.Error("Failing to Prepend in Transpose List. Expecting: ", i, " got: ", curr.value)
		}
	}
}

func TestTransposeAppend(t *testing.T) {
	tList := &transposeList{}
	tList.Append(0)
	tList.Append(1)
	tList.Append(2)
	curr := tList.head
	for i := 0; curr != nil; curr, i = curr.next, i+1 {
		if curr.value != i {
			t.Error("Failing to Append in Transpose List. Expecting: ", i, " got: ", curr.value)
		}
	}
}

func TestTransposeFind(t *testing.T) {
	tList := &transposeList{}
	tList.Append(0)
	tList.Append(1)
	tList.Append(2)
	tList.Append(3)
	tList.Find(3)
	tList.Find(3)
	tList.Find(2)
	tList.Find(1)
	curr := tList.head
	// Expected output is 0 3 1 2
	if curr.value != 0 {
		t.Error("Failed to follow Counter List criteria. Expecting: ", 0, " got: ", curr.value)
	} else if curr.next.value != 3 {
		t.Error("Failed to follow Counter List criteria. Expecting: ", 3, " got: ", curr.next.value)
	} else if curr.next.next.value != 1 {
		t.Error("Failed to follow Counter List criteria. Expecting: ", 1, " got: ", curr.next.next.value)
	} else if curr.next.next.next.value != 2 {
		t.Error("Failed to follow Counter List criteria. Expecting: ", 2, " got: ", curr.next.next.next.value)
	}
}
