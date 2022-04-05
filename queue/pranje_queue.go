package queue

import "fmt"

var queue IQueue = New()

func Push(user string) bool {
	iter := queue.Iterator()
	for iter.HasNext() {
		if iter.Next().Item() == user {
			return false
		}
	}
	queue.Enqueue(user)
	return true
}

func Pop(user string) string {
	iter := queue.Iterator()
	if iter.HasNext() {
		if queue.Peek().Item() == user {
			queue.Dequeue()
			iter := queue.Iterator()
			if iter.HasNext() {
				return iter.Next().Item()
			} else {
				return "last"
			}
		} else {
			return "denied"
		}
	} else {
		return "empty"
	}
}

func PrintQueue() string {
	resp := ""
	iter := queue.Iterator()
	for iter.HasNext() {
		resp += fmt.Sprintf("%s\n", iter.Next().Item())
	}
	if len(resp) != 0 {
		return resp
	} else {
		return "Laundry queue is empty"
	}

}
