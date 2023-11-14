package service

import (
	"context"
	"porters/internal/auth/model"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// PasswordHasher ����� �������
type PasswordHasher interface {
	Hash(password string) (string, error)
}

// UserRepository ����������� ������������
type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetByCredentials(ctx context.Context, username, password string) (model.User, error)
	Close()
}

// User ������������
type User struct {
	repository UserRepository
	hasher     PasswordHasher

	secret []byte
}

// NewUser ������� ������ ������������
func NewUser(repository UserRepository, hasher PasswordHasher, secret []byte) *User {
	return &User{
		repository: repository,
		hasher:     hasher,
		secret:     secret,
	}
}

// SignupUser ������������ ������ ������������
func (s *User) SignupUser(ctx context.Context, request model.SignupUserRequest) error {
	password, err := s.hasher.Hash(request.Password)
	if err != nil {
		return err
	}

	return s.repository.CreateUser(ctx, model.User{
		Username: request.Username,
		Password: password,
		Role:     request.Role,
	})
}

// LoginUser ������������ ����� ������ ������������
func (s *User) LoginUser(ctx context.Context, request model.LoginUserRequest) (string, error) {
	password, err := s.hasher.Hash(request.Password)
	if err != nil {
		return "", err
	}

	user, err := s.repository.GetByCredentials(ctx, request.Username, password)
	if err != nil {
		return "", err
	}

	// ��������� ������ �� 20 ����� ��� ������� ������������
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(20 * time.Minute).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    user.Role,
		Subject:   strconv.Itoa(user.ID),
	})

	// ���������� ������ ��������� ������
	response, err := token.SignedString(s.secret)
	if err != nil {
		return "", err
	}

	return response, nil
}
