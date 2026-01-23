package auth

import "time"

type LoginDTO struct { 
	Email string `json:"email"`
	Password string `json:"password"`
}


type ErrorDTO struct {
	Message string
	Time time.Time
}