package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	initialText := `
	To calculate how much you can spend with a discount, press 1.
	To calculate how much percentage a certain number is from an amount, press 2.
	To calculate how much a percentage is from a certain amount, press 3.
	For more detailed examples, press 4.
	Press ESC to quit
	`
	extendText := `
	Example 1: I have 100 USD and a 20% discount. How much can I spend? (Press 1)
	Example 2: I receive 100 USD, but my salary is 130 USD. How much tax did I pay? (Press 2)
	Example 3: They offer me a 20% raise on my 140 USD income. (Press 3)
	`
	fmt.Println(`
----------------------------------------------------------------------------------------------------------
------------------------------Welcome to Calculator discount and percentage-------------------------------
----------------------------------------------------------------------------------------------------------`)

	showInitialText := true
	checkNumber4 := false

	for {
		if showInitialText {
			fmt.Println(initialText)
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		switch char {
		case '1':
			fmt.Println("press 1")
		case '2':
			fmt.Println("press 2")
		case '3':
			fmt.Println("press 3")
		case '4':
			fmt.Println(extendText)
			showInitialText = false
			checkNumber4 = true
		default:
			if key == keyboard.KeyEsc {
				break
			}
			fmt.Printf("Press ESC to quit")
		}
		if key == keyboard.KeyEsc {
			break
		}

		if !checkNumber4 {
			showInitialText = true
		}
		checkNumber4 = false
	}

}
