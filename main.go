package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/manifoldco/promptui"
)

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println(`
----------------------------------------------------------------------------------------------------------
------------------------------Welcome to Calculator discount and percentage-------------------------------
----------------------------------------------------------------------------------------------------------`)
	for {

		prompt := promptui.Select{
			Label: "Choose an option",
			Items: []string{
				"1 - Calculate how much you can spend with a discount",
				"2 - Calculate how much percentage a certain number is from an amount",
				"3 - Calculate how much a percentage is from a certain amount",
				"4 - Show detailed examples",
				"Quit",
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch result {
		case "1 - Calculate how much you can spend with a discount":
			discountCalculator()
		case "2 - Calculate how much percentage a certain number is from an amount":
			percentageFromAmountCalculator()
		case "3 - Calculate how much a percentage is from a certain amount":
			percentageOfAmountCalculator()
		case "4 - Show detailed examples":
			fmt.Println(`
	Example 1: I have 100 USD and a 20% discount. How much can I spend? 
	Example 2: I receive 100 USD, but my salary is 130 USD. How much tax did I pay? 
	Example 3: They offer me a 20% raise on my 140 USD income. 
	Press ESC to quit.
	`)
		default:
			fmt.Println("Are you sure?")
		}
		fmt.Println("Press ESC to quit.")
		fmt.Printf("Or any other key to restart.")

		if _, key, err := keyboard.GetKey(); err != nil {
			panic(err)
		} else if key == keyboard.KeyEsc {
			break
		}
	}

}

func discountCalculator() {
	var amount float32
	var discount int

	fmt.Print("Enter the total amount you want to spend: ")
	fmt.Scanf("%f", &amount)
	fmt.Println(amount)

	fmt.Print("Enter the discount percentage (e.g., 20 for 20%): ")
	fmt.Scanf("%d", &discount)
	fmt.Println(discount)

	canSpend := amount / (1 - float32(discount)/100)

	fmt.Printf("With a %d%% discount, you can spend: %.2f\n", discount, canSpend)
}

func percentageFromAmountCalculator() {
	var amount, total float32

	fmt.Print("Enter the total amount (e.g., your salary or total income): ")
	fmt.Scanf("%f", &total)
	fmt.Println(total)

	fmt.Print("Enter the amount you have (e.g., your actual earnings): ")
	fmt.Scanf("%f", &amount)
	fmt.Println(amount)

	percentage := ((amount / total) - 1) * -100

	fmt.Printf("The amount %.2f is %.2f%% less than the total amount %.2f\n", amount, percentage, total)
}

func percentageOfAmountCalculator() {
	prompt := promptui.Prompt{
		Label: "Enter the total amount",
		Validate: func(input string) error {
			if _, err := fmt.Sscanf(input, "%f", new(float32)); err != nil {
				return fmt.Errorf("invalid total amount")
			}
			return nil
		},
	}
	totalStr, err := prompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var total float32
	fmt.Sscanf(totalStr, "%f", &total)

	prompt = promptui.Prompt{
		Label: "Enter the percentage to calculate (e.g., 20 for 20%)",
		Validate: func(input string) error {
			if _, err := fmt.Sscanf(input, "%f", new(float32)); err != nil {
				return fmt.Errorf("invalid percentage")
			}
			return nil
		},
	}
	percentageStr, err := prompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var percentage float32
	fmt.Sscanf(percentageStr, "%f", &percentage)

	result := total * (1 + percentage/100)
	fmt.Printf("Applying a %.2f%% increase to %.2f results in: %.2f\n", percentage, total, result)
}
