package handlers

import (
	"database/sql"
	"net/http"

	"auth-login/database"
	"auth-login/middleware"
	"auth-login/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// POST /api/auth/login — handled by Go Backend only
func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "ข้อมูลไม่ถูกต้อง"})
		return
	}

	var user models.User
	err := database.DB.QueryRow(
		"SELECT id, username, password FROM users WHERE username = ?",
		req.Username,
	).Scan(&user.ID, &user.Username, &user.Password)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Username หรือ Password ไม่ถูกต้อง"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "เกิดข้อผิดพลาดภายในระบบ"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Username หรือ Password ไม่ถูกต้อง"})
		return
	}

	token, err := middleware.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "ไม่สามารถสร้าง token ได้"})
		return
	}

	c.JSON(http.StatusOK, models.LoginResponse{
		Token:    token,
		Username: user.Username,
	})
}
