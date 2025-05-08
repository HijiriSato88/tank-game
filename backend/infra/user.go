package infra

import (
	"backend/db"
	"backend/domain/model"
	"backend/domain/repository"
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *model.User) error {
	_, err := db.DB.Exec(`
		INSERT INTO users (username, password_hash)
		VALUES (?, ?)`,
		user.Username, user.PasswordHash,
	)
	return err
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := db.DB.Get(&user, "SELECT id, username, password_hash FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByID(id int) (*model.User, error) {
	var user model.User
	err := db.DB.Get(&user, "SELECT id, username FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserScore(userID, eventID int) (*model.UserScore, error) {
	var score model.UserScore
	err := db.DB.Get(&score, `
		SELECT id, user_id, event_id, high_score
		FROM user_scores
		WHERE user_id = ? AND event_id = ?
		LIMIT 1
	`, userID, eventID)

	if err != nil {
		return nil, err
	}
	return &score, nil
}

func (r *userRepository) CreateUserScore(userID, eventID, score int) error {
	_, err := db.DB.Exec(`
		INSERT INTO user_scores (user_id, event_id, high_score)
		VALUES (?, ?, ?)
	`, userID, eventID, score)
	return err
}

func (r *userRepository) UpdateUserScore(userID, eventID, score int) error {
	_, err := db.DB.Exec(`
		UPDATE user_scores
		SET high_score = ?, updated_at = NOW()
		WHERE user_id = ? AND event_id = ?
	`, score, userID, eventID)
	return err
}
