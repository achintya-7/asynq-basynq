package models

type WelcomeApiModel struct {
	Email   string `form:"email" binding:"required"`
}
