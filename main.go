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

	flags := struct {
		check           bool
		escPressed      bool
		fourPressed     bool
		showInitialText bool
	}{
		escPressed:      false,
		fourPressed:     false,
		check:           false,
		showInitialText: true,
	}

	hideInitialMessage := func() {
		flags.showInitialText = false
		flags.check = true
	}

	for {
		if flags.showInitialText {
			fmt.Println(initialText)
			if flags.escPressed {
				flags.escPressed = false
			}
			if flags.fourPressed {
				flags.fourPressed = false
			}
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		switch char {
		case '1':
			discountCalculator()
		case '2':
			fmt.Println("press 2")
		case '3':
			fmt.Println("press 3")
		case '4':
			hideInitialMessage()
			if !flags.fourPressed {
				fmt.Println(extendText)
			}
			flags.fourPressed = true
		default:
			if key == keyboard.KeyEsc {
				break
			}
			hideInitialMessage()
			if !flags.escPressed {
				fmt.Printf("Press ESC to quit")
			}
			flags.escPressed = true

		}

		if key == keyboard.KeyEsc {
			break
		}

		if !flags.check {
			flags.showInitialText = true
		}
		flags.check = false
	}

}

func discountCalculator() {
	var amount float32
	var discount int

	fmt.Print("Enter the amount you have to spend: ")
	fmt.Scanf("%f", &amount)
	fmt.Println(amount)

	fmt.Print("Enter the discount percentage: ")
	fmt.Scanf("%d", &discount)
	fmt.Println(discount)

	canSpend := amount / (1 - (float32(discount) / 100))
	fmt.Print("You can spend: ")
	fmt.Println(canSpend)
}
