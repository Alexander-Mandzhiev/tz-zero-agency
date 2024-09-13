package entity

type Category struct {
	ID     int64  `json:"id,omitempty" db:"id"`
	UserID string `json:"user_id,omitempty" db:"user_id"`
	Title  string `json:"title,omitempty" db:"title"`
}
