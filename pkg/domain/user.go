package domain

type User struct {
	Id   uint   `json:"id" gorm:"unique;not null"`
	Name string `json:"name"`
	Mail string `json:"mail"`
}
