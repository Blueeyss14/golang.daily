package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func fetchData() []User {
	url := "https://reqres.in/api/users?page=2"
	var response, err = http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer response.Body.Close() //biar gak resource leak

	var result struct {
		Data []User `json:"data"`
	}

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	return result.Data
}

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
