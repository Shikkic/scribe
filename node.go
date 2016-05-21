package main

import (
	"errors"
	"fmt"
	"strings"
)

type Node struct {
	key        string
	value      string
	leftChild  *Node
	rightChild *Node
}

func AppendNodeToRootNode(node *Node) {
	if RootNode == nil {
		RootNode = node
	}

	nodePointer := RootNode
	for nodePointer != nil {
		compareValue := strings.Compare(node.key, nodePointer.key)

		if compareValue == -1 {
			if nodePointer.leftChild == nil {
				nodePointer.leftChild = node
				return
			}
			nodePointer = nodePointer.leftChild
		}

		if compareValue == 0 {
			nodePointer = node
			return
		}

		if compareValue == 1 {
			if nodePointer.rightChild == nil {
				nodePointer.rightChild = node
				return
			}
			nodePointer = nodePointer.rightChild
		}
	}
}

func GetValueFromRootNodeKey(searchKey string) (string, error) {
	if RootNode == nil {
		return "", errors.New("No keys in the keystore")
	}

	nodePointer := RootNode
	for nodePointer != nil {
		compareValue := strings.Compare(searchKey, nodePointer.key)

		if compareValue == -1 {
			nodePointer = RootNode.leftChild
		}

		if compareValue == 0 {
			return nodePointer.value, nil
		}

		if compareValue == 1 {
			nodePointer = RootNode.rightChild
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
		compareValue := strings.Compare(searchKey, RootNode.key)

		if compareValue == -1 {
			nodePointer = nodePointer.leftChild
		}

		if compareValue == 0 {
			foundNodePointer = nodePointer
			// Case 1: the end of the tree
			if nodePointer.leftChild == nil && nodePointer.rightChild == nil {
				foundNodePointer = nil
				return nil
			}
			// Case 2: left child is not nil, but right is nil
			if nodePointer.leftChild != nil && nodePointer.rightChild == nil {
				foundNodePointer = nodePointer.leftChild
				return nil
			}
			// Case : right child is not nil, but left is nil
			if nodePointer.leftChild == nil && nodePointer.rightChild != nil {
				foundNodePointer = nodePointer.rightChild
				return nil
			}

			if nodePointer.leftChild != nil && nodePointer.rightChild != nil {
				greatestChild := nodePointer
				// Now find the greatest value
				for greatestChild.rightChild != nil {
					greatestChild = greatestChild.rightChild
				}

				rightOfFoundNodePointer := foundNodePointer.rightChild
				foundNodePointer = foundNodePointer.leftChild
				greatestChild.rightChild = rightOfFoundNodePointer
			}
		}

		if compareValue == 1 {
			nodePointer = nodePointer.rightChild
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

	fmt.Printf("Key = %s, Value = %s\n", node.key, node.value)
	printNode(node.leftChild)
	printNode(node.rightChild)
}
