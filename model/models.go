package model

type User struct {
	// only fields with capital first letter are visible to external packages
	UserID    int64  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
