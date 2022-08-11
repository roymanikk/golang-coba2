package main


func contains(s []User, e string) bool {
	for _, a := range s {
		if a.name == e {
			return true
		}
	}
	return false
}