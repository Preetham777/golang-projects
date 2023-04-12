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
	"strings"
)

type queue struct {
	arr   []int
	start int
}

var newQ queue

const (
	ceilingBrick = "╔═════╗"

	bottomBrick = "╚═════╝"

	ceilingBrickWithEndSymbol = "   ⇓\n╔═════╗"

	bottomBrickWithStartSymbol = "╚═════╝ ⇖ FI"
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

func pushToQueue() {

	if !isQueueFull() {
		fmt.Printf("Enter nummber to Push\n")
		var pushInt int
		fmt.Scanln(&pushInt)

		if pushInt > 9999 {
			fmt.Printf("%v\n", "Max 5 digit number is allowed")
			return
		}

		newQ.start += 1

		newQ.arr = append(newQ.arr, pushInt)

		printQueue()
	} else {
		fmt.Printf("%v", "Queue is full !!!\n")
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

func popFromQueue() {

	if !isQueueEmpty() {
		fmt.Printf("Popping the Top\n")

		newQ.arr = newQ.arr[1:]
		newQ.start -= 1

		if !isQueueEmpty() {
			printQueue()
		}

	} else {
		fmt.Printf("%v\n", "Queue is empty !!!")
	}
}

func isQueueEmpty() bool {
	return (len(newQ.arr) == 0)
}

func isQueueFull() bool {
	return (len(newQ.arr) == 10)
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

func printQueue() {

	ceilingBricks := ""
	queueUnicode := ""
	bottomBricks := ""

	if !isQueueEmpty() {

		ceilingBricks += ceilingBrickWithEndSymbol
		ceilingBricks += strings.Repeat(ceilingBrick, len(newQ.arr)-1)

		for i := newQ.start; i >= 0; i-- {

			if i == 0 {
				queueUnicode += fmt.Sprintf("╣%5d║", newQ.arr[i])
			} else if i == newQ.start {
				queueUnicode += fmt.Sprintf("║%5d╠", newQ.arr[i])
			} else {
				queueUnicode += fmt.Sprintf("╣%5d╠", newQ.arr[i])
			}

		}

		bottomBricks += strings.Repeat(bottomBrick, len(newQ.arr)-1)
		bottomBricks += bottomBrickWithStartSymbol

		fmt.Printf("\n\n%v\n%v\n%v\n\n\n", ceilingBricks, queueUnicode, bottomBricks)

	} else {
		fmt.Printf("%v\n", "Queue is empty !!!")
	}

}

func main() {

	// Initialize the start to -1
	newQ.start = -1

REPL:
	for {

		fmt.Printf("Enter the operation number: 1.Push \t 2.Pop \t 3.Print Stack \t  4.Exit \n")

		var inpustarteration int
		fmt.Scanln(&inpustarteration)

		switch inpustarteration {
		case 1:
			pushToQueue()
		case 2:
			popFromQueue()
		case 3:
			printQueue()
		default:
			break REPL
		}
	}

}
