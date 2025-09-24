package stack

import (
	"slices"
	"testing"
)

func TestSnapshot(t *testing.T) {

	st := new(Stack[int])

	st.Push(1)
	st.Push(2)
	st.Push(3)

	switch sn := st.Snapshot(); {
	case len(sn) != 3:
		t.Fatal("snapshot length not equal the origin")
	case sn[0] != 1:
		t.Fatalf("snapshot element at index %d not equal origin", 0)
	case sn[1] != 2:
		t.Fatalf("snapshot element at index %d not equal origin", 1)
	case sn[2] != 3:
		t.Fatalf("snapshot element at index %d not equal origin", 2)
	}
}

func TestCloneDistinction(t *testing.T) {

	st := new(Stack[int])

	st.Push(1)
	st.Push(2)
	st.Push(3)

	clone := st.Clone()

	switch {
	case !slices.Equal(st.Snapshot(), clone.Snapshot()):
		t.Fatal("clone inequal")
	}

	st.Pop()
	st.Pop()

	switch {
	case st.Mass() == clone.Mass():
		t.Fatal("")
	case slices.Equal(st.Snapshot(), clone.Snapshot()):
		t.Fatal("clone affected by origin")
	}
}

func TestEquivalent(t *testing.T) {

	st := new(Stack[int])

	st.Push(1)
	st.Push(2)
	st.Push(3)

	clone := st.Clone()

	for !st.IsEmpty() {

		or, _ := st.Pop()
		cl, _ := clone.Pop()

		switch {
		case or != cl:
			t.Fatalf("inequavalent of stack head")
		}
	}
}
