package db

import (
	"github.com/vdjurovic/go-webapp-sample/model"
)

func SaveUser(user model.User) int64 {
	stmt, err := db.Prepare("INSERT INTO users(firstname,lastname,email,pswd) VALUES($1,$2,$3,$4)")
	if err != nil {
		panic(err)
	}
	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()
	return id
}
