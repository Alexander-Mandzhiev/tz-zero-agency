package entity

type NewsCategory struct {
	NewsId     int64 `json:"news_id" db:"news_id"`
	CategoryId int64 `json:"category_id" db:"category_id"`
}
