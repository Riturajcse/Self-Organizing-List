package selfOrganizingList

import "testing"

func TestCounterPrepend(t *testing.T) {
	cList := &counterList{}
	cList.Prepend(0)
	cList.Prepend(1)
	cList.Prepend(2)
	curr := cList.head
	for i := 2; curr != nil; curr, i = curr.next, i-1 {
		if curr.value != i {
			t.Error("Failing to Prepend in Counter List. Expecting: ", i, " got: ", curr.value)
		}
	}
}

func TestCounterAppend(t *testing.T) {
	cList := &counterList{}
	cList.Append(0)
	cList.Append(1)
	cList.Append(2)
	curr := cList.head
	for i := 0; curr != nil; curr, i = curr.next, i+1 {
		if curr.value != i {
			t.Error("Failing to Append in Counter List. Expecting: ", i, " got: ", curr.value)
		}
	}
}

func TestCounterFind(t *testing.T) {
	cList := &counterList{}
	cList.Append(0)
	cList.Append(1)
	cList.Append(2)
	cList.Append(3)
	cList.Find(2)
	cList.Find(2)
	cList.Find(2)
	cList.Find(2)
	cList.Find(3)
	cList.Find(3)
	cList.Find(3)
	cList.Find(1)
	cList.Find(1)
	curr := cList.head
	// Expected output based on number of counters is 2 3 1 0
	if curr.value != 2 {
		t.Error("Failed to follow Counter List criteria. Expecting: ", 2, " got: ", curr.value)
	} else if curr.next.value != 3 {
		t.Error("Failed to follow Counter List criteria. Expecting: ", 3, " got: ", curr.next.value)
	} else if curr.next.next.value != 1 {
		t.Error("Failed to follow Counter List criteria. Expecting: ", 1, " got: ", curr.next.next.value)
	} else if curr.next.next.next.value != 0 {
		t.Error("Failed to follow Counter List criteria. Expecting: ", 0, " got: ", curr.next.next.next.value)
	}
}
