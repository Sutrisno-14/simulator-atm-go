package main

import (
	"fmt"
	"os"
	"os/exec"
)


func callClear() {
    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
}
var user1 string
var user2 string
var balance float32 = 0
var transfer float32 = 0
var anotherTransaction int

func login() {
	fmt.Printf("Login ")
	fmt.Scan(&user1)
	fmt.Println()
	fmt.Printf("hallo, %s !\n\n",user1)
	fmt.Printf("Your balance is $%.0f\n\n", balance)
}
func transaction() {

	var choise  int
	fmt.Printf("Choisee any option \n\n")
	fmt.Println("1. Deposite")
	fmt.Println("2. Transafer")
	fmt.Println("3. Cek account")
	fmt.Printf("4. Logout\n\n")
	fmt.Scan(&choise)

	var amountToDeposit float32
	var amountToWithdraw float32

	switch choise {
		case 1 :
			//Deposite
			fmt.Printf("\nEnter Deposit : ")
			fmt.Scan(&amountToDeposit)
			balance += amountToDeposit
			fmt.Printf("Your balance is $%.0f", balance)
			fmt.Printf("\n\nDo you want another transaction?\nPress 1 to next transaction and 2 to Logout : ")
			fmt.Scan(&anotherTransaction)

			switch anotherTransaction {
			case 1:
				callClear()
				transaction()
			default:
				fmt.Printf("\nGoodbye, %s!\n\n", user1)
		    }

		// //transfer
		case 2:
			fmt.Printf("Transfer to ")
			fmt.Scan(&user2,&transfer)

			if transfer > balance {
				fmt.Printf("\nThere is no insufficient funds in your account")
				fmt.Printf("\n\nDo you want another transaction?\nPress 1 to next transaction and 2 to Logout : ")
			    fmt.Scan(&anotherTransaction)
				switch anotherTransaction {
				case 1:
					callClear()
					transaction()
				default:
					fmt.Printf("\nGoodbye, %s!\n\n", user1)
				}
			}else {
				fmt.Printf("Transferred $%.0f to %s\n\n", transfer, user2)
				balance -= transfer
				fmt.Printf("Your balance is $%.0f", balance)
				fmt.Printf("\n\nDo you want another transaction?\nPress 1 to next transaction and 2 to Logout : ")
			    fmt.Scan(&anotherTransaction)
				switch anotherTransaction {
				case 1:
					callClear()
					transaction()
				default:
					fmt.Printf("\nGoodbye, %s!\n\n", user1)
				}
			}
		//check account
		case 3:
			fmt.Printf("\nhallo, %s !\n\n",user1)
			fmt.Printf("Your balance is $%.0f\n\n", balance)
			amountToWithdraw += transfer
			fmt.Printf("Owned %.0f to %s", amountToWithdraw, user2)
			fmt.Printf("\n\nDo you want another transaction?\nPress 1 to next transaction and 2 to Logout : ")
			    fmt.Scan(&anotherTransaction)
				switch anotherTransaction {
				case 1:
					callClear()
					transaction()
				default:
					fmt.Printf("\nGoodbye, %s!\n\n", user1)
				}
		case 4:
			fmt.Printf("\nGoodbye, %s\n\n", user1)
	}
}

func main() {
	fmt.Printf("Welcome in our ATM\n\n")
	login()
	transaction()
}