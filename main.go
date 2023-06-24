package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	util "github.com/rizkiamr/go-singly-linked-list/util"
)

type person struct {
	name string
	id   int
}

type node struct {
	next   *node
	person person
}

func CreateLinkedList(head *node) (*node, error) {
	var tempId, id int
	var nextchar byte

	if head != nil {
		head, _ = ShowLinkedList(head)
	}

	for {
		id = 0
		tempId = 0
		fmt.Print(" Input id : ")
		_, err := fmt.Scanf("%d%c", &tempId, &nextchar)
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

		n, _ := IsIdExist(head, tempId)

		if n {
			fmt.Println(" ERROR ID ALREADY EXIST")
			continue
		}
		id = tempId
		break
	}

	fmt.Print(" Input name : ")
	nameReader := bufio.NewReader(os.Stdin)
	name, _ := nameReader.ReadString('\n')

	newNode := node{
		person: person{
			id:   id,
			name: name,
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
			fmt.Printf(" id: %d, name: %s", ptr.person.id, ptr.person.name)
			ptr = ptr.next
		}
	} else {
		return head, errors.New(" Linked List is Empty")
	}
	return head, nil
}

func IsIdExist(ptr *node, id int) (bool, error) {
	if ptr != nil {
		if ptr.next != nil {
			for ptr.person.id != id {
				ptr = ptr.next
			}
			if ptr.person.id != id {
				return false, nil
			} else {
				return true, nil
			}
		} else {
			if ptr.person.id != id {
				return false, nil
			} else {
				return true, nil
			}
		}
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

func UpdateNodeById(head *node) (*node, error) {

	if head == nil {
		return head, errors.New(" Linked List is Empty")
	}

	head, err := ShowLinkedList(head)
	if err != nil {
		fmt.Print(err)
	}

	var tempId, id int
	var nextchar byte

	for {
		id = 0
		tempId = 0
		fmt.Print(" Input id : ")
		_, err := fmt.Scanf("%d%c", &tempId, &nextchar)
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

		n, _ := IsIdExist(head, tempId)

		if n {
			id = tempId
			break
		}
		fmt.Println(" ERROR ID DOES NOT EXIST")
		continue
	}

	fmt.Print(" Input name : ")
	nameReader := bufio.NewReader(os.Stdin)
	name, _ := nameReader.ReadString('\n')

	ptr := head

	for ptr.person.id != id {
		ptr = ptr.next
	}

	ptr.person.name = name

	return head, nil
}

func DeleteNodeById(head *node) (*node, error) {

	if head == nil {
		return head, errors.New(" Linked List is Empty")
	}

	head, err := ShowLinkedList(head)
	if err != nil {
		fmt.Println(err)
	}

	var tempId, id int
	var nextchar byte

	for {
		id = 0
		tempId = 0
		fmt.Print(" Input id : ")
		_, err := fmt.Scanf("%d%c", &tempId, &nextchar)
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

		n, _ := IsIdExist(head, tempId)

		if n {
			id = tempId
			break
		}
		fmt.Println(" ERROR ID DOES NOT EXIST")
		continue
	}

	var preptr *node
	ptr := head

	if ptr.next != nil {
		for ptr.person.id != id {
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
	for option != 8 {
		fmt.Print(" ***** Main Menu *****")
		fmt.Print("\n 1. Print hello world")
		fmt.Print("\n 2. Add a New Node")
		fmt.Print("\n 3. Show Linked List")
		fmt.Print("\n 4. Update a Node by Id")
		fmt.Print("\n 5. Delete a Node by Id")
		fmt.Print("\n 6. Delete Entire Linked List")
		fmt.Print("\n 7. Clear Screen")
		fmt.Print("\n 8. Exit")
		fmt.Print("\n Input your option : ")
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
				fmt.Println(" SUCCESSFULLY CREATED A NODE")
			case 3:
				head, err = ShowLinkedList(head)
				if err != nil {
					fmt.Println(err)
				}
			case 4:
				head, err = UpdateNodeById(head)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println(" SUCCESSFULLY UPDATED A NODE")
			case 5:
				head, err = DeleteNodeById(head)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println(" SUCCESSFULLY DELETED A NODE")
			case 6:
				head, err = DeleteEntireLinkedList(head)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println(" SUCCESSFULLY DELETED ENTIRE LINKED LIST")
			case 7:
				util.ClearScreen()
			case 8:
				continue
			default:
				fmt.Print(" ERROR INVALID INPUT")
			}
		}
	}
}
