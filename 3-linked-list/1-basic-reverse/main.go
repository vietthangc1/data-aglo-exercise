package main

import (
	"log"
)

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

func bruteForce(list *Node) *Node {
	if list == nil {
		return nil
	}

	listValue := []int{}

	currentNode := list
	for currentNode != nil {
		listValue = append(listValue, currentNode.Value)
		currentNode = currentNode.Next
	}

	log.Println(listValue)

	newLinkedList := &Node{}

	for _, e := range listValue {
		currentNode := &Node{
			Value: e,
			Next:  newLinkedList,
		}
		newLinkedList = currentNode
	}

	log.Println(newLinkedList)
	return newLinkedList
}

func returnNewNode(list *Node) *Node {
	if list == nil {
		return nil
	}

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
	log.Println(newList)
	return newList
}

func optimalCode(list *Node) {
	prevNode := &Node{}
	currentNode := list

	for currentNode != nil {
		next := currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = next
	}

	log.Println(prevNode)
}

func main() {
	bruteForce(linkedList)
	returnNewNode(linkedList)
	optimalCode(linkedList)
}
