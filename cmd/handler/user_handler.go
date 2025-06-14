package handler

import (
	"net/http"

	"g42-user/cmd/logic"
	"g42-user/repositories"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userLogic *logic.UserLogic
}

func NewUserHandler(userLogic *logic.UserLogic) *UserHandler {
	return &UserHandler{userLogic: userLogic}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignupRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	Mobile      string `json:"mobile,omitempty"`
	DateOfBirth string `json:"dateOfBirth,omitempty"`
}

type LoginResponse struct {
	Token string            `json:"token"`
	User  repositories.User `json:"user"`
}

type SignupResponse struct {
	Message string            `json:"message"`
	User    repositories.User `json:"user"`
}

type GetUserDetailsRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	token, user, err := h.userLogic.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Token: token,
		User:  user,
	})
}

func (h *UserHandler) Signup(c *gin.Context) {
	var req SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Check if user already exists
	existingUser, _ := h.userLogic.GetUserByEmail(req.Email)
	if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
		return
	}

	// Create new user
	user := &repositories.User{
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		Mobile:      req.Mobile,
		DateOfBirth: req.DateOfBirth,
	}

	if err := h.userLogic.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, SignupResponse{
		Message: "User created successfully",
		User:    *user,
	})
}

func (h *UserHandler) Logout(c *gin.Context) {
	// Since we're using JWT, we don't need to do anything server-side
	// The client should remove the token
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (h *UserHandler) GetUserDetails(c *gin.Context) {
	var req GetUserDetailsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Get the authenticated user's email from the context
	authenticatedEmail, exists := c.Get("userEmail")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Check if the authenticated user is requesting their own details
	if authenticatedEmail != req.Email {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only view your own details"})
		return
	}

	user, err := h.userLogic.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (h *UserHandler) GetUserDetailsByID(c *gin.Context) {
	userID := c.Query("userId")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}

	// Get the authenticated user's email from the context
	authenticatedEmail, exists := c.Get("userEmail")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := h.userLogic.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Check if the authenticated user is requesting their own details
	if authenticatedEmail != user.Email {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only view your own details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
