package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	util "github.com/rizkiamr/go-singly-linked-list/util"
)

type person struct {
	id int
}

type node struct {
	next   *node
	person person
}

func CreateLinkedList(head *node) (*node, error) {
	var tempNumber, number int
	var nextchar byte

	for {
		number = 0
		tempNumber = 0
		fmt.Print(" Input number : ")
		_, err := fmt.Scanf("%d%c", &tempNumber, &nextchar)
		if err != nil || nextchar != '\n' {
			reader := bufio.NewReader(os.Stdin)
			for {
				char, _ := reader.ReadByte()
				if char == '\n' {
					break
				}
			}
			fmt.Println(" ERROR INVALID INPUT")
			continue
		}
		number = tempNumber
		break
	}

	newNode := node{
		person: person{
			id: number,
		},
		next: nil,
	}

	ptr := head

	if ptr != nil {
		for (*ptr).next != nil {
			ptr = (*ptr).next
		}
		(*ptr).next = &newNode
	} else {
		head = &newNode
	}
	return head, nil
}

func ShowLinkedList(head *node) (*node, error) {
	ptr := head

	if ptr != nil {
		for ptr != nil {
			fmt.Printf(" %d", ptr.person.id)
			ptr = ptr.next
		}
	} else {
		return head, errors.New(" Linked List is Empty")
	}
	return head, nil
}

func IsNumberExist(ptr *node, number int) (bool, error) {
	if ptr != nil {
		for ptr.person.id != number {
			ptr = ptr.next
		}
		return ptr.person.id == number, nil
	} else {
		return false, nil
	}
}

func DeleteEntireLinkedList(head *node) (*node, error) {
	if head == nil {
		return head, errors.New(" Linked List is Empty")
	}

	head = nil

	return head, nil
}

func DeleteNodeByData(head *node) (*node, error) {
	var tempNumber, number int
	var nextchar byte

	if head == nil {
		return head, errors.New(" Linked List is Empty")
	}

	head, err := ShowLinkedList(head)
	if err != nil {
		fmt.Print(err)
	}

	for {
		number = 0
		tempNumber = 0
		fmt.Print("\n Input number : ")
		_, err := fmt.Scanf("%d%c", &tempNumber, &nextchar)
		if err != nil || nextchar != '\n' {
			reader := bufio.NewReader(os.Stdin)
			for {
				char, _ := reader.ReadByte()
				if char == '\n' {
					break
				}
			}
			fmt.Println(" ERROR INVALID INPUT")
			continue
		}

		n, _ := IsNumberExist(head, tempNumber)

		if n {
			number = tempNumber
			break
		}
		fmt.Println(" ERROR NUMBER NOT EXIST")
		continue
	}

	var preptr *node
	ptr := head

	if ptr.next != nil {
		for ptr.person.id != number {
			preptr = ptr
			ptr = ptr.next
		}

		if ptr.next != nil {
			preptr.next = ptr.next
		} else {
			preptr.next = nil
		}
	} else {
		head = nil
	}

	return head, nil
}

func main() {
	var head *node
	head = nil

	var option int
	var nextchar byte

	util.ClearScreen()
	for option != 6 {
		fmt.Print("\n ***** Main Menu *****")
		fmt.Print("\n 1. Print hello world")
		fmt.Print("\n 2. Add a New Node")
		fmt.Print("\n 3. Show Linked List")
		fmt.Print("\n 4. Delete a Node by Data")
		fmt.Print("\n 5. Delete Entire Linked List")
		fmt.Print("\n 6. Exit")
		fmt.Print("\n Input your number : ")
		_, err := fmt.Scanf("%1d%c", &option, &nextchar)

		if err != nil || nextchar != '\n' {
			reader := bufio.NewReader(os.Stdin)
			for {
				char, _ := reader.ReadByte()
				if char == '\n' {
					break
				}
			}
			fmt.Print(" ERROR INVALID INPUT")
		} else {
			switch option {
			case 1:
				fmt.Print(" Hello World!")
			case 2:
				head, _ = CreateLinkedList(head)
				fmt.Print(" SUCCESSFULLY CREATED A NODE")
			case 3:
				head, err = ShowLinkedList(head)
				if err != nil {
					fmt.Print(err)
				}
			case 4:
				head, err = DeleteNodeByData(head)
				if err != nil {
					fmt.Print(err)
					continue
				}
				fmt.Print(" SUCCESSFULLY DELETED A NODE")
			case 5:
				head, err = DeleteEntireLinkedList(head)
				if err != nil {
					fmt.Print(err)
					continue
				}
				fmt.Print(" SUCCESSFULLY DELETED ENTIRE LINKED LIST")
			case 6:
				continue
			default:
				fmt.Print(" ERROR INVALID INPUT")
			}
		}
	}
}
