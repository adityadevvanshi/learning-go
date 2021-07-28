package linkedlist

import (
	"fmt"

	"go.uber.org/zap"
)

var Logger *zap.Logger

type Node struct {
	data     int
	nextnode *Node
}
type NodeList struct {
	Head *Node
}

func (nodelist *NodeList) InsertAtBeginning(value int) {
	if nodelist == nil {
		Logger.Warn("NodeList Pointer is nil in InsertAtBeginning Method", zap.Int("Value to be inserted", value))
	} else {
		nodelist.Head = &Node{data: value, nextnode: nodelist.Head}
		Logger.Info("Insertion at Beginning Success!", zap.Int("inserted value", value))
	}
}

func (head *Node) PrintLinkedList() {
	if head == nil {
		Logger.Warn("head is nil in PrintLinkedList Method")
	} else {
		for node := head; node != nil; node = node.nextnode {
			fmt.Println(node.data)
		}
	}
}
