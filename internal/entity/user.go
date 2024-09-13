package entity

type UserRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password," db:"password_hash"`
}

type User struct {
	ID           string `json:"id,omitempty" db:"id"`
	Username     string `json:"username" db:"username"`
	Email        string `json:"email" db:"email"`
	PasswordHash []byte `json:"password," db:"password_hash"`
}
