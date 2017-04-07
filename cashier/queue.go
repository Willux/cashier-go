package cashier

type Node interface{}

type workerNode struct {
	worker     Worker
	workerNode *next
}

type Queue interface {
	Enqueue(worker Worker)
	Deque() Worker
}

type workerQueue struct {
	head Node
	tail Node
}

func NewWorkerQueue() Queue {
	return &workerQueue{head: nil, tail: nil}
}

func (q *workerQueue) Enqueue(worker Worker) {
	if q.head == nil {
		q.head = &workerNode{worker: worker, next: nil}
		q.tail = q.head
	} else {
		q.tail.next = &workerNode{worker: worker, next: nil}
		q.tail = q.tail.next
	}
}

func (q *workerQueue) Deque() Worker {
	if q.head == nil {
		return nil
	}

	toRemove := q.head
	q.head = q.head.next

	// If we removed the last node,
	// the tail needs to be nil too.
	if q.head == nil {
		q.tail = nil
	} else {
		toRemove.next = nil
	}

	return toRemove.worker
}
