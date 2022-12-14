package models

import "time"

type User struct {
	ID        int       `json:"id"`
	FullName  string    `json:"fullName" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"password" gorm:"type: varchar(255)"`
	Role      string    `json:"role" gorm:"type: varchar(100)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersProfileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserTransactionResponse struct {
	ID        int    `json:"id"`
	FullName  string `json:"fullName" gorm:"type: varchar(255)"`
	Email     string `json:"email" gorm:"type: varchar(255)"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}

func (UserTransactionResponse) TableName() string {
	return "users"
}
