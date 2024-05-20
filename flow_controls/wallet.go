package flowcontrols

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Menu() {
	var opt int
	var balance float64

	for opt != 4 {
		fmt.Println(`
		***** Welcome to wallet *****
		***** [1] Deposit       *****
		***** [2] Withdrawal    *****
		***** [3] Balance       *****
		***** [4] exit          *****
		`)

		fmt.Print("choose opt: ")
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			deposit(&balance)
		case 2:
			e := withdrawal(&balance)
			if e != nil {
				fmt.Println(e)
			}
		case 3:
			fmt.Println("Balance:", balance)
		case 4:
			os.Exit(1)
		default:
			fmt.Println("Invalid opt")
		}

		time.Sleep(time.Second * 1)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()

		if err != nil {
			fmt.Println("error clearing fails")
		}
	}
}

func deposit(balance *float64) {
	var amount float64
	fmt.Print("Write amount to deposit:")
	fmt.Scanln(&amount)

	*balance += amount
}

func withdrawal(balance *float64) error {
	var amount float64

	fmt.Print("Write amount to withdrawal: ")
	fmt.Scanln(&amount)

	if amount > *balance {
		return fmt.Errorf("error: amount {%v} > balance {%v}", amount, *balance)
	}

	*balance -= amount

	return nil
}
