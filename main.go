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
	prompt := promptui.Prompt{
		Label: "Enter the total amount you want to spend",
		Validate: func(input string) error {
			if _, err := fmt.Sscanf(input, "%f", new(float32)); err != nil {
				return fmt.Errorf("invalid total amount")
			}
			return nil
		},
	}
	amountStr, err := prompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var amount float32
	fmt.Sscanf(amountStr, "%f", &amount)

	prompt = promptui.Prompt{
		Label: "Enter the discount percentage (e.g., 20 for 20%)",
		Validate: func(input string) error {
			if _, err := fmt.Sscanf(input, "%d", new(int)); err != nil {
				return fmt.Errorf("invalid discount percentage")
			}
			return nil
		},
	}
	discountStr, err := prompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var discount int
	fmt.Sscanf(discountStr, "%d", &discount)

	canSpend := amount / (1 - float32(discount)/100)

	fmt.Printf("With a %d%% discount, you can spend: %.2f\n", discount, canSpend)
}

func percentageFromAmountCalculator() {

	amountInterface, err := promptHandler("Enter the total amount (e.g., your salary or total income)", "%f", "invalid total amount")
	amount, ok := amountInterface.(float64)
	if !ok {
		fmt.Println("Unexpected type for amount")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalInterface, err := promptHandler("Enter the amount you have (e.g., your actual earnings)", "%f", "invalid total amount")

	total, ok := totalInterface.(float64)
	if !ok {
		fmt.Println("Unexpected type for amount")
		return
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	percentage := -((amount / total) - 1) * -100

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

func promptHandler(labor string, format string, errorMessage string) (interface{}, error) {

	prompt := promptui.Prompt{
		Label: labor,
		Validate: func(input string) error {
			if format == "%d" {
				var tmp int
				if _, err := fmt.Sscanf(input, format, &tmp); err != nil {
					return fmt.Errorf(errorMessage, err)
				}
			} else if format == "%f" {
				var tmp float64
				if _, err := fmt.Sscanf(input, format, &tmp); err != nil {
					return fmt.Errorf(errorMessage, err)
				}
			} else {
				return fmt.Errorf("unsupported format")
			}
			return nil
		},
	}

	inputStr, err := prompt.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run prompt: %v", err)
	}

	if format == "%d" {
		var input int
		if _, err := fmt.Sscanf(inputStr, format, &input); err != nil {
			return nil, fmt.Errorf("invalid integer format: %v", err)
		}
		return input, nil
	} else if format == "%f" {
		var input float64
		if _, err := fmt.Sscanf(inputStr, format, &input); err != nil {
			return nil, fmt.Errorf("invalid float format: %v", err)
		}
		return input, nil
	}

	return nil, fmt.Errorf("unsupported format")
}
