package handler

import (
	"net/http"
	"backend/model"
	"backend/repository"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// リクエストボディ
type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signup(c echo.Context) error {
	var req SignupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "username and password required"})
	}

	// bcryptでパスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to hash password"})
	}

	user := model.User{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
	}

	err = repository.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create user"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "user created"})
}
