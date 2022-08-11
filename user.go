package main



type User struct {

	name string
	balance int
	owedTo map[string]int
	owedFrom map[string]int
}

func createNewUser (n string) User{
	user := User {
		name: n,
		balance: 0,
		owedTo: map[string]int{},
		owedFrom: map[string]int{},
	}
	return user
}

// func (u *User) updateBalance (b int, err) {
// 	u.balance += b
// }