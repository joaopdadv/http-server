package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var aux *ListNode = nil
	var head *ListNode;

	l1Aux := l1;
	l2Aux := l2;

	sum := 0;

	// iterar até ambos nulos
	for l1Aux != nil || l2Aux != nil {
		
		// somar
		if(l1Aux != nil){
			sum += l1Aux.Val
			l1Aux = l1Aux.Next
		}
		if(l2Aux != nil){
			sum += l2Aux.Val
			l2Aux = l2Aux.Next
		}

		newNode := ListNode{
			Val: sum % 10,
			Next: nil,
		}

		// gravar o valor da unidade
		if (aux != nil) {
			aux.Next = &newNode;
		} else {
			head = &newNode;
		}
		aux = &newNode;

		sum /= 10;
	}

	// no fim, se resto, add resto
	if(sum > 0){
		// add to last node .next
		aux.Next = &ListNode{
			Val: sum,
			Next: nil,
		};
	}

	return head;
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

	l1 := makeListFromArray([]int{9,9,9,9,9,9,9});
	printList(l1)

	l2 := makeListFromArray([]int{9,9,9,9});
	printList(l2)

	sum := addTwoNumbers(l1,l2);
	printList(sum)
}