package piscine

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		return
	}

	curentnode := l.Head
	for curentnode.Next != nil {
		curentnode = curentnode.Next
	}
	curentnode.Next = n
}
