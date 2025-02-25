package memberships

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"spotify_api/internal/model/memberhsips"
)

func (s *Service) SignUp(request memberhsips.SignUpRequest) error {
	existsUser, err := s.repository.GetUser(request.Username, request.Email, 0)
	if err != nil || errors.Is(err, sql.ErrNoRows) {
		log.Printf("failed to get user from db: %s", err.Error())
		return err
	}

	if existsUser != nil {
		return errors.New("email or username already used")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	userModel := memberhsips.User{
		Username:  request.Username,
		Email:     request.Username,
		Password:  string(hashedPassword),
		CreatedBy: request.Email,
		UpdatedBy: request.Email,
	}
	return s.repository.CreateUser(userModel)

}
