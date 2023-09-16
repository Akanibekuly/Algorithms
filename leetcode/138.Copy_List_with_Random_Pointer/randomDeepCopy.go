package copylistwithrandompointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	for p := head; p != nil; p = p.Next.Next {
		p.Next = &Node{p.Val, p.Next, nil}
	}
	for p := head; p != nil; p = p.Next.Next {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
	}
	headOrigin, headCopy := Node{}, Node{}
	prevOrigin, prevCopy := &headOrigin, &headCopy
	for p := head; p != nil; p = p.Next.Next {
		prevOrigin, prevOrigin.Next = p, p
		prevCopy, prevCopy.Next = p.Next, p.Next
	}
	prevOrigin.Next = nil
	return headCopy.Next
}

func copyRandomListHash(head *Node) *Node {
	oldToCopy := make(map[*Node]*Node)
	for t := head; t != nil; t = t.Next {
		oldToCopy[t] = &Node{Val: t.Val}
	}

	for old := head; old != nil; old = old.Next {
		cp := oldToCopy[old]
		cp.Next = oldToCopy[old.Next]
		cp.Random = oldToCopy[old.Random]
	}

	return oldToCopy[head]
}
