package response

import system "Miyo_Assignment/model/user"

type RegisterResponse struct {
	User system.User `json:"user"`
}

type LoginResponse struct {
	User      system.User `json:"user"`
	Token     string      `json:"token"`
	ExpiresAt int64       `json:"expires_at"`
}
