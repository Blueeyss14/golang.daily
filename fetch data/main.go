package main

import "fmt"

func main() {
	users := fetchData()
	if users == nil {
		fmt.Println("Gak ada data")
		return
	}

	for i := 0; i < len(users); i++ {
		fmt.Printf("Name: %s %s\nEmail: %s\n", users[i].FirstName, users[i].LastName, users[i].Email)
	}
}
