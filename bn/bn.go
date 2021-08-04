package main

import (
	"errors"
	"fmt"
)

type Singlink interface {
	orderInsert(node *Node)
	pop(id int) (node *Node, err error)
	list()
}
type Node struct {
	id   int
	name string
	next *Node
}
type SingleLk struct {
	head *Node
}

func NewSinglink(head *Node) Singlink {
	return &SingleLk{head}
}

func (singleLink *SingleLk) orderInsert(node *Node) {
	curNode := singleLink.head
	if curNode.id == 0 {
		singleLink.head = node
	} else {
		if curNode.id > node.id {
			node.next = curNode
			singleLink.head = node
		} else {
			for {
				if curNode.next == nil {
					curNode.next = node
					break
				}
				if curNode.next.id > node.id {
					node.next = curNode.next
					curNode.next = node
					break
				}
				curNode = curNode.next
			}
		}
	}
}
func (singleLink *SingleLk) pop(id int) (node *Node, err error) {
	curNode := singleLink.head
	if curNode.id == id {
		singleLink.head = singleLink.head.next
		return curNode, nil
	}
	flag := false
	for {
		if curNode.next == nil {
			break
		}
		if curNode.next.id == id {
			flag = true
			selectNode := curNode.next
			curNode.next = curNode.next.next
			return selectNode, nil
		}
		curNode = curNode.next
	}
	if !flag {
		return nil, errors.New("node not found")
	}
	return
}

func (singleLink *SingleLk) list() {
	if singleLink.head == nil || singleLink.head.id == 0 {
		fmt.Println("link is empty")
	} else {
		curNode := singleLink.head
		for {
			if curNode == nil {
				break
			}
			fmt.Print(*curNode, "==>")
			curNode = curNode.next
		}
	}
	fmt.Println()

}

func main() {
	var originalHead = Node{
		id: 0,
	}
	var s = SingleLk{
		head: &originalHead,
	}
	var node1 = Node{
		id:   1,
		name: "001",
	}
	var node2 = Node{
		id:   2,
		name: "002",
	}
	var node3 = Node{
		id:   3,
		name: "003",
	}
	var node4 = Node{
		id:   4,
		name: "004",
	}
	var node5 = Node{
		id:   5,
		name: "005",
	}

	s.orderInsert(&node4)
	s.list()
	s.orderInsert(&node2)
	s.list()
	s.orderInsert(&node3)
	s.list()
	s.orderInsert(&node5)
	s.list()
	s.orderInsert(&node1)
	s.list()
	popNode, err := s.pop(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("pop node successfully:", *popNode)
	s.list()
	var head *Node
	bn := NewSinglink(head)
	bn.list()
}
