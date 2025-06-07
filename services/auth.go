package services

import (
	"errors"
	"time"

	"payslip-system/controllers"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	userCtrl *controllers.UserController
	jwtKey   []byte
}

func NewAuthService(userCtrl *controllers.UserController, jwtKey []byte) *AuthService {
	return &AuthService{
		userCtrl: userCtrl,
		jwtKey:   jwtKey,
	}
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userCtrl.FindByUsername(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid username or password")
	}

	// Verify password (assume stored hashed)
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"is_admin": user.IsAdmin,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// For logout, usually JWT tokens are stateless,
// but you can implement blacklist or token revocation if needed.
// For now, we'll just have a stub:

func (s *AuthService) Logout(tokenString string) error {
	// Implement token blacklist if needed.
	// Otherwise, frontend deletes token client-side.
	return nil
}
