package main

import (
	"fmt"
	"go.uber.org/zap"
)

var logger *zap.Logger

func Init(){
	logger_, _ := zap.NewDevelopment()
	defer logger_.Sync()
	logger=logger_
}
type Node struct {
	data                    int
	left_child, right_child *Node
}

func (node *Node) insert(value int) {
	fmt.Print(logger)
	if node.data < value {
		if node.right_child == nil {
			logger.Info("New Right Child Node has been created.",zap.Int("Node value",value),zap.Int("Parent value",node.data))
			node.right_child = &Node{data: value}
		} else {
			node.right_child.insert(value)
		}
	} else if node.data > value {
		if node.left_child == nil {
			logger.Debug("New Left Child Node has been created.",zap.Int("Node value",value),zap.Int("Parent value",node.data))
			node.left_child = &Node{data: value}
		} else {
			node.left_child.insert(value)
		}
	} else {
		logger.Warn("Cannot created new Node, value already exists!!",zap.Int("Node Value",value))
	}
}
func (n *Node) buildBST() {
	n.insert(10)
	n.insert(2)
	n.insert(30)
	n.insert(4)
	n.insert(1)
	n.insert(50)
}

func (n*Node) inOrderTraversal(){
	if n==nil {
		return
	}
	n.left_child.inOrderTraversal()
	fmt.Println(n.data)
	n.right_child.inOrderTraversal()
}
func main() {
	Init()
	root := &Node{data: 3}
	root.buildBST()
	root.inOrderTraversal()
}
