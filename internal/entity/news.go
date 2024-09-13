package entity

type News struct {
	ID         int    `json:"id,omitempty" db:"id"`
	UserID     string `json:"user_id,omitempty" db:"user_id"`
	Title      string `json:"title,omitempty" db:"title"`
	Content    string `json:"content,omitempty" db:"content"`
	Categories []int  `json:"categories,omitempty" db:"-"`
}
