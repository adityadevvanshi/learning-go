package binarysearchtree

import (
	"fmt"

	"go.uber.org/zap"
)

var Logger *zap.Logger

type Node struct {
	data                    int
	left_child, right_child *Node
}

func (root *Node) insert(value int) {
	if root.data < value {
		if root.right_child == nil {
			Logger.Info("Right Child Node has been created.", zap.Int("Key", value), zap.Int("Current Node Value", root.data))
			root.right_child = &Node{data: value}
		} else {
			Logger.Debug("insert function called for right subtree",zap.Int("Key", value),zap.Int("Current Node Value", root.data))
			root.right_child.insert(value)
		}
	} else if root.data > value {
		if root.left_child == nil {
			Logger.Info("Left Child Node has been created.", zap.Int("Key", value), zap.Int("Current Node value", root.data))
			root.left_child = &Node{data: value}
		} else {
			Logger.Debug("insert function called for left subtree",zap.Int("Key", value),zap.Int("Current Node Value", root.data))
			root.left_child.insert(value)
		}
	} else {
		Logger.Warn("Key already exists!! cannot create new node", zap.Int("Key", value),zap.Int("Current Node Value", root.data))
	}
}
func (root *Node) BuildBST() {
	root.insert(10)
	root.insert(2)
	root.insert(0)
	root.insert(30)
	root.insert(4)
	root.insert(1)
	root.insert(50)
}

func (root *Node) InOrderTraversal() {
	if root == nil {
		return
	}
	Logger.Debug("Entering Left Subtree",zap.Int("Current Node",root.data))
	root.left_child.InOrderTraversal()
	Logger.Debug("Exit Left Subtree",zap.Int("Current Node",root.data))
	fmt.Println(root.data)
	Logger.Debug("Entering Right Subtree",zap.Int("Current Node",root.data))
	root.right_child.InOrderTraversal()
	Logger.Debug("Entering Right Subtree",zap.Int("Current Node",root.data))
}
