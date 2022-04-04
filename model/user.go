package model

import (
	"time"
)

type User struct {
	Id        int       `json:"id" grom:"primarikey"`
	UserName  string    `json:"userName" grom:"user_name"`
	Password  string    `json:"password" grom:"password"`
	Phone     string    `json:"phone" grom:"phone"`
	Role      string    `json:"role" grom:"role"`
	CreatedAt time.Time `json:"createAt" gorm:"create_at"`
	UpdatedAt time.Time `json:"updatedAt" grom:"updated_at"`
	Delete    *time.Time `json:"deleteAt" grom:"delete_at"`
}

type UserRepository interface {
	GetById(id int) (*User, error)
	GetAll() ([]*User, error)
	Create(newRecord *User) (*User, error)
	Update(id int, newRecord *User) (*User, error)
	Delete(id int) error
}
