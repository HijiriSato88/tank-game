package handler

import (
	"backend/pkg/jwtutil"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: u}
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *UserHandler) Signup(c echo.Context) error {
	var req SignupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	user, err := h.userUsecase.Signup(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create user"})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"message":   "user created",
		"username":  user.Username,
		"createdAt": user.CreatedAt,
	})
}

func (h *UserHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	user, err := h.userUsecase.Login(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid username or password"})
	}

	token, err := jwtutil.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func (h *UserHandler) Me(c echo.Context) error {
	claims, err := jwtutil.ExtractUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
	}
	user, err := h.userUsecase.GetUser(claims.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "user not found"})
	}
	return c.JSON(http.StatusOK, echo.Map{"username": user.Username})
}
