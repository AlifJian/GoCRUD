package controller

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AlifJian/GoCRUD/config"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func AddUser(user User) {
	db := config.GetDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO user (name, email, password) VALUES (?,?,?)", user.Name, user.Email, user.Password)

	if err != nil {
		fmt.Println("ERROR : ", err)
		return
	}

	fmt.Println("SUCCESS")
}

func GetUser() *sql.Rows {
	db := config.GetDB()

	ctx := context.Background()
	row, err := db.QueryContext(ctx, "SELECT * FROM user")
	if err != nil {
		fmt.Println("ERROR : ", err)
	}

	return row
}

func UpdateUser(id int, user User) {
	db := config.GetDB()
	defer db.Close()

	ctx := context.Background()
	row, err := db.ExecContext(ctx, "UPDATE user SET name=?, email=?, password=? WHERE id=?", user.Name, user.Email, user.Password, id)

	if err != nil {
		fmt.Print("ERROR", err)
	}
	rowsAffected, _ := row.RowsAffected()
	if rowsAffected == 0 {
		fmt.Printf("No student found with Id: %d. Nothing was updated.\n", id)
		return
	}
	fmt.Println("Successfully updated student!")
}

func DeleteUser(id int) {
	db := config.GetDB()

	ctx := context.Background()
	_, err := db.ExecContext(ctx, "DELETE FROM user WHERE id=?", id)

	if err != nil {
		fmt.Print("EROR: ", err)
	}

	fmt.Print("SUCCESS DELETE")
}
