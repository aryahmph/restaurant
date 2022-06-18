package http

import (
	"time"
)

type (
	Error struct {
		Code    int               `json:"code"`
		Message string            `json:"message"`
		Errors  map[string]string `json:"errors"`
	}

	Login struct {
		ID       string `json:"id"`
		Username string `json:"username" binding:"required,gte=5,lte=25"`
		Password string `json:"password" binding:"required,gte=8,lte=16"`
	}

	Delete struct {
		ID string `uri:"id" binding:"uuid4"`
	}

	User struct {
		ID           string    `json:"id"`
		Username     string    `json:"username" binding:"required,gte=5,lte=25"`
		PasswordHash string    `json:"password_hash" binding:"required,gte=8,lte=16"`
		Role         string    `json:"role" binding:"required,oneof=WAITER CASHIER ADMIN"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	AddUser struct {
		ID       string `json:"id"`
		Username string `json:"username" binding:"required,gte=5,lte=25"`
		Password string `json:"password" binding:"required,gte=8,lte=16"`
		Role     string `json:"role" binding:"required,oneof=WAITER CASHIER ADMIN"`
	}

	UpdateUser struct {
		ID       string `uri:"id" binding:"uuid4"`
		Username string `json:"username" binding:"required,gte=5,lte=25"`
		Role     string `json:"role" binding:"required,oneof=WAITER CASHIER ADMIN"`
	}

	Category struct {
		ID        string    `json:"id"`
		Name      string    `json:"name" binding:"required,lte=25"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	UpdateCategory struct {
		ID   string `uri:"id" binding:"uuid4"`
		Name string `json:"name" binding:"required,lte=25"`
	}
)
