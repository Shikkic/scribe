package main

import (
	"errors"
	"fmt"
	"strings"
)

type Node struct {
	Key        string
	Value      string
	LeftChild  *Node
	RightChild *Node
}

func AppendNodeToRootNode(node *Node) {
	if RootNode == nil {
		RootNode = node
	}

	nodePointer := RootNode
	for nodePointer != nil {
		compareValue := strings.Compare(node.Key, nodePointer.Key)

		if compareValue == -1 {
			if nodePointer.LeftChild == nil {
				nodePointer.LeftChild = node
				return
			}
			nodePointer = nodePointer.LeftChild
		}

		if compareValue == 0 {
			nodePointer = node
			return
		}

		if compareValue == 1 {
			if nodePointer.RightChild == nil {
				nodePointer.RightChild = node
				return
			}
			nodePointer = nodePointer.RightChild
		}
	}
}

func GetValueFromRootNodeKey(searchKey string) (string, error) {
	if RootNode == nil {
		return "", errors.New("No keys in the keystore")
	}

	nodePointer := RootNode
	for nodePointer != nil {
		compareValue := strings.Compare(searchKey, nodePointer.Key)

		if compareValue == -1 {
			nodePointer = RootNode.LeftChild
		}

		if compareValue == 0 {
			return nodePointer.Value, nil
		}

		if compareValue == 1 {
			nodePointer = RootNode.RightChild
		}
	}

	return "", errors.New("Key was not found in keystore")
}

func RemoveKeyFromRootNode(searchKey string) error {
	if RootNode == nil {
		return errors.New("No keys in the keystore")
	}

	nodePointer := RootNode
	var foundNodePointer *Node

	for nodePointer != nil {
		compareValue := strings.Compare(searchKey, nodePointer.Key)
		fmt.Println(compareValue)

		if compareValue == -1 {
			nodePointer = nodePointer.LeftChild
		}

		if compareValue == 0 {
			foundNodePointer = nodePointer
			// Case 1: the end of the tree
			if nodePointer.LeftChild == nil && nodePointer.RightChild == nil {
				foundNodePointer = nil
				return nil
			}
			// Case 2: left child is not nil, but right is nil
			if nodePointer.LeftChild != nil && nodePointer.RightChild == nil {
				foundNodePointer = nodePointer.LeftChild
				return nil
			}
			// Case : right child is not nil, but left is nil
			if nodePointer.LeftChild == nil && nodePointer.RightChild != nil {
				foundNodePointer = nodePointer.RightChild
				return nil
			}

			if nodePointer.LeftChild != nil && nodePointer.RightChild != nil {
				greatestChildParent := nodePointer
				greatestChild := nodePointer.LeftChild
				// Now find the greatest value
				for greatestChild.RightChild != nil {
					greatestChildParent = greatestChild
					greatestChild = greatestChild.RightChild
				}

				if greatestChild.LeftChild != nil {
					greatestChildParent.RightChild = greatestChild.LeftChild
				} else {
					greatestChildParent.RightChild = nil
				}

				foundNodePointer.Key = greatestChild.Key
				foundNodePointer.Value = greatestChild.Value

				return nil
			}
		}

		if compareValue == 1 {
			nodePointer = nodePointer.RightChild
		}
	}

	return errors.New("Key not found in the keystore")
}

func PrintRootTree() {
	// Pre-order traversal
	printNode(RootNode)
}

func printNode(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("Key = %s, Value = %s\n", node.Key, node.Value)
	printNode(node.LeftChild)
	printNode(node.RightChild)
}
