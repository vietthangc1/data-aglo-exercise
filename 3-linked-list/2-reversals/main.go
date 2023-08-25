package main

import "log"

// https://leetcode.com/problems/reverse-linked-list-ii/

type Node struct {
	Value int
	Next  *Node
}

var linkedList = &Node{
	Value: 1,
	Next: &Node{
		Value: 2,
		Next: &Node{
			Value: 3,
			Next: &Node{
				Value: 4,
				Next: &Node{
					Value: 5,
					Next:  nil,
				},
			},
		},
	},
}

func solution(list *Node, left, right int) *Node {
	if left < 0 || right <= left {
		return nil
	}

	currentIndex := 1
	currentNode := list
	newNode := &Node{}
	currentNewNode := newNode.Next

	for currentIndex < left {
		currentNode = currentNode.Next
		newNode.Value = currentNode.Value
		currentNewNode = newNode.Next
		currentIndex++
	}
	return currentNewNode
}

func reverse(list *Node) *Node {
	if list.Next == nil {
		return list
	}

	newList := &Node{}

	currentNode := list

	for currentNode != nil {
		newList = &Node{
			Value: currentNode.Value,
			Next:  newList,
		}
		currentNode = currentNode.Next
	}
	return newList
}

func copyList(list *Node) *Node {
	newNode := &Node{
		Value: list.Value,
		Next: &Node{},
	}
	nextNode := newNode.Next

	currentNode := list
	for currentNode != nil {
		nextNode.Value = currentNode.Value
		nextNode = nextNode.Next
		currentNode = currentNode.Next
	}
	
	newNode = newNode.Next
	log.Println(newNode)
	return newNode
}

func main() {
	copyList(linkedList)
}
