package priorityQueue

import (
	"testing"
	"fmt"
	"math/rand"
	"math"
)

type Something struct {
	msg      string
	priority int
}

func (x *Something) HigherPriorityThan(o Interface) bool {
	return x.priority > o.(*Something).priority
}

func TestPriorityQueue(t *testing.T) {

	const ITEMS = 500000

	pq := New()

	for i := 0; i < ITEMS; i++ {
		priority := rand.Intn(20)
		something := &Something{msg: fmt.Sprintf("Priority %d", priority), priority: priority}
		pq.Push(something)
	}

	// Take the items out; they arrive in decreasing priority order.
	expected := math.MaxInt64
	for i := 0; i < ITEMS; i++ {
		item := pq.Pop()
		if item == nil {
			t.Error("Expecting a value. Got nil")
		}
		if got := item.(*Something).priority; got >= expected {
			t.Errorf("Expecting a priority equal or lower than %d, got %d", expected, got)
		}
	}

	// Now, the queue is empty, so we expect a nil value
	if item := pq.Pop(); item != nil {
		t.Error("Expecting a nil value.")
	}
}
