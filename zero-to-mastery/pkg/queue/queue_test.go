package queue

import "testing"

func TestCreateQueue(t *testing.T) {
	q := CreateQueue(3) // create Queue with capacity 3
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect queue element count : %v, want %v", len(q.items), i)
		}
		if !q.AppendToQueue(i) {
			t.Errorf("Failed to append item %v to queue", i)
		}
	}
	if q.AppendToQueue(4) {
		t.Errorf("Should not be able to add to a full queue")
	}
}

func TestPopFirstFromQueue(t *testing.T) {
	q := CreateQueue(3)
	for i := range 3 {
		q.AppendToQueue(i)
	}

	for i := range 3 {
		item, ok := q.PopFirstFromQueue()
		if !ok {
			t.Errorf("Should be able to get item from queue")
		}
		if item != i {
			t.Errorf("Item %v should be equal to %v", item, i)
		}
	}
	// Queue should be empty by this point
	item, ok := q.PopFirstFromQueue()
	if ok {
		t.Errorf("Should not be any more items in queue, got: %v", item)
	}

}
