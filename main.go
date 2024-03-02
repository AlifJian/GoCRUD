package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AlifJian/GoCRUD/config"
	"github.com/AlifJian/GoCRUD/controller"
)

func main() {
	db := config.GetDB()
	if db == nil {
		fmt.Print("ERROR CONECTION")
	}
	addUser()
	getUser()
	updateUser()
	deleteUser()

}

func makeUser() controller.User {
	var user controller.User

	input := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Name: ")
	input.Scan()
	user.Name = input.Text()

	fmt.Print("Enter Email: ")
	input.Scan()
	user.Email = input.Text()

	fmt.Print("Enter Password: ")
	input.Scan()
	user.Password = input.Text()

	return user
}

func addUser() {
	controller.AddUser(makeUser())
}

func updateUser() {
	var id int
	fmt.Print("Masukkan Id dari user: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("Invalid input")
	}
	fmt.Scanln()
	controller.UpdateUser(id, makeUser())
}

func deleteUser() {
	var id int
	fmt.Print("Masukkan Id dari user: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("Invalid input")
	}
	fmt.Scanln()
	controller.DeleteUser(id)
}

func getUser() {
	users := controller.GetUser()

	for users.Next() {
		var user controller.User
		err := users.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Name:%s\nEmail:%s\n", user.Name, user.Email)
	}
	if err := users.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
