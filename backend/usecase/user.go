package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"fmt"

	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Signup(username, password string) (*model.User, error)
	Login(username, password string) (*model.User, error)
	GetUser(userID int) (*model.User, error)
	UpdateHighScore(userID int, eventSlug string, newScore int) error
}

type userUsecase struct {
	userRepo    repository.UserRepository
	rankingRepo repository.RankingRepository
}

func NewUserUsecase(u repository.UserRepository, r repository.RankingRepository) UserUsecase {
	return &userUsecase{
		userRepo:    u,
		rankingRepo: r,
	}
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

func (u *userUsecase) UpdateHighScore(userID int, eventSlug string, newScore int) error {
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		return err
	}

	event, err := u.rankingRepo.GetEventBySlug(eventSlug)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 現在のハイスコア取得（user_scoresテーブルを使う）
	current, err := u.userRepo.GetUserScore(userID, event.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if current != nil && newScore <= current.HighScore {
		return nil
	}

	if current == nil {
		if err := u.userRepo.CreateUserScore(userID, event.ID, newScore); err != nil {
			return err
		}
	} else {
		if err := u.userRepo.UpdateUserScore(userID, event.ID, newScore); err != nil {
			return err
		}
	}

	adjustedScore := int64(newScore)*1e9 + (event.EndAt.Unix() - time.Now().Unix())
	redisKey := "ranking:" + event.Slug

	return u.rankingRepo.ZAddScore(redisKey, user.Username, adjustedScore)
}
