package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
	return nil
}

func makeListFromArray(array []int) *ListNode {

	first := ListNode{
		Val: array[len(array) - 1],
		Next: nil,
	}

	if (len(array) == 1) {
		return &first;
	}

	aux := &first
	for i := len(array) - 2; i >= 0; i-- {
		newNode := ListNode{
			Val: array[i],
			Next: nil,
		}

		aux.Next = &newNode;
		aux = &newNode;
	}

	return &first;
}

func printList(list *ListNode) {

	aux := list;

	for {
		if (aux == nil){
			break
		}
		fmt.Print(aux.Val)
		aux = aux.Next
	}

	fmt.Print("\r\n")
}

func main() {

	l1 := makeListFromArray([]int{3,4,2});
	printList(l1)

	l2 := makeListFromArray([]int{4,6,5});
	printList(l2)

	addTwoNumbers(l1,l2);

}