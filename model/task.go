package model

import "time"

type Task struct{
	ID 		  uint 		`json:"id" gorm:"primary_key"`
	Title 	  string 	`json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User 	  User 		`json:"user" gorm:"foreignKey:UserId; constrtaiant:OnDelete:CASCADE"`//紐づいているユーザが削除されれば、そのユーザのタスクも削除される
	UserId 	  uint 		`json:"user_id" gorm:"not null"`
}

type TaskResponse struct{
	ID 		  uint 		`json:"id" gorm:"primary_key"`
	Title 	  string 	`json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}