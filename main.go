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
			fmt.Printf("Hello %v!\n Your balance is $%v\n", currentUser.name, currentUser.balance)
		} else {
			for i := range listUser {
				if listUser[i].name == name {
					currentUser = listUser[i]
					owedFrom:= listUser[i].owedFrom
					owedTo:= listUser[i].owedTo
					fmt.Printf("Hello %v!\n Your balance is $%v\n", currentUser.name, currentUser.balance)
					
					// check if user has any debts
					if len(owedFrom) > 0 {
					for k, v := range owedFrom {
						fmt.Printf("owed %v from $%v\n", v, k)
						}
					}
					if len(owedTo) > 0 {
					for k, v := range owedTo {
						fmt.Printf("owed %v to $%v\n", v, k)
						}
					}	
				}
			}
		}
		prompt()
		
	case "deposit":
		deposit := strings.Split(input, " ")[1]
		d, _ := strconv.Atoi(deposit)
		owedTo := currentUser.owedTo

		//deduct from owedTo
		if len(owedTo) > 0 {
			for k, v := range owedTo {
				for i := range listUser {
					if listUser[i].name == k {
						listUser[i].balance = listUser[i].balance + d
						if (v + d) > 0 {
							fmt.Printf("Transferred $%v to %v\n", d - (v + d), k)
						} else {
							fmt.Printf("Transferred $%v to %v\n", d, k)
						}
						
						// fmt.Printf("d: $%v v: %v\n", d, v)
						listUser[i].owedFrom[currentUser.name] = v + d
						currentUser.owedTo[k] = v + d
						if (currentUser.owedTo[k] > 0) {
							delete(currentUser.owedTo, k)
							// fmt.Println("test", currentUser.owedTo)
							currentUser.balance = v + d
							// fmt.Printf("test", len(owedTo))
						}
						fmt.Printf("Your balance is $%v\n", currentUser.balance)
						if len(currentUser.owedTo) > 0 {
						fmt.Printf("Owed $%v to %v\n", currentUser.owedTo[k], k)
					}

						for i:= range listUser {
							if listUser[i].name == currentUser.name {
								listUser[i].owedTo = currentUser.owedTo
								listUser[i].balance = currentUser.balance
							}
						}
					}
				}	
			}
		} else {
		//update balance current user
		for i := range listUser {
			if listUser[i].name == currentUser.name {
				listUser[i].balance += d
				currentUser.balance = listUser[i].balance
			}
		}
		fmt.Println("Your balance is: $", currentUser.balance)
	}
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
		if currentUser.balance < a  {
			fmt.Println("ngutang")
			for i := range listUser {
				//update balance current user
				if listUser[i].name == currentUser.name {
					listUser[i].owedTo = map[string]int{recipient: currentUser.balance - a}
					currentUser.owedTo = map[string]int{recipient: currentUser.balance - a}
					fmt.Printf("You transferred $%v to %v, your balance now is 0. You owed %v to %v", currentUser.balance , recipient, listUser[i].owedTo[recipient], recipient)
					
					//tambah saldo recipient
					for i := range listUser {
						if listUser[i].name == recipient {
							listUser[i].balance += currentUser.balance
							listUser[i].owedFrom = map[string]int{currentUser.name: a - currentUser.balance}
							
						}
					}

					listUser[i].balance = 0
					currentUser.balance = 0
					
				}
			}
		// check if the recipient is in the list of user owedFrom
		} else if len(currentUser.owedFrom) > 0 {
			for i := range listUser {
				if listUser[i].name == recipient {
					// fmt.Println("test", listUser[i].owedTo[currentUser.name])
					listUser[i].owedTo[currentUser.name] += a
					// fmt.Println("test", listUser[i].owedTo[currentUser.name])
					currentUser.owedFrom[recipient] += a
					fmt.Printf("Your balance is %v\n", currentUser.balance)
					fmt.Printf("Owed %v from $%v\n", currentUser.owedFrom[recipient], recipient)
					
					//update balance current user in listUser
					for i := range listUser {
						if listUser[i].name == currentUser.name {
							listUser[i].owedFrom[recipient] = currentUser.owedFrom[recipient]
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
		// fmt.Println("current user :", currentUser.name, "current balance :", currentUser.balance)
		prompt()
		default:
			fmt.Println("Invalid command")
			prompt()
			
	}
}

func main() {
	prompt()
}