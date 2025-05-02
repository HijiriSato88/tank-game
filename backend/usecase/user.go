package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Signup(username, password string) (*model.User, error)
	Login(username, password string) (*model.User, error)
	GetUser(userID int) (*model.User, error)
	UpdateHighScore(userID int, newScore int) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: r}
}

func (u *userUsecase) Signup(username, password string) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
	}
	err = u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Login(username, password string) (*model.User, error) {
	user, err := u.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) GetUser(userID int) (*model.User, error) {
	return u.userRepo.GetByID(userID)
}

func (u *userUsecase) UpdateHighScore(userID int, newScore int) error {
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		return err
	}
	
	if newScore > user.HighScore {
		return u.userRepo.UpdateHighScore(userID, newScore)
	}
	return nil
}
