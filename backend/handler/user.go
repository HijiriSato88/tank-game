package handler

import (
	"net/http"

	"backend/model"
	"backend/pkg/jwtutil"
	"backend/repository"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// リクエストボディ
type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
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

	// 作成されたユーザー情報
	return c.JSON(http.StatusCreated, echo.Map{
		"message":   "user created",
		"username":  user.Username,
		"createdAt": user.CreatedAt,
	})
}

func Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	user, err := repository.GetUserByUsername(req.Username)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "user not found"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid password"})
	}

	token, err := jwtutil.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func Me(c echo.Context) error {
	claims, err := jwtutil.ExtractUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
	}

	u, err := repository.GetUserByID(claims.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "user not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"username":  u.Username,
	})
}
