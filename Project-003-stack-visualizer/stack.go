package main

/*
   .-----------------------.
   |C>_ Author: Preetham G |
   |                       |
 __|_______________________|__
|  __________________________--|
`-/.::::::::::::::::::::::::.\-'a
 `----------------------------'

*/

import (
	"fmt"
)

type stack struct {
	arr []int
	top int
}

var newStack stack

const (
	ceilingBrick = "╔═════╗"

	midBrick = "╠═════╣"

	bottomBrick = "╚═════╝"

	topSymbol = " ⟸   top "
)

/*

*************************************
* ██████╗ ██╗   ██╗███████╗██╗  ██╗ *
* ██╔══██╗██║   ██║██╔════╝██║  ██║ *
* ██████╔╝██║   ██║███████╗███████║ *
* ██╔═══╝ ██║   ██║╚════██║██╔══██║ *
* ██║     ╚██████╔╝███████║██║  ██║ *
* ╚═╝      ╚═════╝ ╚══════╝╚═╝  ╚═╝ *
*************************************

 */

func pushToStack() {

	if !isStackFull() {
		fmt.Printf("Enter nummber to Push\n")
		var pushInt int
		fmt.Scanln(&pushInt)

		if pushInt > 9999 {
			fmt.Printf("%v\n", "Max 5 digit number is allowed")
			return
		}

		newStack.top += 1
		newStack.arr = append(newStack.arr, pushInt)

		printStack()
	} else {
		fmt.Printf("%v", "Stack is full !!!\n")
	}

}

/*

*************************************
* ██████╗  ██████╗ ██████╗          *
* ██╔══██╗██╔═══██╗██╔══██╗         *
* ██████╔╝██║   ██║██████╔╝         *
* ██╔═══╝ ██║   ██║██╔═══╝          *
* ██║     ╚██████╔╝██║              *
* ╚═╝      ╚═════╝ ╚═╝              *
*************************************

 */

func popFromStack() {

	if !isStackEmpty() {
		fmt.Printf("Popping the top\n")

		newStack.arr = newStack.arr[:len(newStack.arr)-1]
		newStack.top -= 1

		printStack()
	} else {
		fmt.Printf("%v\n", "Stack is empty !!!")
	}
}

func isStackEmpty() bool {
	return (newStack.top == -1)
}

func isStackFull() bool {
	return (newStack.top == 9)
}

/*

*******************************************
* ██████╗ ██████╗ ██╗███╗   ██╗████████╗  *
* ██╔══██╗██╔══██╗██║████╗  ██║╚══██╔══╝  *
* ██████╔╝██████╔╝██║██╔██╗ ██║   ██║     *
* ██╔═══╝ ██╔══██╗██║██║╚██╗██║   ██║     *
* ██║     ██║  ██║██║██║ ╚████║   ██║     *
* ╚═╝     ╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝   ╚═╝     *
*******************************************

 */

func printStack() {

	fmt.Printf("%v\n", ceilingBrick)

	if !isStackEmpty() {
		for i := newStack.top; i >= 0; i-- {

			if i == newStack.top {
				fmt.Printf("║%5d║%v\n", newStack.arr[i], topSymbol)
			} else {
				fmt.Printf("║%5d║\n", newStack.arr[i])
			}

			if i != 0 {
				fmt.Printf("%v\n", midBrick)
			}
		}
	}

	fmt.Printf("%v\n", bottomBrick)
}

func main() {

	newStack.top = -1

REPL:
	for {

		fmt.Printf("Enter the operation number: 1.Push \t 2.Pop \t 3.Print Stack \t  4.Exit \n")

		var inputOperation int
		fmt.Scanln(&inputOperation)

		switch inputOperation {
		case 1:
			pushToStack()
		case 2:
			popFromStack()
		case 3:
			printStack()
		default:
			break REPL
		}
	}

}
