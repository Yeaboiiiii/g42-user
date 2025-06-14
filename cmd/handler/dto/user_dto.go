package dto

type SignupRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	Mobile      string `json:"mobile,omitempty"`
	DateOfBirth string `json:"dateOfBirth,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type GetUserDetailsRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type SignupResponse struct {
	Message string       `json:"message"`
	User    UserResponse `json:"user"`
}

type UserResponse struct {
	ID          string `json:"userId"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile,omitempty"`
	DateOfBirth string `json:"dateOfBirth,omitempty"`
}
