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

	resto := 0;
	// iterar até ambos nulos
	for {
		if(l1Aux == nil && l2Aux == nil){
			break
		}
	
		// somar
		l1Val := 0;
		l2Val := 0;

		if(l1Aux != nil){
			l1Val = l1Aux.Val
		}
		if(l2Aux != nil){
			l2Val = l2Aux.Val
		}

		sum := l1Val + l2Val + resto;
		resto = 0;
		
		// salvar o resto (se maior que 10 -> val - 10)
		if (sum >= 10){
			resto = 1; // resto 1 pois nunca vai ser mais que 18 a soma (9+9)
			sum = sum - 10;
		}

		newNode := ListNode{
			Val: sum,
			Next: nil,
		}

		// gravar o valor da unidad
		if (aux != nil) {
			aux.Next = &newNode;
		} else {
			head = &newNode;
		}
		aux = &newNode;

		if(l1Aux != nil){
			l1Aux = l1Aux.Next
		}
		if(l2Aux != nil){
			l2Aux = l2Aux.Next
		}
	}

	// no fim, se resto, add resto
	if(resto != 0){
		// add to last node .next
		newNode := ListNode{
			Val: resto,
			Next: nil,
		}

		aux.Next = &newNode;
		aux = &newNode
		
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