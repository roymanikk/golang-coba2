package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//slice of users
var listUser []User
var currentUser User

func getInput (prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	input = strings.TrimSpace(input)
	return input, err
}

func prompt() {
	
	reader := bufio.NewReader(os.Stdin)
	input, _ := getInput("\n", reader)
	command := strings.Split(input, " ")[0]

	switch command {
	case "login":
		name := strings.Split(input, " ")[1]
		if !contains(listUser, name) {
			newUser := createNewUser(name)
			currentUser = newUser
			listUser = append(listUser, newUser)
			fmt.Printf("Hello %v! Your balance is $%v\n", currentUser.name, currentUser.balance)
		} else {
			for i := range listUser {
				if listUser[i].name == name {
						currentUser = listUser[i]
						fmt.Printf("Hello %v! Your balance is $%v\n", currentUser.name, currentUser.balance)
				}
			}
		}
		prompt()
		
	case "deposit":
		deposit := strings.Split(input, " ")[1]
		d, _ := strconv.Atoi(deposit)
		currentUser.balance = currentUser.balance + d

		//update balance current user
		for i := range listUser {
			if listUser[i].name == currentUser.name {
				listUser[i].balance = currentUser.balance
			}
		}
		fmt.Println("current user :", currentUser)
		fmt.Println("Your balance is: $", currentUser.balance)
		prompt()

	case "withdraw":
		withdraw := strings.Split(input, " ")[1]
		w, _ := strconv.Atoi(withdraw)

		// cek apakah saldo cukup
		if currentUser.balance < w {
			fmt.Println("Saldo tidak cukup")
		} else {
			//update balance current user
			currentUser.balance = currentUser.balance - w
			for i := range listUser {
				if listUser[i].name == currentUser.name {
					listUser[i].balance = currentUser.balance
				}
			}
			fmt.Printf("You withdraw an amount of $%v Your balance now is: $%v", w ,currentUser.balance)
		}

		prompt()

	case "transfer":
		recipient := strings.Split(input, " ")[1]
		amount := strings.Split(input, " ")[2]
		a, _ := strconv.Atoi(amount)
		
		// cek apakah saldo cukup
		if currentUser.balance < a {
			fmt.Println("ngutang")
			for i := range listUser {
				//update balance current user
				if listUser[i].name == currentUser.name {
					listUser[i].balance = currentUser.balance - a
					listUser[i].owedTo = map[string]int{recipient: currentUser.balance - a}
					fmt.Printf("You transferred $%v to %v your balance now is $0. You owed %v to %v",  currentUser.balance, recipient, currentUser.owedTo[recipient], recipient)
					
					//tambah saldo recipient
					for i := range listUser {
						if listUser[i].name == recipient {
							listUser[i].balance += currentUser.balance
							listUser[i].owedFrom = map[string]int{currentUser.name: a - currentUser.balance}
							
						}
					}		
					
				}
			}
			
		} else {
			//update balance current user
			currentUser.balance = currentUser.balance - a
			for i := range listUser {
				if listUser[i].name == currentUser.name {
					listUser[i].balance = currentUser.balance
				}
			}

			//tambah saldo recipient
			for i := range listUser {
				if listUser[i].name == recipient {
					listUser[i].balance += a
				}
			}

			//update current user
			for i := range listUser {
				if listUser[i].name == currentUser.name {
					listUser[i].balance = currentUser.balance
				}
			}

			fmt.Printf("You transferred $%v to %v your balance now is %v",  a, recipient, currentUser.balance)
		}
		prompt()
	case "logout":
		fmt.Printf("Goodbye %v!", currentUser.name)
		currentUser.name = ""
		fmt.Println("current user :", currentUser.name, "current balance :", currentUser.balance)
		prompt()
		default:
			fmt.Println("Invalid command")
			prompt()
			
	}
}

func main() {
	prompt()
}